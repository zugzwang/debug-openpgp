package commands

import (
	"fmt"
	"os"
)

// DecryptPrivateKey takes a ciphertext and a private key and decrypts.
func DecryptPrivateKey() {
	// Load input files
	if len(os.Args) < 2 {
		println("Usage: debug_openpgp /path/to/key")
		os.Exit(1)
	}

	// Import string key(s)
	entities := importSecretKey(os.Args[1])

	fmt.Println("Proceeding with first key")
	key := entities[0].PrivateKey
	if key == nil {
		fmt.Println("No private key found in first key of entities")
		os.Exit(1)
	}

	if !key.Encrypted {
		fmt.Println("Key is NOT encrypted.")
		os.Exit(0)
	} else {
		passphrase := askForPassphrase()
		decryptKey(key, passphrase)
		progressiveDump(key)
	}
}
