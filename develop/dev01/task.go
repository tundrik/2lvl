package main
/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)


func main() {
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		os.Exit(1)
	}

	currentTime := time.Now()
	fmt.Println("текущее время: ", currentTime)
	fmt.Println("точное время: ", ntpTime)
}