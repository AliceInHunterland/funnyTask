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

* concatenateWords(wordCosts map[[3]interface{}][2]interface{}, lettersMatrix map[byte]map[byte]int) map[[3]interface{}][3]interface{}:returns a new map where the keys are a set of concatenated words, and the values are a tuple of the total cost of concatenation and the list of concatenated words. It accomplishes this by iterating over all possible pairs of keys and updating a matrix of costs to concatenate each pair of keys. The cost of concatenating two sets of words is the sum of the costs of the individual sets plus the cost of moving from the last letter of the first set to the first letter of the second set in the keyboard matrix. The final result is the lowest cost concatenation for each set of words with the same first and last letter and combined length of at most 24.

* findLowestCostWords(inputDict map[[3]interface{}][3]interface{}, lettersMatrix map[byte]map[byte]int) []string: checks if the combined length of the two keys is between 20 and 24, combines the sets of words from the two keys, and checks if the resulting set has 4 unique words and finds the lowest cost.

Time and complexity- O(n^2)

* why not dp?
need to find unique 4 elements.

* Structure: (first_letter, last_letter, word_length): [cost, word1, word2]

# Usage

The user should enter unique words consisting of English alphabet letters separated by spaces.

* Input examples: word_list.txt








