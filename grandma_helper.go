package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strings"
	"time"
)

var keyboard = map[byte][]byte{
	'a': {'s', 'z', 'q'},
	'b': {'v', 'g', 'n'},
	'c': {'x', 'd', 'f', 'v'},
	'd': {'s', 'e', 'f', 'c', 'x'},
	'e': {'w', 'r', 'd'},
	'f': {'d', 'r', 'g', 'v', 'c'},
	'g': {'f', 't', 'h', 'b', 'v'},
	'h': {'g', 'y', 'j', 'n', 'b'},
	'i': {'u', 'o', 'k'},
	'j': {'h', 'u', 'k', 'm', 'n'},
	'k': {'j', 'i', 'l', 'm'},
	'l': {'k', 'o', 'p'},
	'm': {'n', 'j', 'k'},
	'n': {'b', 'h', 'j', 'm'},
	'o': {'i', 'p', 'l'},
	'p': {'o', 'l'},
	'q': {'w', 'a'},
	'r': {'e', 't', 'f'},
	's': {'a', 'w', 'd', 'x', 'z'},
	't': {'r', 'y', 'g'},
	'u': {'y', 'i', 'j'},
	'v': {'c', 'f', 'g', 'b'},
	'w': {'q', 's', 'e'},
	'x': {'z', 's', 'd', 'c'},
	'y': {'t', 'u', 'h'},
	'z': {'a', 's', 'x'},
}

// Function to calculate the cost between two letters
func calculateCost(lastLetter byte, firstLetter byte, keyboard map[byte][]byte) int {
	if lastLetter == firstLetter {
		return 0
	}
	if neighbors, ok := keyboard[lastLetter]; ok {
		for _, neighbor := range neighbors {
			if neighbor == firstLetter {
				return 1
			}
		}
		return shortestPath(lastLetter, firstLetter, keyboard)
	}
	return 0
}

// Function to find the shortest path between two letters (BFS)
func shortestPath(start byte, end byte, keyboard map[byte][]byte) int {
	queue := []byte{start}
	visited := map[byte]bool{start: true}
	distances := map[byte]int{start: 0}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if neighbors, ok := keyboard[curr]; ok {
			for _, neighbor := range neighbors {
				if !visited[neighbor] {
					visited[neighbor] = true
					distances[neighbor] = distances[curr] + 1
					queue = append(queue, neighbor)
				}
			}
		}
		if curr == end {
			return distances[curr]
		}
	}
	return -1
}

// Create Matrix with cost of move from letter to letter
func createLettersMatrix(keyboard map[byte][]byte) map[byte]map[byte]int {
	letters := make([]byte, 0, len(keyboard))
	for letter := range keyboard {
		letters = append(letters, letter)
	}
	lettersMatrix := make(map[byte]map[byte]int, len(keyboard))
	for _, lastLetter := range letters {
		lettersMatrix[lastLetter] = make(map[byte]int, len(keyboard))
		for _, firstLetter := range letters {
			lettersMatrix[lastLetter][firstLetter] = calculateCost(lastLetter, firstLetter, keyboard)
		}
	}
	return lettersMatrix
}

// For each word in List find the lowest cost with same key's parameters (first letter, last letter, lenght): [cost, word]
func calculateWordCosts(wordList []string, lettersMatrix map[byte]map[byte]int) map[[3]interface{}][2]interface{} {
	wordCosts := make(map[[3]interface{}][2]interface{}, len(wordList))
	for _, word := range wordList {
		firstLetter := word[0]
		lastLetter := word[len(word)-1]
		length := len(word)

		cost := 0
		for i := 1; i < len(word); i++ {
			cost += lettersMatrix[word[i-1]][word[i]]
		}

		key := [3]interface{}{firstLetter, lastLetter, length}
		if value, ok := wordCosts[key]; ok {
			if cost < value[0].(int) {
				wordCosts[key] = [2]interface{}{cost, []string{word}}
			}
		} else {
			wordCosts[key] = [2]interface{}{cost, []string{word}}
		}
	}
	return wordCosts
}

// Concatenate words in value and recalculate keys
func concatenateWords(wordCosts map[[3]interface{}][2]interface{}, lettersMatrix map[byte]map[byte]int) map[[3]interface{}][3]interface{} {
	concatenatedWordCosts := make(map[[3]interface{}][3]interface{})

	// Create a slice of all the keys in the wordCosts map
	keys := make([][3]interface{}, 0, len(wordCosts))
	for key := range wordCosts {
		keys = append(keys, key)
	}

	// Initialize the matrix with the cost of concatenating each word with itself, which is 0
	matrix := make([][]int, len(keys))
	for i := range matrix {
		matrix[i] = make([]int, len(keys))
		matrix[i][i] = 0
	}

	// Iterate through all possible pairs of words and update the matrix if the cost of concatenating the pair is less than the current value in the matrix
	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			key1 := keys[i]
			key2 := keys[j]
			firstLetter1 := key1[0].(byte)
			lastLetter1 := key1[1].(byte)
			length1 := key1[2].(int)
			words1 := wordCosts[key1][1].([]string)
			firstLetter2 := key2[0].(byte)
			lastLetter2 := key2[1].(byte)
			length2 := key2[2].(int)
			words2 := wordCosts[key2][1].([]string)

			if length1+length2 <= 24 {
				cost := wordCosts[key1][0].(int) + lettersMatrix[lastLetter1][firstLetter2] + wordCosts[key2][0].(int)
				newKey := [3]interface{}{firstLetter1, lastLetter2, length1 + length2}

				// Update the matrix if the cost of concatenating the pair is less than the current value in the matrix
				if value, ok := concatenatedWordCosts[newKey]; !ok || cost < value[0].(int) {
					concatenatedWordCosts[newKey] = [3]interface{}{cost, words1[0], words2[0]}
				}
			}
		}
	}

	return concatenatedWordCosts
}

// Find 4 unique words for password with the lowest cost and length in range of 20 to 24
func findLowestCostWords(inputDict map[[3]interface{}][3]interface{}, lettersMatrix map[byte]map[byte]int) []string {

	// Initialize variables to hold the lowest cost and length, and corresponding words for the two keys
	lowestCost := math.MaxInt64
	lowestCostAndLengthWordsForKey := []string{}

	// Loop through all key-value pairs in the input dictionary
	for key1, value1 := range inputDict {
		lastLetter1 := key1[1].(byte)
		words1 := value1[1].(string)
		words2 := value1[2].(string)

		for key2, value2 := range inputDict {
			// Skip the current key or any key that has already been processed
			if key2 == key1 || key2[0].(byte) <= key1[1].(byte) {
				continue
			}

			firstLetter2 := key2[0].(byte)
			words3 := value2[1].(string)
			words4 := value2[2].(string)
			// Check if the combined length of the two keys is between 20 and 24
			length := key1[2].(int) + key2[2].(int)
			if length < 20 || length > 24 {
				continue
			}

			// Combine the sets of words from the two keys, and check if the resulting set has 4 unique words
			wordSet := make(map[string]bool)
			for _, w := range []string{words1, words2, words3, words4} {
				wordSet[w] = true
			}
			if len(wordSet) != 4 {
				continue
			}

			// Calculate the cost of combining the two keys
			cost := value1[0].(int) + lettersMatrix[lastLetter1][firstLetter2] + value2[0].(int)

			// Update the lowest cost and corresponding words if the current cost is lower
			if cost < lowestCost {
				lowestCost = cost
				lowestCostAndLengthWordsForKey = []string{words1, words2, words3, words4}
			}
		}
	}

	return lowestCostAndLengthWordsForKey
}

func checkWordList(wordList []string) error {
	// check if there are at least 4 words
	if len(wordList) < 4 {
		return fmt.Errorf("the word list should contain at least 4 words")
	}

	// check if len of list is only 4 words then sum of lengths of each word should be more then 20
	if len(wordList) == 4 {
		totalLength := 0
		for _, word := range wordList {
			totalLength += len(word)
		}
		if totalLength < 20 {
			return fmt.Errorf("the total length of the words in the list should be at least 20")
		}
	}

	// check if there are non-English letters in the word list
	regex := regexp.MustCompile("^[a-zA-Z]+$")
	for _, word := range wordList {
		if !regex.MatchString(word) {
			return fmt.Errorf("the word list contains non-English letters")
		}
	}

	return nil
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: ./program_name path_to_word_list_file")
		os.Exit(1)
	}

	filePath := os.Args[1]
	// Read word list from file
	wordListBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	wordList := strings.Split(string(wordListBytes), "\n")
	for i, word := range wordList {
		wordList[i] = strings.ToLower(word) // convert to lowercase
	}

	err = checkWordList(wordList)
	if err != nil {
		panic(err)
	}

	start := time.Now()
	lettersMatrix := createLettersMatrix(keyboard)
	wordCosts := calculateWordCosts(wordList, lettersMatrix)
	wordCosts1 := concatenateWords(wordCosts, lettersMatrix)
	wordCosts2 := findLowestCostWords(wordCosts1, lettersMatrix)

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)
	str := strings.Join(wordCosts2, " ")
	fmt.Println("Possible password is: ", str)

}
