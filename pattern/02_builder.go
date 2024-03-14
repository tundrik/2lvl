package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
	Использование:
	1) построение сложного объекта от его представления

	+:
	1) позволяет изменить внутреннее представление продукта
	2) инкапсулирует код для построения и представления
	3) обеспечивает контроль за этапами процесса строительства

	-:
	1) для каждого типа продукта должен быть создан отдельный строитель
	2) классы строителя должны быть изменяемыми
	3) может затруднить внедрение зависимостей
*/



// Сложный объект для строительства
type house struct {
	RoomNum      int
	EntrancesNum int
	Roof         string
	OutdoorColor string
	IndoorColor  string
}

/*

	КЛАССИЧЕСКИЙ СТРОИТЕЛЬ (с директором и конкретными строителями)

*/

// Абстрактный строитель
type IHouseBuilder interface {
	setRoomNumbers()
	setEntrances()
	setRoofType()
	setOutdoorColor()
	setIndoorColor()
	getHouse() house
}

// Конкретный строитель1
type houseBuilder1 struct {
	roomNum      int
	entrancesNum int
	roof         string
	outdoorColor string
	indoorColor  string
}

func (b *houseBuilder1) setRoomNumbers()  { b.roomNum = 6 }
func (b *houseBuilder1) setEntrances()    { b.entrancesNum = 2 }
func (b *houseBuilder1) setRoofType()     { b.roof = "Скатная" }
func (b *houseBuilder1) setOutdoorColor() { b.outdoorColor = "Кирпичный" }
func (b *houseBuilder1) setIndoorColor()  { b.indoorColor = "Бежевый" }

func (b *houseBuilder1) getHouse() house {
	return house{
		RoomNum:      b.roomNum,
		EntrancesNum: b.entrancesNum,
		Roof:         b.roof,
		OutdoorColor: b.outdoorColor,
		IndoorColor:  b.indoorColor,
	}
}

/* Представим, что тут находится houseBuilder2, который строит по другому */

// Директор. Нужен для того, чтобы контролировать порядок исполнения строительства
type director struct {
	builder IHouseBuilder
}

func newDirector(b IHouseBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b IHouseBuilder) {
	d.builder = b
}

func (d *director) buildHouse() house {
	d.builder.setRoomNumbers()
	d.builder.setEntrances()
	d.builder.setRoofType()
	d.builder.setOutdoorColor()
	d.builder.setIndoorColor()
	return d.builder.getHouse()
}

/*

	Fluent Builder (пользователь сам себе директор)

*/

type fluentHouseBuilder struct {
	h *house
}

func NewHouse() *fluentHouseBuilder {
	return &fluentHouseBuilder{}
}

func (*fluentHouseBuilder) SetStructure() *houseStructureBuilder {
	return &houseStructureBuilder{}
}

func (b *fluentHouseBuilder) Build() *house {
	return b.h
}

type houseStructureBuilder struct {
	fluentHouseBuilder
}

func (*houseStructureBuilder) SetAppearance() *houseAppearanceBuilder {
	return &houseAppearanceBuilder{}
}

func (b *houseStructureBuilder) SetRoomNumbers(roomNum int) *houseStructureBuilder {
	b.h.RoomNum = roomNum
	return b
}
func (b *houseStructureBuilder) SetEntrances(entrancesNum int) *houseStructureBuilder {

	b.h.EntrancesNum = entrancesNum
	return b
}

func (b *houseStructureBuilder) SetRoofType(roof string) *houseStructureBuilder {
	b.h.Roof = roof
	return b
}

type houseAppearanceBuilder struct {
	fluentHouseBuilder
}

func (b *houseAppearanceBuilder) SetOutdoorColor(outdoorColor string) *houseAppearanceBuilder {
	b.h.OutdoorColor = outdoorColor
	return b
}

func (b *houseAppearanceBuilder) SetIndoorColor(indoorColor string) *houseAppearanceBuilder {
	b.h.IndoorColor = indoorColor
	return b
}

/*
func clientFunction() {
	house := NewHouse().
		SetStructure().
			SetEntrances(2).
			SetRoomNumbers(3).
			SetRoofType("Двускатная").
		SetAppearance().
			SetIndoorColor("Черный").
			SetOutdoorColor("Зеленый").
		Build()
	fmt.Println(house)
}
*/