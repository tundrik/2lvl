package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"errors"
	"fmt"
	"unicode"
)


var ErrIncorrectString = errors.New("incorrect string")

func Unpack(s string) (string, error) {
    sr := []rune(s)
    var s2 string
    var n int
    var backslash bool
    
    for i, item := range sr {
        if unicode.IsDigit(item) && i == 0 {
            return "", ErrIncorrectString
        }

        if unicode.IsDigit(item) && unicode.IsDigit(sr[i - 1]) && sr[i - 2] != '\\' {
            return "", ErrIncorrectString
        }             
        if item == '\\' && !backslash {
            backslash = true
            continue
        }   
        if backslash && unicode.IsLetter(item) {
            return "", ErrIncorrectString
        }
        if backslash {
            s2 += string(item)
            backslash = false
            continue
        }
        if unicode.IsDigit(item) {
            n = int(item - '0')
            if n == 0 {
                s2 = s2[:len(s2) - 1]
                continue
            }
            for j := 0; j < n - 1; j ++ {   
                s2 += string(sr[i - 1])
            } 
            continue     
        }     
        s2 += string(item)
    }

    return s2, nil
}



func main() {
    s, err := Unpack("a4bc2d5e")
	if err != nil {
		fmt.Printf("err %s", err)
	}

    fmt.Printf("a4bc2d5e - %s", s)
}
