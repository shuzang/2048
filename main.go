package main

import (
	"fmt"
	"time"

	"github.com/shuzang/2048/game"
)

func main() {
	fmt.Println("Use {W A S D} or Arrow keys to move the board")
	time.Sleep(time.Second * 2)
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
