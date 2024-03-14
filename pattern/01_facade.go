package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/
/*
	применимость:
	Упрощение сложной системы. 
	Если у вас сложная подсистема с многочисленными классами и взаимодействиями,
	шаблон «Фасад» может упростить взаимодействие клиента, предоставляя единую точку входа.

	Уменьшение зависимостей. Если вы хотите уменьшить связь между клиентским кодом и классами подсистемы, 
	шаблон Фасад может помочь, предоставляя уровень абстракции.

	Улучшение удобства сопровождения. Фасады могут сделать код более удобным в сопровождении, 
	инкапсулируя изменения в подсистеме, уменьшая влияние на клиентский код.

	+:
	Упрощенный клиентский код. Клиентам подсистемы не нужно знать подробности того, как работает каждый компонент. 
	Они взаимодействуют с фасадом, который предоставляет понятный и простой API.

	Уменьшенная связь: шаблон «Фасад» уменьшает связь между клиентским кодом и подсистемой, 
	упрощая изменение или замену компонентов подсистемы, не затрагивая клиенты.

	Улучшенная ремонтопригодность: по мере развития подсистемы изменения инкапсулируются внутри фасада, 
	ограничивая влияние на клиентский код.

	-:
	рискует стать антипаттерном "божественный объект"
*/

import (
    "fmt"
)

// Подсистемы

type Product struct {
	name string
}

func (p *Product) GetProduct() {
	p.name = "Pizza"
	fmt.Println("Product Fetched: ", p.name)
}

type Payment struct {
	sum    float32
	status bool
}

func (p *Payment) DoPayment(sum float32) {
	p.sum = sum
	p.status = true
	fmt.Println("Payment Done Successfuly")
}

type Invoice struct {
	status bool
}

func (i *Invoice) SendInvoice(payment Payment) {
	if payment.status {
		fmt.Println("Invoice Sent Successfuly")
		i.status = true
	}
	fmt.Println("Invoice Failed")
	i.status = false
}

// Фасад. Создает унифицированный интерфейс с подсистемами выше

type Order struct {
}

func (po *Order) PlaceOrder() {
	fmt.Println("Start placing order")
	prod := Product{}
	prod.GetProduct()

	payment := Payment{}
	payment.DoPayment(100)

	invoice := Invoice{}
	invoice.SendInvoice(payment)
}