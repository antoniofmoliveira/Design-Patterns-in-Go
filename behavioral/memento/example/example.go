package main

import "fmt"

type Command interface {
	GetValue() interface{}
}

type Volume byte

func (v Volume) GetValue() interface{} {
	return v
}

type Mute bool

func (m Mute) GetValue() interface{} {
	return m
}

type Memento struct {
	memento Command
}

type originator struct {
	Command Command
}

func (o *originator) NewMemento() Memento {
	return Memento{memento: o.Command}
}

func (o *originator) ExtractAndStoreCommand(m Memento) {
	o.Command = m.memento
}

type careTaker struct {
	mementoStack []Memento
}

func (c *careTaker) Add(m Memento) {
	c.mementoStack = append(c.mementoStack, m)
}
func (c *careTaker) Pop() Memento {
	if len(c.mementoStack) > 0 {
		tempMemento := c.mementoStack[len(c.mementoStack)-1]
		c.mementoStack = c.mementoStack[0 : len(c.mementoStack)-1]
		return tempMemento
	}
	return Memento{}
}

type MementoFacade struct {
	originator originator
	careTaker  careTaker
}

func (m *MementoFacade) SaveSettings(s Command) {
	m.originator.Command = s
	m.careTaker.Add(m.originator.NewMemento())
}
func (m *MementoFacade) RestoreSettings(i int) Command {
	m.originator.ExtractAndStoreCommand(m.careTaker.Pop())
	return m.originator.Command
}

func assertAndPrint(c Command) {
	switch cast := c.(type) {
	case Volume:
		fmt.Printf("Volume:\t%d\n", cast)
	case Mute:
		fmt.Printf("Mute:\t%t\n", cast)
	}
}

func main() {
	m := MementoFacade{}
	m.SaveSettings(Volume(4))
	m.SaveSettings(Mute(false))
	assertAndPrint(m.RestoreSettings(0))
	assertAndPrint(m.RestoreSettings(1))
}
