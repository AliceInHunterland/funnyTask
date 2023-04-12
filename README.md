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

* calculateCost(lastLetter, firstLetter byte) int: calculates the additional cost based on the last letter of the previous word and the first letter of the current word using the shortestPath function.

* shortestPath(start, end rune) int: computes the shortest path between two letters on a keyboard layout using a breadth-first search algorithm. It uses a queue to keep track of characters to visit, a map to keep track of visited characters, and a map to store distances from the start character.

* chooseWords(words []Word) []Word: iterates through combinations of four words and calculates the total cost and length of each combination. It keeps track of the best combination with the lowest cost and a sum length within the desired range.

Time and space complexity- O(n^2)

# Usage

The user should enter unique words consisting of English alphabet letters separated by spaces.

* Input examples:

```
lemon lime mango nectarine orange peach quince raspberry strawberry tangerine watermelon apricot blueberry cantaloupe blackberry cherry currant date elderberry fig gooseberry grapefruit honeydew ice cream
```
* Output: 
```
Best words:
Cost: 10 AddCost: 0 Length: 4 First Letter: l Last Letter: e Full Word: lime
Cost: 10 AddCost: 2 Length: 6 First Letter: c Last Letter: y Full Word: cherry
Cost: 13 AddCost: 4 Length: 5 First Letter: c Last Letter: m Full Word: cream
Cost: 15 AddCost: 8 Length: 6 First Letter: q Last Letter: e Full Word: quince
OVERALL COST: 62
PASSWORD: lime cherry cream quince

```
```
qwertyuiop ass lkj nm bvc nbvcxz

```
* Output: 
```
Best words:
Cost: 2 AddCost: 0 Length: 3 First Letter: l Last Letter: j Full Word: lkj
Cost: 2 AddCost: 2 Length: 3 First Letter: b Last Letter: c Full Word: bvc
Cost: 5 AddCost: 3 Length: 6 First Letter: n Last Letter: z Full Word: nbvcxz
Cost: 9 AddCost: 2 Length: 10 First Letter: q Last Letter: p Full Word: qwertyuiop
OVERALL COST: 25
PASSWORD: lkj bvc nbvcxz qwertyuiop

```
# Future
What can be improved:
* memory usage 
* addcost importance
* handle edge cases and errors
* chooseWords alternatives should be considered (dp, A*)
