package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

var keyboard = map[rune][]rune{
	'q': {'w', 'a'},
	'w': {'q', 'e', 's'},
	'e': {'w', 'r', 'd'},
	'r': {'e', 't', 'f'},
	't': {'r', 'y', 'g'},
	'y': {'t', 'u', 'h'},
	'u': {'y', 'i', 'j'},
	'i': {'u', 'o', 'k'},
	'o': {'i', 'p', 'l'},
	'p': {'o', 'l'},
	'a': {'q', 's', 'z'},
	's': {'w', 'a', 'd', 'z', 'x'},
	'd': {'e', 's', 'f', 'x', 'c'},
	'f': {'r', 'd', 'g', 'c', 'v'},
	'g': {'t', 'f', 'h', 'v', 'b'},
	'h': {'y', 'g', 'j', 'b', 'n'},
	'j': {'u', 'h', 'k', 'n', 'm'},
	'k': {'i', 'j', 'l', 'm'},
	'l': {'o', 'k'},
	'z': {'a', 's', 'x'},
	'x': {'s', 'd', 'z', 'c'},
	'c': {'d', 'f', 'x', 'v'},
	'v': {'f', 'g', 'c', 'b'},
	'b': {'g', 'h', 'v', 'n'},
	'n': {'h', 'j', 'b', 'm'},
	'm': {'j', 'k', 'n'},
}

type Word struct {
	Cost        int
	Length      int
	FirstLetter byte
	LastLetter  byte
	FullWord    string
}

func pathLength(word string) int {
	length := 0
	var prevChar rune
	for _, char := range word {
		if prevChar != 0 {
			length = length + shortestPath(prevChar, char)
		}
		prevChar = char
	}
	return length
}

func calculateCost(lastLetter, firstLetter byte) int {
	// Function to calculate the additional cost based on the last letter of the previous word
	// and the first letter of the current word

	if lastLetter == firstLetter {
		return 0
	}
	if lastLetter != 0 {
		return shortestPath(rune(lastLetter), rune(firstLetter))
	}
	return 0
}

func shortestPath(start, end rune) int {
	queue := []rune{start}                // queue to store characters to visit
	visited := map[rune]bool{start: true} // map to keep track of visited characters
	distances := map[rune]int{start: 0}   // map to store distances from start character
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		neighbors := keyboard[curr]
		for _, neighbor := range neighbors {
			if _, ok := visited[neighbor]; !ok {
				visited[neighbor] = true
				distances[neighbor] = distances[curr] + 1
				queue = append(queue, neighbor)
			}
		}
		if curr == end {
			return distances[curr]
		}
	}
	return -1 // return -1 if end character is not reachable from start character
}

func chooseWords(words []Word) []Word {
	n := len(words)

	// Initialize variables to keep track of the best combination
	bestSumCost := math.MaxInt64
	var bestWords []Word

	// Precompute costs
	costs := make([][]int, n)
	for i := 0; i < n; i++ {
		costs[i] = make([]int, n)
		for j := 0; j < n; j++ {
			costs[i][j] = calculateCost(words[i].LastLetter, words[j].FirstLetter)
		}
	}

	// Iterate through combinations of 4 words
	for i := 0; i < n-3; i++ {
		sumCost := 0
		sumLength := 0

		for j := i; j < i+4; j++ {
			// Calculate the additional cost based on precomputed costs
			var additionalCost int
			if j > 0 {
				additionalCost = costs[j-1][j]
			}
			sumCost += words[j].Cost + additionalCost
			sumLength += words[j].Length
		}

		// Check if the current combination has a lower cost and the sum length is within the desired range
		if sumLength >= 20 && sumLength <= 24 {
			if sumCost < bestSumCost {
				bestSumCost = sumCost
				bestWords = words[i : i+4]
			}
			return bestWords
		}
	}

	return bestWords
}

func main() {

	//my_words := []string{"hello", "world", "python", "qwerty", "asdfghp", "zxcvbn", "name", "my", "kate", "i", "am", "a", "good", "software", "engineer"}
	fmt.Println("Enter words separated by spaces:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') // Read input until newline character is encountered
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	my_words := strings.Fields(input) // Split input into words
	fmt.Println("Words read from console:", my_words)

	words := []Word{}
	for _, w := range my_words {
		word := Word{
			Cost:        pathLength(w), // Set Cost to the length of the word
			Length:      len(w),        // Set Length to the length of the word
			FirstLetter: w[0],          // Set FirstLetter to the first byte of the word
			LastLetter:  w[len(w)-1],   // Set LastLetter to the last byte of the word
			FullWord:    w,             // Set FullWord to the word itself
		}
		words = append(words, word)
	}

	//fmt.Println(words)
	sort.Slice(words, func(i, j int) bool {
		return words[i].Cost < words[j].Cost
	})
	//fmt.Println(words)
	// Call the chooseWords function to get the best combination of 4 words
	bestWords := chooseWords(words)
	var sumCost = 0
	// Print the result
	if len(bestWords) == 4 {
		fmt.Println("Best words:")
		for i := 0; i < len(bestWords); i++ {
			word := bestWords[i]
			sumCost = sumCost + word.Cost
			var addCost int
			if i > 0 {
				addCost = calculateCost(bestWords[i-1].LastLetter, word.FirstLetter)
				sumCost = sumCost + addCost
			}

			fmt.Println("Cost:", word.Cost, "AddCost:", addCost, "Length:", word.Length, "First Letter:", string(word.FirstLetter), "Last Letter:", string(word.LastLetter), "Full Word:", word.FullWord)

		}
		fmt.Println("OVERALL COST:", sumCost)
		fmt.Print("PASSWORD: ")
		for i := 0; i < len(bestWords); i++ {
			fmt.Print(bestWords[i].FullWord)
			if i < len(bestWords)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	} else {
		fmt.Println("No valid combination found.")
	}
}
