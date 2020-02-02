package main

import (
	"os"
	"io/ioutil"
)

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
