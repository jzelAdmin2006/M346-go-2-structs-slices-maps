package main

import "fmt"

func main() {
	var name string    // a single string
	var names []string // a slice of strings
	fmt.Println(name, names)

	var days = []string{"Mo", "tu", "We", "Th", "Fr", "Sa"}
	days = append(days, "Su")
	days[1] = "Tu"
	fmt.Println(days)
	fmt.Println(len(days))
	workdays := days[0:5]
	weekend := days[5:7]
	fmt.Println(workdays, len(workdays))
	fmt.Println(weekend, len(weekend))
	firstDay := days[0]
	lastDay := days[len(days)-1]
	fmt.Println("from", firstDay, "to", lastDay)

	var numbers = make([]int, 0)
	var moreNumbers = make([]int, 3)
	fmt.Println(numbers, len(numbers), cap(numbers))
	fmt.Println(moreNumbers, len(moreNumbers), cap(moreNumbers))
	var extract = moreNumbers[0:2]
	fmt.Println(extract, len(extract), cap(extract))
}
