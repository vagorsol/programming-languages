/**
* A function that increments based on the value given.
* Author: ayang
* Created: September 26, 2021
 */
package main

import "fmt"

// this needs a better name
type incr func() int

// var count int = 0
// turns out global var in golang are on global scope. who would have guessed /s

/**
* A function that returns an incrementer function.
 */
func increment(incrBy int) incr {
	var count int = 0 // variable to store the incremented value (for each incrementor)
	incrementor := func() int {
		ret := count           // store count separtely so can increment but start from 0
		count = count + incrBy // do the incrementing by given number here
		return ret
	}
	return incrementor
}

/**
* A function to run a test that demonstrates: a) each instance of the function increment is
* independent of each other b) it's "unbounded" property (can increment forever).
* increment can have an unbounded (theoretically limiteless) number, but because of instruction ambiguity and for
* my own sake, I chose not to and assume what I have shown already in the code is good enough on that front.
 */
func incrTest(incrVal1, incrVal2, incrVal3 int) {
	byVal1 := increment(incrVal1)
	byVal2 := increment(incrVal2)
	byVal3 := increment(incrVal3)
	for {
		fmt.Printf("%s %v\n", "First Value:", byVal1())
		fmt.Printf("%s %v\n", "Secound Value:", byVal2())
		fmt.Printf("%s %v\n", "Third Value:", byVal3())
	}
}

func main() {
	//TODO: write testing function, script (or equivalent) in windows
	byThree := increment(3)
	bySeven := increment(7)

	fmt.Printf("%v\n", byThree()) // 0
	fmt.Printf("%v\n", bySeven()) // 0
	fmt.Printf("%v\n", byThree()) // 3
	fmt.Printf("%v\n", bySeven()) // 7
	fmt.Printf("%v\n", byThree()) // 6
	fmt.Printf("%v\n", bySeven()) // 14
	fmt.Printf("%v\n", byThree()) // 9
	fmt.Printf("%v\n", bySeven()) // 21

	fmt.Printf("%s\n", "Testing Function Output: ")
	incrTest(3, 7, 9)
}
