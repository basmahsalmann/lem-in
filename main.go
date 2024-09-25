package main

import (
	"fmt"
	"os"
	f "lem-in/supportfiles"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid arguments: Enter File Name")
		os.Exit(1)
	}

	fileName := os.Args[1]

	f.ReadFile(fileName)

	f.Algorithm(fileName, f.GetNumberofAnts())

}