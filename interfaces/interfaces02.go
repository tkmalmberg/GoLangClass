/* RZFeeser | Alta3 Research
   Interfaces - Getting at interfaces */

package main

import "fmt"

type animal interface {
	breathe()
	walk()
	snooze() int
}

type tiger struct {
	age int
}

func (l tiger) breathe() {
	fmt.Println("tiger breathes")
}

func (l tiger) walk() {
	fmt.Println("tiger walk")
}

func (l tiger) snooze() int {
	return 5
}

type giraffe struct {
	age int
}

func (l giraffe) breathe() {
	fmt.Println("giraffe breathes")
}

func (l giraffe) walk() {
	fmt.Println("giraffe walk")
}

func (l giraffe) snooze() int {
	return 5
}

func main() {
	l := tiger{age: 10}
	callBreathe(l)
	callWalk(l)
	callSnooze(l)

	d := giraffe{age: 5}
	callBreathe(d)
	callWalk(d)
	callSnooze(d)
}

func callBreathe(a animal) {
	a.breathe()
}

func callWalk(a animal) {
	a.walk()
}

func callSnooze(a animal) {
	a.snooze()
}
