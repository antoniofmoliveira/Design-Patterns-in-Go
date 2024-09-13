package main

import "fmt"

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

// 1st method
type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

func Swim() {
	fmt.Println("Swimming!")
}

// 2nd method
type Animal struct{}

func (r *Animal) Eat() {
	println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}

// 3rd method
type Swimmer interface {
	Swim()
}
type Trainer interface {
	Train()
}
type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim() {
	println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

func main() {
	// 1st method
	swimmer := CompositeSwimmerA{
		MySwim: Swim,
	}
	swimmer.MyAthlete.Train()
	swimmer.MySwim()

	// 2nd method
	fish := Shark{
		Swim: Swim,
	}
	fish.Eat()
	fish.Swim()

	// 3rd method
	swimmer2 := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}
	swimmer2.Train()
	swimmer2.Swim()
}
