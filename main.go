package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"github.com/zugzwang/parseopenpgp"
)

func main() {
	// Load input file
	if len(os.Args) < 2 {
		println("Must give a filename...")
		os.Exit(1)
	}
	filename := os.Args[1]
	fmt.Println("Filename: ", filename)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// Prompt user choice
	PrintFile("msg/welcome.txt")
	var choice int
	for {
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			panic(err)
		}
		if choice == 1 {
			parseopenpgp.ParseArmored(string(contents))
			break
		}
		println("Invalid choice...")
	}
}
