package main

import (
	"fmt"

	"github.com/zugzwang/debug-openpgp/commands"
)

func main() {

	// Prompt user choice
	println("1. Parse packets of armored string")
	println("2. Decrypt a private key from armored string")

	print("\nYour choice: ")
	var choice int
	if _, err := fmt.Scanf("%d", &choice); err != nil {
		panic(err)
	}
	switch choice {
	case 1:
		commands.ParseArmored()
	case 2:
		commands.DecryptPrivateKey()
	default:
	println("Invalid choice...")
	}
}
