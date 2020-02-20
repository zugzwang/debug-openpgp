package commands

import (
	"os"
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

// DecryptCiphertext takes a ciphertext and a private key and decrypts.
func DecryptCiphertext() {
	// Load input files
	if len(os.Args) < 3 {
		println("Usage: debug_openpgp /path/to/key /path/to/ciphertext")
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

	fmt.Print("Proceeding with first key\n\n")
	key := entities[0].PrivateKey

	// Decrypt key if necessary
	passphrase := ""
	if key.Encrypted {
		passphrase = askForPassphrase()
		decryptKey(key, passphrase)
	}

	// Key ok. Import ciphertext
	fmt.Println("Key available, import ciphertext")
	enterToContinue()
	cipherFilename := os.Args[2]
	fmt.Println("Importing: ", cipherFilename)
	println()
	cipherFile, err := os.Open(cipherFilename)
	if err != nil {
		panic(err)
	}
	decodedCipher, err := armor.Decode(cipherFile)
	if err != nil {
		fmt.Println("Error decoding ciphertext: ", err.Error())
	}
	fmt.Println("Ciphertext imported ok:")
	singleDump(decodedCipher, 3)
	enterToContinue()

	// Define prompt in order to ReadMessage
	prompt := func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
		err := keys[0].PrivateKey.Decrypt([]byte(passphrase))
		if err != nil {
			fmt.Printf("Prompt: error decrypting key: %s", err)
			return nil, err
		}
		return nil, nil
	}

	// Read message
	messageDetails, err := openpgp.ReadMessage(decodedCipher.Body, entities, prompt, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Ciphertext decrypted ok:")
	singleDump(messageDetails, 3)

	fmt.Println("\n\nDecryption follows")
	enterToContinue()
	decrypted, err := ioutil.ReadAll(messageDetails.UnverifiedBody)
	if err != nil {
		fmt.Println("Error reading UnverifiedBody: ", err.Error())
		os.Exit(1)
	}
	fmt.Println(string(decrypted))
}
