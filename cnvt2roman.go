package main

import (
	"fmt"
)

func cnvt2Roman(in_a int) (out_a string) {

	div := 0
	a_num := [9]int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	b_string := [9]string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	i := 8

	for in_a > 0 {

		div = in_a / a_num[i]
		in_a = in_a % a_num[i]

		for div > 0 {
			out_a = out_a + b_string[i]
			div--
		}
		i--
	}

	return out_a
}

func main() {

	resRoman := ""

	for ai := 1; ai <= 100; ai++ {
		resRoman = cnvt2Roman(ai)
		fmt.Println("Decimal = ", ai, "Roman = ", resRoman)
	}

}
