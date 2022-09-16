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

type Profile struct {
	FullName         FullName
	BirthDate        BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		FullName:         FullName{"Johannes", "Zeller"},
		BirthDate:        BirthDate{1, 4, 2006},
		NumberOfSiblings: 2,
		ZodiacSign:       '\u2648',
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
