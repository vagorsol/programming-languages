/**
* A test meant to test and demonstrate the kind of scope (static) Golang has.
* @author ayang
* Created: September 25, 2021
 */

package main

import "fmt"

/**
* A function to test whether the value of a variable will be based on the last
* value of that variable in stack
* e.g., if dynamic, both the inner and outer values of the variable "b" should be
* equal to 7 because "b := 7" will be the most recently called value of "b" on the stack.
* Else, 7 and 5 respectively.
 */
func inOutValTest() {
	// value of b in the "outer" function (for the "outer" scope)
	b := 5
	{
		// value of b in the "inner" function (for the inner scope)
		b := 7
		fmt.Printf("%s %d\n", "Inner Variable Value:", b)
	}
	fmt.Printf("%s %d\n", "Outer Variable Value:", b)
}

/**
* I wanted to also test calling variables outside of function scope in Golang,
* ideally using a throw/catch thing. I don't know how to do it, I don't think it's necessary to show
* that, and I don't have a good function name for it anyways.
func iNeedAGoodFuncName() {
	// a := 2
	fmt.Printf("%s %d\n", "Value of Variable that's Not of this Function's Scope: ", b)
}
*/

func main() {
	// run the scope test here
	inOutValTest()
}
