package main

import "fmt"

func main() {
	//var countryPopulation map[string]uint
	//var numbersSquareRoots map[int]float32
	//var numbersIsPrime map[int]bool
	//countryPopulation := make(map[string]uint, 0)
	//numbersSquareRoots := make(map[int]float32, 0)
	//numbersIsPrime := make(map[int]bool, 0)
	countryPopulation := map[string]uint{
		"AT": 8_917_000,
		"CH": 8_637_000,
		"DE": 83_240_000,
	}
	numbersSquareRoots := map[int]float32{
		1: 1.0,
		2: 1.41421356237,
		4: 2.0,
	}
	numbersIsPrime := map[int]bool{
		1: false,
		2: true,
		3: true,
		4: false,
	}
	countryPopulation["IT"] = 59_550_000
	numbersSquareRoots[16] = 4.0
	numbersIsPrime[13] = true
	fmt.Println(countryPopulation)
	fmt.Println(numbersSquareRoots)
	fmt.Println(numbersIsPrime)

	fmt.Println("Swiss Population:", countryPopulation["CH"])
	fmt.Println("Square Root of 16:", numbersSquareRoots[16])
	fmt.Println("Is 13 Prime?", numbersIsPrime[13])

	fmt.Println("French Population:", countryPopulation["FR"])

	frenchPopulation, ok := countryPopulation["FR"]
	fmt.Println("Value:", frenchPopulation)
	fmt.Println("Was stored in map?", ok)

	delete(countryPopulation, "FR")
	fmt.Println(countryPopulation)
}
