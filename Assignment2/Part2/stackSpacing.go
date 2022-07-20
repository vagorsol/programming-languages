/**
* A program to test the maximum height the stack in Go can go,
* based on the size of the item it is stacking
* @author ayang
* Created: September 25, 2021
* How high the stack went:
* 	-- basicRecurse: 6100738
*	-- moreRecurse: 4793437 4793437
*	-- mostRecurse: 3947536 3947536 3947536
*	-- int16Recurse: 5890
* 	-- int64Recurse: 6100738
*	-- float64Recurse: 6100720.000000
 */

package main

import "fmt"

/**
* A basic recursion function, keeps counting up by increments of 1 until
* Go's stack runs full.
 */
func basicRecurse(i int) {
	fmt.Printf("%d\n", i)
	basicRecurse(i + 1)
}

/**
* A modified version of basicRecurse, passing in 2 variables instead of 1 for increased size.
 */
func moreRecurse(i int, j int) {
	fmt.Printf("%d %d\n", i, j)
	moreRecurse(i+1, j+1)
}

/**
* A modified version of basicRecurse, passing in 3 variables instead of 1 for increased size.
 */
func mostRecurse(i int, j int, k int) {
	fmt.Printf("%d %d %d\n", i, j, k)
	mostRecurse(i+1, j+1, k+1)
}

/**
* A variation of basicRecurse where the type is int16 instead of int
 */
func int16Recurse(i int16) {
	fmt.Printf("%d\n", i)
	int16Recurse(i + 1)
}

/**
* A variation of basicRecurse where the type is int64 instead of int.
* Meant as a comparison to float64.
 */
func int64Recurse(i int64) {
	fmt.Printf("%d\n", i)
	int64Recurse(i + 1)
}

/**
* A variation of basicRecurse where the type is float64 instead of int.
* Increments by 1 instead of some small decimal value because a) I don't know
* what to make that decimal value and b) so I can compare the results with int64
 */
func float64Recurse(i float64) {
	fmt.Printf("%f\n", i)
	float64Recurse(i + 1)
}

func main() {
	basicRecurse(0)
	// moreRecurse(0, 0)
	// mostRecurse(0, 0, 0)
	// int16Recurse(0)
	// int64Recurse(0)
	// float64Recurse(0)
}
