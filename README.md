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

* pathLength(word string) int: calculates the total path length between the first and last letters of a word by summing the shortest path distances between adjacent letters using the shortestPath function.

* shortestPath(start, end rune) int: computes the shortest path between two letters on a keyboard layout using a breadth-first search algorithm. It uses a queue to keep track of characters to visit, a map to keep track of visited characters, and a map to store distances from the start character.

* findOptimalCombination(dictionary map[string]int) string: iterates through combinations of four words and calculates the total cost and length of each combination. It keeps track of the best combination with the lowest cost and a sum length within the desired range.

Time and complexity- O(n^4)

# Usage

The user should enter unique words consisting of English alphabet letters separated by spaces.

* Input examples: word_list.txt

# Future
What can be improved:
* memory usage 
* handle edge cases and errors
* chooseWords alternatives should be considered (dp, A*)
* tests
