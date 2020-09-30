package main

import (
	"fmt"
	"log"

	"github.com/shuzang/2048/game"
)

func main() {
	fmt.Println("Use {W A S D} or Arrow keys to move the board")
	fmt.Printf("Press and key to start\n")
	_, err := game.GetCharKeystroke()
	if err != nil {
		log.Fatalln("error while taking input to start the game")
	}
	g := game.New()

	g.AddElement()
	g.AddElement()
	for true {
		if g.IsOver() {
			break
		}
		g.AddElement()
		g.Display()
		g.TakeInput()
	}
	fmt.Printf("******** Game Over ********\n")
	max, total := g.CountScore()
	fmt.Printf("Score: Max Tile Value:    %v \n", max)
	fmt.Printf("Score: Total Tiles Value: %v \n", total)
}
