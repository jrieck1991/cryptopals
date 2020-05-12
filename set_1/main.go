package main

import (
	"fmt"
	"strings"
)

func main() {

	hexStr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	b := ASCIIToHex(hexStr)
	fmt.Println(string(b))
	x := ByteStringToBase64(b)
	fmt.Println(x)
}

// convert ascii chars to hex bytes
func ASCIIToHex(s string) []byte {
	s = strings.ToUpper(s)

	var bs []byte

	// skip by 2
	for i := 0; i < len(s)-1; i += 2 {

		nib1 := s[i]
		nib2 := s[i+1]

		// if hex digit is >= A we need to subtract
		if nib1 >= 'A' {

			// subtract to get correct hex digit, A == 0A
			nib1 = nib1 - '7'
		}

		// for first char, shift bits left
		nib1 = nib1 << 4

		if nib2 >= 'A' {

			// subtract
			nib2 = nib2 - '7'
		}

		// get last 4 bits of second char
		nib2 = nib2 & 0x0F

		// combine into 1 byte
		b := nib1 | nib2

		// append byte to slice
		bs = append(bs, b)
	}

	return bs
}

// read hex bytes in groups of 3
// get first 6 bits of byte 1
// add 2 bits of byte 1 and first 4 bits of byte 2
// add last 4 bits of byte 2 and first 2 bits of byte 3
// get last 6 bits of byte 3
func ByteStringToBase64(b []byte) string {

	// lookup table
	base64LookupTable := []string{
		"A", "B", "C", "D", "E",
		"F", "G", "H", "I", "J", "K",
		"L", "M", "N", "O", "P", "Q",
		"R", "S", "T", "U", "V", "W", "X",
		"Y", "Z", "a", "b", "c", "d", "e", "f",
		"g", "h", "i", "j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u", "v",
		"w", "x", "y", "z", "0", "1", "2", "3",
		"4", "5", "6", "7", "8", "9", "+", "/",
	}

	// for remainder
	padding := "="

	var result []string

	for i := 0; i < len(b)-1; i += 3 {

		// copy first byte
		cb1 := b[i]

		// grab first 6 bits
		c1 := b[i] >> 2

		// get last 2 bits of byte 1
		l := cb1 & 0x03
		l = l << 4

		if i+1 >= len(b)-1 {
			result = append(result, base64LookupTable[c1], padding, padding, padding)
			continue
		}

		// copy byte 2
		cb2 := b[i+1]

		// get first 4 bits of byte 2
		x := b[i+1] >> 4

		// combine last 2 of byte 1 and first 4 of byte 2
		c2 := l | x

		// get last 4 bits of byte 2
		q := cb2 & 0x0F
		q = q << 2

		if i+2 > len(b)-1 {
			result = append(result, base64LookupTable[c1], base64LookupTable[c2], padding, padding)
			continue
		}

		// copy byte 3
		cb3 := b[i+2]

		// get first 2 bits of byte 3
		r := b[i+2] >> 6

		// combine 4 of byte 2 and first 2 of byte 3
		c3 := q | r

		// get last 6 bits of byte 3
		c4 := cb3 & 0x3F

		// append string
		result = append(result, base64LookupTable[c1], base64LookupTable[c2], base64LookupTable[c3], base64LookupTable[c4])
	}

	return strings.Join(result, "")
}
