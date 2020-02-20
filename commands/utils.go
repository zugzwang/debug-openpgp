package commands

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davecgh/go-spew/spew"
	"golang.org/x/crypto/openpgp/packet"
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

func askForPassphrase() (passphrase string) {
	fmt.Print("Private key is passphrase encrypted. Enter text:\n")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		passphrase = scanner.Text()
		fmt.Printf("Input was: %q\n", passphrase)
	}
	return
}

func decryptKey(key *packet.PrivateKey, passphrase string) {
	err := key.Decrypt([]byte(passphrase))
	if err != nil {
		fmt.Printf("Error decrypting key: %s\n\n", err.Error())
		fmt.Printf("Maybe wrong password?\n")
		fmt.Printf("Maybe the key needs to be bcrypted?\n")
		os.Exit(1)
	} else {
		fmt.Println("Private key decrypted correctly")
		singleDump(key, 3)
	}
}

func progressiveDump(element interface{}) {
	println("****** BEGIN PROGRESSIVE DUMP ******")
	for j := 1; j < maxDepth; j++ {
		println()
		fmt.Println("----------- DEPTH ", j, "-----------------")
		spew.Config = spew.ConfigState{
			Indent:   "\t",
			MaxDepth: j,
		}
		spew.Dump(element)
		fmt.Println("----------------------------")
		println()
		enterToContinue()
	}
	println("****** END PROGRESSIVE DUMP *******")
}

func singleDump(element interface{}, depth int) {
	println("****** BEGIN DUMP ******")
	spew.Config = spew.ConfigState{
		Indent:   "\t",
		MaxDepth: depth,
	}
	spew.Dump(element)
	println("****** END DUMP *******")
}

func enterToContinue() {
	println("Press Enter to continue")
	if _, err := fmt.Scanln(); err != nil {
		panic(err)
	}
}
