package main

import "fmt"

//	"strings"

type Game struct {
	X int
	Y int
}

func (g *Game) PlayerAGetPoint() {
	g.X++
}

func (g *Game) PlayerBGetPoint() {
	g.Y++
}

func (g *Game) CurrentScore() {
	//	fmt.Println(g.X, g.Y)
	output_a := ""
	switch g.X {
	case 0:
		output_a = output_a + "Love"
	case 1:
		output_a = output_a + "Fifteen"
	case 2:
		output_a = output_a + "Thirty"
	case 3:
		output_a = output_a + "Forty"
	}
	output_a = output_a + " - "
	switch g.Y {
	case 0:
		output_a = output_a + "Love"
	case 1:
		output_a = output_a + "Fifteen"
	case 2:
		output_a = output_a + "Thirty"
	case 3:
		output_a = output_a + "Forty"
	}
	if g.X == 3 {
		output_a = "A Wins"
	}
	if g.Y == 3 {
		output_a = "B Wins"
	}
	if g.X == 3 && g.Y == 3 {
		output_a = "Tied"
	}
	fmt.Println(output_a)
}

func main() {

	g := Game{}
	g.PlayerBGetPoint()
	g.PlayerBGetPoint()
	g.PlayerAGetPoint()
	g.PlayerBGetPoint()

	g.CurrentScore()

	gc := Game{}
	gc.PlayerBGetPoint()
	gc.PlayerBGetPoint()
	gc.CurrentScore()

	ga := Game{}
	ga.PlayerBGetPoint()
	ga.PlayerBGetPoint()
	ga.PlayerBGetPoint()
	ga.PlayerAGetPoint()
	ga.PlayerAGetPoint()
	ga.PlayerAGetPoint()
	ga.CurrentScore()
}
