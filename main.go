package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
	"github.com/shuzang/2048/game"
)

func main() {
	fmt.Println("Use {W A S D} or Arrow keys to move the board")
	fmt.Printf("Press any key to start\n")
	_, _, err := keyboard.GetSingleKey()
	if err != nil {
		log.Fatalln("error while taking input to start the game")
	}
	b := game.NewBoard()
	b.AddElement()
	b.AddElement()
	for true {
		if b.IsOver() {
			break
		}
		b.AddElement()
		b.Display()
		res := b.TakeInput()
		if !res {
			return
		}

	}
	fmt.Println("\n**********  game over  **********")
	max, total := b.CountScore()
	fmt.Printf("Max Score: %v \n", max)
	fmt.Printf("Total Score %v \n", total)
}
