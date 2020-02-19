package commands

import (
	"io"
	"os"
	"fmt"
	"strings"
	"io/ioutil"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

// ParseArmored dumps the packets of the given armored string
func ParseArmored() {
	// Load input file
	if len(os.Args) < 2 {
		println("Must give a filename...")
		os.Exit(1)
	}
	filename := os.Args[1]
	fmt.Println("Filename: ", filename)
	println()
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	input := string(contents)
	// Unarmor ciphertext
	sr := strings.NewReader(input)
	data, err := armor.Decode(sr)
	if err != io.EOF && err != nil {
		println("Could not unarmor ciphertext!")
		println(err.Error())
	}

	// Parse packets
	packets := packet.NewReader(data.Body)
	parsedPackets := make([]packet.Packet, 0)
	errs := make([]error, 0)

	for {
		var p packet.Packet
		p, err = packets.Next();
		if err == io.EOF || err != nil {
			break
		}
		parsedPackets = append(parsedPackets, p)
		errs = append(errs, err)
	}
	// Print details of each packet
	for {
		fmt.Println("Parsed", len(parsedPackets), "packets:")
		for i, p := range parsedPackets {
			fmt.Printf("%d: %T (error: %v)\n", i, p, errs[i])
		}
		var choice int
		fmt.Println("Choose packet to print info:")
		_, err = fmt.Scanf("%d", &choice)
		Dump(parsedPackets[choice])
	}
}
