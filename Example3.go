package main

import (
	"fmt"
	"math"
)

func main() {
	p1 := Person{"foo", 10, 100.0, true}
	fmt.Println("first print: ")
	p1.printPropsOfPerson()

	p1.setName("newfoo")
	p1.setAge(20)
	p1.setWage(200.0)
	p1.setActive(false)
	fmt.Println("second print: ")
	p1.printPropsOfPerson()

	rect := Rectangle{10, 20}
	c := Circle{10}
	t := Triangle{10, 20}

	getArea(rect)
	getArea(c)
	getArea(t)

}

//interface vs struct example
type Shape interface {
	area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	x float64
	y float64
}

func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (tri Triangle) area() float64 {
	return tri.x * tri.y / 2
}

func getArea(s Shape) {
	fmt.Println("area = ", s.area())
}

//struct example
type Person struct {
	name   string
	age    int32
	wage   float64
	active bool
}

func (p *Person) setName(name string) {
	p.name = name
}

func (p *Person) getName() string {
	return p.name
}

func (p *Person) setAge(age int32) {
	p.age = age
}

func (p *Person) getAge() int32 {
	return p.age
}

func (p *Person) setWage(wage float64) {
	p.wage = wage
}

func (p *Person) getWage() float64 {
	return p.wage
}

func (p *Person) setActive(active bool) {
	p.active = active
}

func (p *Person) isActive() bool {
	return p.active
}

func (p *Person) printPropsOfPerson() {
	fmt.Println("name = ", p.name)
	fmt.Println("age = ", p.age)
	fmt.Println("wage = ", p.wage)
	fmt.Println("active = ", p.active)
}
