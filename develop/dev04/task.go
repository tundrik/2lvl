package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagramSets(words *[]string) *map[string][]string {
	anagramSets := make(map[string][]string)

	for _, word := range *words {
		word = strings.ToLower(word)
		sortedWord := sortString(word)
		anagramSets[sortedWord] = append(anagramSets[sortedWord], word)
	}

	for key, value := range anagramSets {
		if len(value) == 1 {
			delete(anagramSets, key)
		}
	}

	return &anagramSets
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}

	anagramSets := findAnagramSets(&words)

	for _, value := range *anagramSets {
		fmt.Printf("Множество анаграмм для слова %s: %v\n", value[0], value)
	}
}
