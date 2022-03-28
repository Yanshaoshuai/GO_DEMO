package main

type IAnimal interface {
	Sleep()
	Eat()
	Name() string
}
type FlyableAnimal interface {
	Fly()
	IAnimal
}

type Owl struct {
}

func (o *Owl) Fly() {
	println(o.Name() + " is flying")
}

func (o *Owl) Name() string {
	return "Owl"
}

func (o *Owl) Sleep() {
	//TODO implement me
	panic("implement me")
}

func (o *Owl) Eat() {
	//TODO implement me
	panic("implement me")
}

func Fly(animal FlyableAnimal) {
	animal.Fly()
}

func main() {
	owl := Owl{}
	Fly(&owl)
}
