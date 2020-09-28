package main

import (
	"fmt"

	"github.com/shuzang/2048/game"
)

func main() {
	fmt.Println("Getting started!")
	g := game.New()
	g.Display()
}
