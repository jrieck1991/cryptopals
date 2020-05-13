package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "1c0111001f010100061a024b53535009181c"
	s2 := "686974207468652062756c6c277320657965"

	fmt.Println(XORHexStrings(s1, s2))
}

// ascii characters are 1 byte
// hex chars are 4 bits
// return hex encoded byte string
func ASCIIToHex(input string) []byte {

	if len(input)%2 != 0 {
		return nil
	}

	var bs []byte

	// step by 2
	for i := 0; i < len(input)-1; i += 2 {

		// read 2 bytes
		nib1 := input[i]
		nib2 := input[i+1]

		// for hex chars A-F subtract by decimal 55/hex 7
		if nib1 >= 'A' {

			// get 0<DIGIT>
			nib1 = nib1 - '7'
		}

		// shift bits left, leaving the last 4 moved up
		nib1 = nib1 << 4

		if nib2 >= 'A' {

			nib2 = nib2 - '7'
		}

		// get last 4 bits
		nib2 = nib2 & 0x0F

		// combine nibbles to get full byte
		fullByte := nib1 | nib2

		bs = append(bs, fullByte)
	}

	return bs
}

// expects hex strings
// strings must be of equal length
// return hex
func XORHexStrings(s1, s2 string) string {

	// return empty if strings arent same len
	if len(s1) != len(s2) {
		return ""
	}

	// get byte strings of hex strings
	s1bs := ASCIIToHex(s1)
	s2bs := ASCIIToHex(s2)

	// XOR bytes
	var result []byte
	for i := 0; i < len(s1bs); i++ {

		x := s1bs[i] ^ s2bs[i]

		result = append(result, x)
	}

	// convert hex byte string to ascii
	return HexBytesToASCII(result)
}

//
func HexBytesToASCII(hexBytes []byte) string {

	var hexChars []string
	for _, b := range hexBytes {
		hexChars = append(hexChars, fmt.Sprintf("%x", b))
	}

	return strings.Join(hexChars, "")
}
