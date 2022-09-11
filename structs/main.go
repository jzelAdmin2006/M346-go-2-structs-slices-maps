package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	DayOfBirth   byte
	MonthOfBirth byte
	YearOfBirth  int16
}

type Person struct {
	Name FullName
	Born BirthDate
}

type Teacher struct {
	FullName
	BirthDate
	TeachesModule string
}

func main() {
	var myName FullName = FullName{
		FirstName: "Patrick",
		LastName:  "Bucher",
	}
	var myBirthDate BirthDate = BirthDate{
		DayOfBirth:   24,
		MonthOfBirth: 6,
		YearOfBirth:  1987,
	}
	var teacher = Person{
		Name: myName,
		Born: myBirthDate,
	}

	fmt.Println("Teacher's last name:", teacher.Name.LastName)

	teacher.Name.FirstName = "Padraigh"
	fmt.Println("Teacher's Irish name:", teacher.Name.FirstName)

	fmt.Println(myName, myBirthDate, teacher)

	pbucher := Teacher{
		myName,
		myBirthDate,
		"Modul 346",
	}
	fmt.Println("Teacher's full name:", pbucher.FirstName, pbucher.LastName)
	fmt.Printf("Teacher's birth date: %d.%d.%d\n", pbucher.DayOfBirth,
		pbucher.MonthOfBirth, pbucher.YearOfBirth)

	fmt.Printf("%v\n", pbucher)
	fmt.Printf("%q\n", pbucher)
}
