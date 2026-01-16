package main

import (
	"fmt"
	"math"
	"strings"
)

// Time:	O(n)
// Space:	O(n)
func intToRoman(num int) string {
	if num == 0 {
		return ""
	}
	symbols := map[int]rune{
		1:    'I',
		5:    'V',
		10:   'X',
		50:   'L',
		100:  'C',
		500:  'D',
		1000: 'M',
	}
	var res strings.Builder
	digits := numberOfDigits(num)
	for num != 0 {
		pow := int(math.Pow10(digits - 1))
		d := num / pow
		if d == 4 {
			res.WriteRune(symbols[1*pow])
			res.WriteRune(symbols[5*pow])
		} else if d == 9 {
			res.WriteRune(symbols[1*pow])
			res.WriteRune(symbols[10*pow])
		} else {
			for d != 0 {
				if d < 4 {
					res.WriteRune(symbols[1*pow])
					d--
				} else {
					res.WriteRune(symbols[5*pow])
					d -= 5
				}
			}
		}
		digits--
		num %= pow
	}
	return res.String()
}

func numberOfDigits(num int) int {
	length := 0
	for num != 0 {
		length++
		num /= 10
	}
	return length
}

func main() {
	n1 := 3749
	n2 := 58
	n3 := 1994

	fmt.Println(intToRoman(n1))
	fmt.Println(intToRoman(n2))
	fmt.Println(intToRoman(n3))
}
