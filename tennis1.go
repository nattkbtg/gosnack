package main

import (
	"fmt"
	"strings"
)

func calloutPoint(input_a string) string {
	s := strings.Split(input_a, "-")
	err := "invalid tennis's point input"
	output_a := ""

	if len(s) > 2 {
		return err
	} else {
		for i := 0; i < 2; i++ {

			switch s[i] {
			case "0":
				output_a = output_a + "Love"
			case "15":
				output_a = output_a + "Fifteen"
			case "30":
				output_a = output_a + "Thirty"
			case "40":
				output_a = output_a + "Forty"
			default:
				return err
			}
			if i == 0 {
				output_a = output_a + " - "
			}
		}
	}
	return output_a
}

func main() {

	input := "15-0"
	output := calloutPoint(input)
	fmt.Printf("input = " + input + ", output = " + output + "\n")

	input = "15-15"
	output = calloutPoint(input)
	fmt.Printf("input = " + input + ", output = " + output + "\n")

	input = "15-15-40"
	output = calloutPoint(input)
	fmt.Printf("input = " + input + ", output = " + output + "\n")

	input = "20-15"
	output = calloutPoint(input)
	fmt.Printf("input = " + input + ", output = " + output + "\n")

	input = "30-20"
	output = calloutPoint(input)
	fmt.Printf("input = " + input + ", output = " + output + "\n")

	input = "40-30"
	output = calloutPoint(input)
	fmt.Printf("input = " + input + ", output = " + output + "\n")

	input = "0-40"
	output = calloutPoint(input)
	fmt.Printf("input = " + input + ", output = " + output + "\n")

}
