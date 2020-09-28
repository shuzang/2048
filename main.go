package main

import (
	"fmt"

	"github.com/shuzang/2048/game"
)

func main() {
	fmt.Println("Getting started!")
	g := game.New()
	for i := 0; i < 10; i++ {
		g.Display()
		g.AddElement()
		g.TakeInput()
		fmt.Println("new game")
	}
}
