package main

import (
	"flag"
	"fmt"
)

type Animal struct {
	runSpeed int
	name     string
	voice    string
}

type Turtle struct {
	Animal        Animal
	age           int
	shellDiameter int
}

type Tiger struct {
	Animal Animal
	weight int
	food   string
}

type Fish struct {
	Animal       Animal
	finsQuantity int
	area         string
}

type Race struct {
	Distance int
	Turtle   Turtle
	Tiger    Tiger
	Fish     Fish
}

func main() {
	distance := flag.Int("distance", 0, "input distance")
	turtleSpeed := flag.Int("turtleSpeed", 0, "input turtle speed")
	tigerSpeed := flag.Int("tigerSpeed", 0, "input tiger speed")
	fishSpeed := flag.Int("fishSpeed", 0, "input fish speed")
	flag.Parse()

	var turtleAnimal = Animal{
		runSpeed: *turtleSpeed,
		name:     "Turtle",
		voice:    "I am turtle",
	}
	var turtle = Turtle{
		age:           300,
		shellDiameter: 50,
		Animal:        turtleAnimal,
	}
	var tigerAnimal = Animal{
		runSpeed: *tigerSpeed,
		name:     "Tiger",
		voice:    "I am tiger",
	}
	var tiger = Tiger{
		weight: 170,
		food:   "meat",
		Animal: tigerAnimal,
	}
	var fishAnimal = Animal{
		runSpeed: *fishSpeed,
		name:     "Fish",
		voice:    "I am fish",
	}
	var fish = Fish{
		finsQuantity: 5,
		area:         "sea",
		Animal:       fishAnimal,
	}
	var race = Race{Distance: *distance}

	race.createTeam(&turtle, &tiger, &fish)
	race.start()
}

func (a *Animal) winnerSay() {
	fmt.Println("I am winner!")
}

func (a *Animal) loserSay() {
	fmt.Println("I am looser")
}

func (r *Race) start() {
	var round = 0

	var finishedRace = 0
	for {
		var runners = []Animal{r.Turtle.Animal, r.Tiger.Animal, r.Fish.Animal}
		for _, v := range runners {
			if v.runSpeed*round >= 1000 && v.runSpeed*round-v.runSpeed < 1000 {
				switch finishedRace {
				case 0:
					fmt.Printf("%s - переміг. Час ітерацій становить %d \n", v.name, round)
					v.winnerSay()
					finishedRace++
				case 1:
					fmt.Printf("%s - другий. Час ітерацій становить %d \n", v.name, round)
					finishedRace++
				case 2:
					fmt.Printf("%s - програв. Час ітерацій становить %d \n", v.name, round)
					v.loserSay()
					finishedRace++
				}
			}
		}

		if finishedRace == 3 {
			break
		}

		round++
	}
}

func (r *Race) createTeam(turtle *Turtle, tiger *Tiger, fish *Fish) {
	r.Tiger = *tiger
	r.Turtle = *turtle
	r.Fish = *fish
}
