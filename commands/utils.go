package commands

import (
	"os"
	"fmt"
	"io/ioutil"
	"github.com/davecgh/go-spew/spew"
)

const maxDepth = 5


// PrintFile just reads the content of a file and prints it to console
func PrintFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	print(string(contents))
}

func Dump(element interface{}) {
	for j := 1; j < maxDepth; j++ {
		println()
		fmt.Println("----------- DEPTH ", j, "-----------------")
		spew.Config = spew.ConfigState{
			Indent: "\t",
			MaxDepth: j,
		}
		spew.Dump(element)
		fmt.Println("----------------------------")
		println()
		println("Press Enter to continue")
		if _, err := fmt.Scanf("\n"); err != nil {
			panic(err)
		}

	}
}
