package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)


func main() {
	// Инициализация флагов
	columnPtr := flag.Int("k", 0, "указание колонки для сортировки")
	numericPtr := flag.Bool("n", false, "сортировать по числовому значению")
	reversePtr := flag.Bool("r", false, "сортировать в обратном порядке")
	uniquePtr := flag.Bool("u", false, "не выводить повторяющиеся строки")
	monthPtr := flag.Bool("M", false, "сортировать по названию месяца")
	ignoreTrailingSpacePtr := flag.Bool("b", false, "игнорировать хвостовые пробелы")
	checkSortedPtr := flag.Bool("c", false, "проверять отсортированы ли данные")
	numericSuffixPtr := flag.Bool("h", false, "сортировать по числовому значению с учетом суффиксов")
	flag.Parse()

	// Проверка наличия файла в качестве аргумента
	if flag.NArg() == 0 {
		log.Fatal("Необходимо указать файл для сортировки")
	}

	// Открытие файла на чтение
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Can not open file. Error: %s", err)
		}
	}(file)

	// Считывание строк из файла
	lines, err := readLines(file)
	if err != nil {
		log.Fatal(err)
	}

	// Функция сравнения строк для сортировки
	less := func(i, j int) bool {
		a := lines[i]
		b := lines[j]

		if *ignoreTrailingSpacePtr { // Обработка флага -b
			a = strings.TrimRight(a, " ")
			b = strings.TrimRight(b, " ")
		}

		if *numericSuffixPtr { // Обработка флага -h
			aVal, aSuffix := extractNumericSuffix(a)
			bVal, bSuffix := extractNumericSuffix(b)

			if aSuffix != bSuffix {
				return aSuffix < bSuffix
			}

			if *numericPtr {
				a = aVal
				b = bVal
			}
		}

		if *numericPtr { // Обработка флага -n
			aNum, errA := strconv.ParseFloat(a, 64)
			bNum, errB := strconv.ParseFloat(b, 64)

			if errA == nil && errB == nil {
				return aNum < bNum
			}
		}

		if *monthPtr { // Обработка флага -M
			aTime, aErr := parseMonth(a)
			bTime, bErr := parseMonth(b)

			if aErr == nil && bErr == nil {
				return aTime.Before(bTime)
			}
		}

		if *columnPtr > 0 { // Обработка флага -k
			aFields := strings.Fields(a)
			bFields := strings.Fields(b)

			if len(aFields) >= *columnPtr && len(bFields) >= *columnPtr {
				a = aFields[*columnPtr-1]
				b = bFields[*columnPtr-1]
			}
		}

		return a < b
	}

	// Сортировка строк
	sort.Slice(lines, func(i, j int) bool {
		if *reversePtr {
			return !less(i, j)
		}
		return less(i, j)
	})

	// Обработка флага -u
	if *uniquePtr {
		lines = removeDuplicates(lines)
	}

	// Вывод отсортированных строк
	for _, line := range lines {
		fmt.Println(line)
	}

	// Обработка флага -c
	if *checkSortedPtr {
		if !isSorted(lines, less) {
			log.Fatal("Данные не отсортированы")
		}
	}
}

// Функция для чтения строк из файла
func readLines(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Функция для удаления дубликатов из среза строк
func removeDuplicates(lines []string) []string {
	seen := make(map[string]struct{})
	result := make([]string, 0)

	for _, line := range lines {
		if _, ok := seen[line]; !ok {
			seen[line] = struct{}{}
			result = append(result, line)
		}
	}

	return result
}

// Функция для извлечения числового значения и суффикса из строки
func extractNumericSuffix(s string) (string, string) {
	for i := len(s) - 1; i >= 0; i-- {
		if !isNumeric(s[i]) {
			return s[:i+1], s[i+1:]
		}
	}
	return s, ""
}

// Функция для проверки наличия числового значения
func isNumeric(c byte) bool {
	return c >= '0' && c <= '9'
}

// Функция для парсинга строки в формате месяца
func parseMonth(s string) (time.Time, error) {
	return time.Parse("January", s)
}

// Функция для проверки отсортированности данных
func isSorted(lines []string, less func(i, j int) bool) bool {
	for i := 1; i < len(lines); i++ {
		if less(i, i-1) {
			return false
		}
	}
	return true
}
