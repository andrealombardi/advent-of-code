package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("day_08/1/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	data := make([]io, 0)
	for scanner.Scan() {
		elements := strings.Split(scanner.Text(), "|")
		inputs := strings.Fields(elements[0])
		outputs := strings.Fields(elements[1])
		data = append(data, io{inputs, outputs})
	}

	result := 0

	for _, d := range data {
		is8 := false
		for _, any := range d.i {
			if len(any) == 8 {
				is8 = true
			}
		}
		for _, any := range d.o {
			if len(any) == 8 {
				is8 = true
			}
		}
		if !is8 {
			fmt.Printf("DOES NOT")
		}
	}

	fmt.Printf("result: %d ", result)

}

func isOneFourSevenOrEight(s string) bool{
	return len(s) == 2 || //one
	    len(s) == 4 || //four
		len(s) == 3 || //seven
		len(s) == 7 //eight
}


type io struct {
	i []string
	o []string
}

//Digit	Display	p	a	b	c	d	e	f	g	pabcdefg	hex     pabcdefg	hex-pgfedcba
////0		0		on	on	on	on	on	on		01111110	0x7E	00111111	0x3F
////1		1			on	on					00110000	0x30	00000110	0x06
////2		2		on	on		on	on		on	01101101	0x6D	01011011	0x5B
////3		3		on	on	on	on			on	01111001	0x79	01001111	0x4F
////4		4			on	on			on	on	00110011	0x33	01100110	0x66
////5		5		on		on	on		on	on	01011011	0x5B	01101101	0x6D
////6		6		on		on	on	on	on	on	01011111	0x5F	01111101	0x7D
////7		7		on	on	on					01110000	0x70	00000111	0x07
////8		8		on	on	on	on	on	on	on	01111111	0x7F	01111111	0x7F
////9		9		on	on	on	on		on	on	01111011	0x7B	01101111	0x6F

//																			         compare with 7 gives me a
////1		1			on	on					00110000	0x30	00000110	0x06 compare with 3 (the only 5 digit with b and c)
////4		4			on	on			on	on	00110011	0x33	01100110	0x66
////7		7		on	on	on					01110000	0x70	00000111	0x07


////3		3		on	on	on	on			on	01111001	0x79	01001111	0x4F  compare with 4 gives me f

////6		6		on		on	on	on	on	on	01011111	0x5F	01111101	0x7D



////0		0		on	on	on	on	on	on		01111110	0x7E	00111111	0x3F
////2		2		on	on		on	on		on	01101101	0x6D	01011011	0x5B
////5		5		on		on	on		on	on	01011011	0x5B	01101101	0x6D
////9		9		on	on	on	on		on	on	01111011	0x7B	01101111	0x6F


//5  -> 5,2,3
//6  -> 6,9,0