package main

import "fmt"

func main() {
	type Student struct {
		FirstName string
		LastName  string
	}

	var firstStudent Student = Student{"Max", "Muster"}
	var secondStudent Student = Student{"Franz", "Müller"}
	var thirdStudent Student = Student{"Fritz", "Kempf"}
	var fourthStudent Student = Student{"Fabian", "Schwander"}
	var fifthStudent Student = Student{"Daniel", "Lagger"}
	var sixthStudent Student = Student{"Ulrich", "Boll"}

	type Class struct {
		Students []Student
	}

	var firstClass = Class{[]Student{firstStudent, secondStudent, thirdStudent}}
	var secondClass = Class{[]Student{fourthStudent, fifthStudent, sixthStudent}}

	modules := map[uint][]Class{
		104: {firstClass},
		117: {firstClass, secondClass},
		346: {secondClass},
	}

	fmt.Println(modules)
}
