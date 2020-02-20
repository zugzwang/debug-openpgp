package commands

import (
	"fmt"
	"os"

	"golang.org/x/crypto/openpgp"
)

// DecryptPrivateKey takes a ciphertext and a private key and decrypts.
func DecryptPrivateKey() {
	// Load input files
	if len(os.Args) < 2 {
		println("Usage: debug_openpgp /path/to/key")
		os.Exit(1)
	}

	// Import string key(s)
	keyFilename := os.Args[1]
	fmt.Println("Importing Key: ", keyFilename)
	println()
	keyFile, err := os.Open(keyFilename)
	if err != nil {
		panic(err)
	}
	// Unarmor
	entities, err := openpgp.ReadArmoredKeyRing(keyFile)
	if err != nil {
		panic(err)
	}
	fmt.Println("Imported entities:")
	singleDump(entities, 3)
	fmt.Println("Proceeding with first key")
	key := entities[0].PrivateKey

	if !key.Encrypted {
		fmt.Println("Key is NOT encrypted.")
		os.Exit(0)
	} else {
		passphrase := askForPassphrase()
		decryptKey(key, passphrase)
		progressiveDump(key)
	}
}
