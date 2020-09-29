package main

import (
	"fmt"

	"github.com/shuzang/2048/game"
)

func main() {
	fmt.Println("Getting started!")
	g := game.New()
	g.AddElement()
	g.AddElement()
	for true {
		g.AddElement()
		g.Display()
		g.TakeInput()
		//g.Display()
		fmt.Println("new game")
	}
}
