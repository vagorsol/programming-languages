/**
* Takes user input, converts from kilometers and meters,
* and vice versa, and returns the conversion
* @author ayang
* Created: September 14, 2021
 */

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var dist float64
	var err error
	dist, err = strconv.ParseFloat(os.Args[1], 64)
	var unit = os.Args[2]
	// the values to be outputted
	var convDist float64
	var convUnit string

	// if the user did not enter a value, remind them and stop (hopefully)
	if err != nil {
		fmt.Printf("Please enter a value")
		return
	}

	if unit == "m" || unit == "M" {
		// if unit == m, m -> k convert
		convUnit = "k"
		convDist = dist * 1.609
	} else { // unit == "k" || unit == "K"
		// if unit == k, k -> m convert
		convUnit = "m"
		convDist = dist / 1.609
	}
	//return the converted dist
	fmt.Printf("%f %s\n", convDist, convUnit)
}
