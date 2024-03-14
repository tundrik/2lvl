package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

import "errors"


type ICar interface {
	SetName(name string)
	SetEngine(name string)
	GetName() string
	GetEngine() string
}

type Car struct {
	name string
	engine string
}

func (c *Car) SetName(name string) {
	c.name = name
}

func (c *Car) SetEngine(name string) {
	c.engine = name
}

func (c *Car) GetName() string {
	return c.name
}

func (c *Car) GetEngine() string {
	return c.engine
}

type BMW struct {
	Car
}

func NewBMW() ICar {
	return &BMW{
		Car: Car{
			name: "BMW",
			engine: "V8",
		},
	}
}

type Shkoda struct {
	Car
}

func NewShkoda() ICar {
	return &Shkoda{
		Car: Car{
			name: "Shkoda",
			engine: "V6",
		},
	}
}

func GetAvto(avtoType string) (ICar, error) {
	switch avtoType {
	case "BMW":
		return NewBMW(), nil
	case "Shkoda":
		return NewShkoda(), nil
	default:
		return nil, errors.New("wrong type")
	}
}

/*
	Использование:
	1) когда заранее неизвестны типы и зависимости объектов, с которыми должен работать код 
	2) когда нужна возможность пользователям расширять части фреймворка или библиотеки
	3) когда нужно экономить системные ресурсы, повторно используя уже созданные объекты, вместо порождения новых

	+:
	1) избавляет классы от привязки к конкретным классам продуктов
	2) выделяет код производства продуктов в одно место, упрощая поддержку кода
	3) упрощает добавление новых продуктов в программу
	4) реализует принцип открытости/закрытости

	-:
	1) может привести к созданию больших параллельных иерархий классов, так как для каждого класса
	продукта надо создать свой подкласс создателя
*/