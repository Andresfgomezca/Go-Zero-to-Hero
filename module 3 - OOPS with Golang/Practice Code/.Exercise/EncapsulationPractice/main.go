package main

import "fmt"

//NBA app to add players to the teams without exporting packages
type player struct {
	name, lastName, nickName                            string
	age, playerNumber, pointsPerGame, AssistancePerGame int
}
type team struct {
	players []player
}

func (t *team) addPlayer(p player) {
	t.players = append(t.players, p)
}
func (t *team) printPlayersInfo() {
	for i, p := range t.players {
		fmt.Println("index#", i, ": Player name: "+p.name+", Lastname: "+p.lastName+", Nickname: "+p.nickName+", Age: ", p.age, ", Player Number: ", p.playerNumber)
	}
}

func main() {
	player1 := player{
		name:              "Luka",
		lastName:          "Doncic",
		nickName:          "luka",
		age:               22,
		playerNumber:      77,
		pointsPerGame:     28,
		AssistancePerGame: 9,
	}
	dallasMaverics := team{}
	dallasMaverics.addPlayer(player1)
	dallasMaverics.printPlayersInfo()
}
