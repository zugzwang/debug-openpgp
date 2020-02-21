package commands

import (
	"os"
	"fmt"
	"bytes"
	"io/ioutil"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

// EncryptDecrypt takes a key pair, tries to encrypt and then decrypt a
// user-queried message.
func EncryptDecrypt() {
	// Load input files
	if len(os.Args) < 3 {
		println("Usage: debug_openpgp /path/to/publickey /path/to/privatekey")
		os.Exit(1)
	}
	// Import string public key
	pkEntities := importPublicKey(os.Args[1])

	// Import string private key(s)
	skEntities := importSecretKey(os.Args[2])

	fmt.Println("Encryption follows")
	enterToContinue()

	// Encrypt message to pk
	var buf bytes.Buffer
	messageWriter, err := openpgp.EncryptText(&buf, pkEntities, nil, nil, nil)
	message := askForMessage()
	_, err = messageWriter.Write([]byte(message))
	if err != nil {
		panic(err)
	}
	err = messageWriter.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println("Encryption Ok:")
	singleDump(messageWriter, 4)

	fmt.Println("Armored ciphertext follows")
	enterToContinue()
	var b bytes.Buffer
	w, err := armor.Encode(&b, "PGP MESSAGE", nil)
	if err != nil {
		panic(err)
	}
	if _, err = w.Write(buf.Bytes()); err != nil {
		panic(err)
	}
	if err := w.Close(); err != nil {
		panic(err)
	}
	fmt.Println(b.String())

	// Define prompt in order to ReadMessage
	prompt := func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
		passphrase := askForPassphrase()
		err := keys[0].PrivateKey.Decrypt([]byte(passphrase))
		if err != nil {
			fmt.Printf("Prompt: error decrypting key: %s\n", err)
			return nil, err
		}
		return nil, nil
	}

	// Read message
	messageDetails, err := openpgp.ReadMessage(&buf, skEntities, prompt, nil)
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
