package main

import (
	"fmt"

	"github.com/zugzwang/debug-openpgp/commands"
)

func main() {

	options := `
	1. Parse packets of armored string
	2. Decrypt a private key from armored string
	3. Decrypt ciphertext
	4. Encrypt + decrypt roundtrip
	`
	println(options)

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
	case 3:
		commands.DecryptCiphertext()
	case 4:
		commands.EncryptDecrypt()
	default:
		println("Invalid choice...")
	}
}
