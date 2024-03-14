package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/
/*
	Использование:
	1) когда программа должна обрабатывать разнообразные запросы несколькими способами,
	но заранее неизвестно, какие конкретно запросы будут приходить и какие обработчики для них понадобятся
	2) когда важно, чтобы обработчики выполнялись один за другим в строгом порядке
	3) когда набор объектов, способных обработать запрос, должен задаваться динамически

	+:
	1) уменьшает зависимость между клиентом и обработчиками
	2) реализует принцип единственной обязаности
	3) реализует принцип открытости/закрытости

	-:
	1) запрос может остаться никем не обработанным
*/
import "fmt"


type BrokenCar struct {
	name string
	isEngine bool
	isWiring bool
	isWheels bool
}

func NewBrokenCar(name string, isEngine, isWiring, isWheels bool) *BrokenCar {
	return &BrokenCar{
		name: name,
		isEngine: isEngine,
		isWiring: isWiring,
		isWheels: isWheels,
	} 
}

type CarService interface {
	Execute(*BrokenCar)
	SetNext(CarService)
}

type EngineMaster struct {
	next CarService
}

func (m *EngineMaster) Execute(brokenCar *BrokenCar) {
	if brokenCar.isEngine {
		fmt.Println("Engine is OK")
		return
	}
	fmt.Println("Repair engine")
	brokenCar.isEngine = true
}

func (m *EngineMaster) SetNext(next CarService) {
	m.next = next
}

type WheelsMaster struct {
	next CarService
}

func (m *WheelsMaster) Execute(brokenCar *BrokenCar) {
	if brokenCar.isWheels {
		fmt.Println("Wheels is OK")
		m.next.Execute(brokenCar)
		return
	}
	fmt.Println("Repair wheels")
	brokenCar.isWheels = true
	m.next.Execute(brokenCar)
}

func (m *WheelsMaster) SetNext(next CarService) {
	m.next = next
}

type WiringMaster struct {
	next CarService
}

func (m *WiringMaster) Execute(brokenCar *BrokenCar) {
	if brokenCar.isWiring {
		fmt.Println("Wiring is OK")
		m.next.Execute(brokenCar)
		return
	}
	fmt.Println("Repair wiring")
	brokenCar.isWiring = true
	m.next.Execute(brokenCar)
}

func (m *WiringMaster) SetNext(next CarService) {
	m.next = next
}

