package main

import "fmt"

func main() {
	type Player struct {
		FirstName string
		LastName  string
	}
	type Team map[byte]Player
	teamA := Team{
		1: Player{
			FirstName: "Joe",
			LastName:  "Doe",
		},
		2: Player{
			FirstName: "Jay",
			LastName:  "Day",
		},
	}
	teamB := Team{
		1: Player{
			FirstName: "Jim",
			LastName:  "Jam",
		},
		2: Player{
			FirstName: "Jam",
			LastName:  "Bam",
		},
	}
	teams := []Team{
		teamA,
		teamB,
	}
	fmt.Println(teams)
}
