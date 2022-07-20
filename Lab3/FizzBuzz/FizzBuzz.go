/**
* Prints out 1 to 100.
* If the number is divisible by 3, print "Fizz"
* If the number is divisible by 5, print "Buzz"
* If the number is divisible by 3 and 5, print "FizzBuzz"
* @author ayang
* Created: September 14, 2021
 */

package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {
		// checking if it's divisible by 3 AND 5
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("FizzBuzz ")
		}
		// checking if it's divisible by 3 but not 5
		if i%3 == 0 && i%5 != 0 {
			fmt.Printf("Fizz ")
		}
		// checking if it's divisible by 5 but not 3
		if i%5 == 0 && i%3 != 0 {
			fmt.Printf("Buzz ")
		}
		// checking if it's not divisilbe by 3 and 5
		if i%3 != 0 && i%5 != 0 {
			fmt.Printf("%d ", i)
		}
	}
}
