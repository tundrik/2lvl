package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	after := flag.Int("A", 0, "печатать +N строк после совпадения")
	before := flag.Int("B", 0, "печатать +N строк до совпадения")
	context := flag.Int("C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	count := flag.Bool("c", false, "(количество строк)")
	ignoreCase := flag.Bool("i", false, "(игнорировать регистр)")
	invert := flag.Bool("v", false, "(вместо совпадения, исключать)")
	fixed := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	withLineNum := flag.Bool("n", false, "печатать номер строки")
	flag.Parse()

	pattern := flag.Arg(0)

	file, err := os.Open(flag.Arg(1))
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	defer file.Close()

	matchCount := 0

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()

		matches := false

		if *fixed {
			if strings.Contains(line, pattern) {
				matches = true
			}
		} else {
			if *ignoreCase {
				line = strings.ToLower(line)
				pattern = strings.ToLower(pattern)
			}

			if strings.Contains(line, pattern) {
				matches = true
			}
		}

		if (*invert && !matches) || (!*invert && matches) {
			if *withLineNum {
				fmt.Print(lineNum, ": ")
			}

			fmt.Println(line)

			if *count {
				matchCount++
				continue
			}

			if *before > 0 {
				for i := lineNum - *before; i < lineNum; i++ {
					if i > 0 {
						scanner.Scan()
						fmt.Println(scanner.Text())
					}
				}
			}

			if *after > 0 {
				for i := lineNum + 1; i <= lineNum+*after; i++ {
					scanner.Scan()
					fmt.Println(scanner.Text())
				}
			}

			if *context > 0 {
				for i := lineNum - *context; i < lineNum+*context; i++ {
					if i > 0 {
						scanner.Scan()
						fmt.Println(scanner.Text())
					}
				}
			}
		}

		lineNum++
	}

	if *count {
		fmt.Println("lines count:", matchCount)
	}
}
