package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/
/*
	Использование:
	1) когда вы хотите параметризировать объекты выполняемым действием
	2) когда вы хотите ставить операции в очередь, выполнять их по расписанию или передовать по сети 
	3) когда вам нужна операция отмены

	+:
	1) убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют
	2) позволяет реализовать простую отмену и повтор операций
	3) позволяет реализовать отложенный запуск операций
	4) позволяет собирать сложные команды из простых
	5) реализует принцип открытости/закрытости

	-:
	1) усложняет код программы из-за введения множества дополнительных классов
*/
import "fmt"


type Command interface {
	execute()
}

type Button struct {
	Command
}

func (b *Button) Press() {
	b.Command.execute()
}

type Device interface {
	on()
	off()
}

type OnCommand struct {
	Device
}

func (c *OnCommand) execute() {
	c.Device.on()
}

type OffCommand struct {
	Device
}

func (c *OffCommand) execute() {
	c.Device.off()
}

type TV struct {
	isWorking bool
}

func (t *TV) on() {
	t.isWorking = true
	fmt.Println("Turning TV on")
}

func (t *TV) off() {
	t.isWorking = false
	fmt.Println("Turning TV off")
}

