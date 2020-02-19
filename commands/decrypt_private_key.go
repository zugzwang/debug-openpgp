package commands

import (
	"os"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/crypto/openpgp"
)


// DecryptPrivateKey takes a ciphertext and a private key and decrypts.
func DecryptPrivateKey() {
	// Load input files
	if len(os.Args) < 2 {
		println(
			"Usage: debug_openpgp /path/to/key")
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
		spew.Config = spew.ConfigState{
			MaxDepth: 3,
			Indent: "\t",
		}

		fmt.Println("Got this:")
		spew.Dump(entities)
		fmt.Println("Proceeding with first key")
		key := entities[0].PrivateKey

		if !key.Encrypted {
			fmt.Println("Key is NOT encrypted.")
			os.Exit(0)
		} else {

			fmt.Print("Private key is passphrase encrypted. Enter text: ")
			var passphrase string
			if _, err := fmt.Scanf("%s\n", &passphrase); err != nil {
				panic(err)
			}
			err = key.Decrypt([]byte(passphrase))
			if err != nil {
				fmt.Printf("prompt: error decrypting key: %s\n", err.Error())
				fmt.Printf("Also, maybe the key is bcrypted?")
				os.Exit(1)
			} else {
				fmt.Println("Private key decrypted correctly. Got the following:")
			}
			Dump(key)
		}
	}
