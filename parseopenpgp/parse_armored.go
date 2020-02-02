package parseopenpgp

import (
	"io"
	"fmt"
	"strings"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

const maxDepth = 6

// ParseArmored dumps the packets of the given armored string
func ParseArmored(input string) {
	spew.Config = spew.ConfigState{
		Indent: "\t",
	}
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
		parsedPackets = append(parsedPackets, p)
		errs = append(errs, err)
		if err == io.EOF || err != nil {
			break
		}
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
		for j := 1; j < maxDepth; j++ {
			fmt.Println("----------- DEPTH ", j, "-----------------")
			spew.Config = spew.ConfigState{
				Indent: "\t",
				MaxDepth: j,
			}
			spew.Dump(parsedPackets[choice])
			fmt.Println("----------------------------")
			println()
			println("Press Enter to continue")
			_, err = fmt.Scanf("Continue")
		}
	}
}
