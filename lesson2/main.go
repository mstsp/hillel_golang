package main

import "fmt"

type Game struct {
	TotalPlayers int
	Genre        string
	Difficulty   string
	Brand        string
	Theme        string
}

type TableGame struct {
	game Game
	Name string
}

func main() {

	var dixit = Game{TotalPlayers: 8, Genre: "Creative", Difficulty: "Easy", Brand: "Libellud", Theme: "Fantasy"}

	var tableGameDixit = TableGame{game: dixit, Name: "Dixit"}

	games := make([]TableGame, 0, 3)

	games = append(games, tableGameDixit)

	var mysterium = Game{TotalPlayers: 7, Genre: "Mystery", Difficulty: "Middle", Brand: "Libellud", Theme: "Mystery"}
	var trekkingTheNationalParks = Game{TotalPlayers: 5, Genre: "Strategy", Difficulty: "Easy", Brand: "Underdog Games", Theme: "Mystery"}

	var tableGameMysterium = TableGame{game: mysterium, Name: "Mysterium"}
	var tableGametrekkingTheNationalParks = TableGame{game: trekkingTheNationalParks, Name: "Trekking The National Parks"}

	games = append(games, tableGameMysterium)
	games = append(games, tableGametrekkingTheNationalParks)

	gamesMap := make(map[int]TableGame)

	for i := 0; i < len(games); i++ {
		gamesMap[i] = games[i]
	}

	totalPlayers := 0
	for _, v := range gamesMap {
		fmt.Println(v)
		totalPlayers = totalPlayers + v.game.TotalPlayers
	}

	fmt.Println(totalPlayers)
}
