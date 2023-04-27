# Task
"Бабушке нужно сгенерировать пароль, она слышала, что если взять четыре
слова из английского словаря, то можно получить хороший вариант. Но
проблема в том, что бабушка печатает одним пальцем и перемещать палец по
клавиатуре ей затруднительно, поэтому необходимо использовать такие слова,
которые эти перемещения минимизируют (считаются перемещения по четырём
сторонам, например, от "F" до "H" необходимо выполнить два перемещения, а
от "A" до "E" три), при том, что общая длина пароля будет от 20 до 24 символов.
Требуется найти наилучший пароль для бабушки."

* Реализовать решение на языках Си или Golang под Linux.

# Solution
Key functions:

* shortestPath(start byte, end byte, keyboard map[byte][]byte):  computes the shortest path between two letters on a keyboard layout using a breadth-first search algorithm. It uses a queue to keep track of characters to visit, a map to keep track of visited characters, and a map to store distances from the start character.

* createLettersMatrix(keyboard map[byte][]byte) map[byte]map[byte]int: creates a matrix of costs to move from one letter to another on the keyboard

* calculateWordCosts(wordList []string, lettersMatrix map[byte]map[byte]int) map[[3]interface{}][2]interface{}: returns a map of the lowest cost for each set of words with the same first letter, last letter, and length.

* concatenateWords(wordCosts map[[3]interface{}][2]interface{}, lettersMatrix map[byte]map[byte]int) map[[3]interface{}][3]interface{}:rThis function creates a new map that contains concatenated words as keys and their total cost and list of words as values. It works by calculating the cost of concatenating each pair of keys and choosing the lowest cost for sets of words with the same first and last letter, and a combined length of up to 24.

* findLowestCostWords(inputDict map[[3]interface{}][3]interface{}, lettersMatrix map[byte]map[byte]int) []string: checks if the combined length of the two keys is between 20 and 24, combines the sets of words from the two keys, and checks if the resulting set has 4 unique words and finds the lowest cost.

Time and complexity- O(n^2)

* why not dp?
need to find unique 4 elements.

* Structure: (first_letter, last_letter, word_length): [cost, word1, word2]

# Usage

The user should enter unique words consisting of English alphabet letters separated by spaces.

* Input examples: word_list.txt








