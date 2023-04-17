package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

// QueueItem represents an item in the queue for BFS
type QueueItem struct {
	char rune
	dist int
}

// Function to find the shortest path between two letters
func shortestPath(start rune, end rune, keyboard map[rune][]rune) int {
	queue := []QueueItem{{char: start, dist: 0}}
	visited := map[rune]bool{start: true}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		neighbors := keyboard[curr.char]
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				visited[neighbor] = true
				if neighbor == end {
					return curr.dist + 1
				}
				queue = append(queue, QueueItem{char: neighbor, dist: curr.dist + 1})
			}
		}
	}

	return -1
}

func pathLength(word string, keyboard map[rune][]rune) int {
	length := 0
	prevChar := rune(0)

	for _, char := range word {
		if prevChar != rune(0) {
			length += shortestPath(prevChar, char, keyboard)
		}
		prevChar = char
	}

	return length
}
func findOptimalCombination(dictionary map[string]int) string {
	words := make([]string, 0, len(dictionary))
	for word := range dictionary {
		words = append(words, word)
	}
	sort.Slice(words, func(i, j int) bool {
		return dictionary[words[i]] > dictionary[words[j]]
	}) // Sort words by price in descending order
	bestPrice := int(^uint(0) >> 1)
	var bestCombination []string

	// Generate all combinations of 4 words with pruning for length constraint
	for i := 0; i < len(words)-3; i++ {
		for j := i + 1; j < len(words)-2; j++ {
			for k := j + 1; k < len(words)-1; k++ {
				for l := k + 1; l < len(words); l++ {
					combination := []string{words[i], words[j], words[k], words[l]}
					wordLen := 0
					for _, word := range combination {
						wordLen += len(word)
					}
					if 20 <= wordLen && wordLen <= 24 {
						price := 0
						for _, word := range combination {
							price += dictionary[word]
						}
						if price < bestPrice {
							// Update the best price and combination if a lower price is found
							bestPrice = price
							bestCombination = combination
						}
					}
				}
			}
		}
	}
	if len(bestCombination) > 0 {
		return fmt.Sprintf("Optimal combination: %s", strings.Join(bestCombination, " "))
	} else {
		return "No valid combination found."
	}
}
func main() {
	// Read word list from file
	wordListBytes, err := ioutil.ReadFile("/Users/ekaterinapavlova/Projects/my-golang-workspace/src/word_list.txt")
	if err != nil {
		panic(err)
	}
	wordList := strings.Split(string(wordListBytes), "\n")

	keyboard := map[rune][]rune{
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

	mWords := make(map[string]int)
	for _, word := range wordList {
		mWords[word] = pathLength(word, keyboard)
	}

	fmt.Println(mWords)
	result := findOptimalCombination(mWords)
	fmt.Println(result)
}
