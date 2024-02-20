package main

import (
	"fmt"
	"log"

	"github.com/xavier2910/tundra"
)

func main() {
	fmt.Println("The Tundra take 2 version 0.0.0. Nothing implemented yet.")

	err := run()

	if err != nil {
		log.Fatal(err)
	}
}

func run() error {

	universe := tundra.NewWorld(
		tundra.NewPlayer(),
		[]*tundra.Location{},
	)

	fmt.Println(universe)

	return nil
}
