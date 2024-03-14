package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/
/*
	Посетитель — это поведенческий паттерн проектирования, который позволяет создавать новые операции,
	не меняя классы объектов на которыми могут выполняться.

	Применяется когда над объектами сложной структуры объектов нужно выполнить некоторые, не связанные между собой операции,
	но вы не хотите "засорять" классы такими операциями.

	+ Новая функциональность в несколько классов добавляется сразу, не изменяя код этих классов.
	+ Объединяет родственные операции в одном классе.

	- Паттерн не оправдан, если иерархия компонентов часто меняется.
	- Может привести к нарушению инкапсуляции компонентов.
*/

type employee interface {
	accept(visitor)
}

type manager struct{}
type programmer struct{}
type hr struct{}

func (manager) accept(v visitor)    { v.visitManager() }
func (programmer) accept(v visitor) { v.visitProgrammer() }
func (hr) accept(v visitor)         { v.visitHr() }

type visitor interface {
	visitManager()
	visitProgrammer()
	visitHr()
}

type messageSender struct{}
type wageAssigner struct{}

func (messageSender) visitManager()    { fmt.Println("Make more money!") }
func (messageSender) visitProgrammer() { fmt.Println("Do less bugs!") }
func (messageSender) visitHr()         { fmt.Println("Bring me more slaves!") }

func (wageAssigner) visitManager()    { fmt.Println("too big to say") }
func (wageAssigner) visitProgrammer() { fmt.Println("300K/ns") }
func (wageAssigner) visitHr()         { fmt.Println("a cookie") }


func clientFunction() {
	employees := []employee{manager{}, programmer{}, hr{}}

	for _, e := range employees {
		fmt.Printf("Hey, %T! ", e)
		e.accept(messageSender{})
		fmt.Print("Your wage is ")
		e.accept(wageAssigner{})
	}
}
