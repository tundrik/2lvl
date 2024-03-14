package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)


func main() {
	if err := cut(os.Stdin, os.Stdout, os.Args[1:]...); err != nil {
		_, err := fmt.Fprintln(os.Stderr, "Error:", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

func cut(input io.Reader, output io.Writer, args ...string) error {
	fields, delimiter := parseArgs(args)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, delimiter) {
			fields := processLine(line, delimiter, fields)
			_, err := fmt.Fprintln(output, strings.Join(fields, delimiter))
			if err != nil {
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func parseArgs(args []string) (fields []int, delimiter string) {
	var fieldStr string
	var delimiterStr string

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-f":
			if i+1 < len(args) {
				fieldStr = args[i+1]
				i++
			}
		case "-d":
			if i+1 < len(args) {
				delimiterStr = args[i+1]
				i++
			}
		}
	}

	delimiter = "\t"
	if delimiterStr != "" {
		delimiter = delimiterStr
	}
	fields = parseFields(fieldStr)

	return fields, delimiter
}

func parseFields(fieldStr string) []int {
	fields := make([]int, 0)

	if fieldStr == "" {
		return fields
	}

	fieldNumbers := strings.Split(fieldStr, ",")
	for _, number := range fieldNumbers {
		index := parseFieldNumber(number)
		if index > -1 {
			fields = append(fields, index)
		}
	}

	return fields
}

func parseFieldNumber(fieldNumber string) int {
	number := strings.TrimSpace(fieldNumber)
	if number == "" {
		return -1
	}
	index, err := strconv.Atoi(number)
	if err != nil {
		return -1
	}
	return index - 1
}

func processLine(line string, delimiter string, fields []int) []string {
	columns := strings.Split(line, delimiter)
	output := make([]string, 0)

	if len(fields) > 0 {
		for _, index := range fields {
			if index > -1 && index < len(columns) {
				output = append(output, columns[index])
			}
		}
	} else {
		output = append(output, columns...)
	}

	return output
}
