/**
* Testing for slice equality in Go
* @author ayang
* Created: Nov 15, 2021
 */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// make a slice of size 10
	var aa = make([]interface{}, 10)

	// for the first five indexes, put a number
	for i := 0; i < 5; i++ {
		aa[i] = i
	}

	// for index 5 and 6, put in something that clearly ISN'T an int(64)
	aa[5] = int32(32)
	aa[6] = "asdf"

	// make a slice of size 5, and put a number in each index
	var bb = make([]interface{}, 5)
	for i := 0; i < 5; i++ {
		bb[i] = i
	}

	// put in index 7 of aa the entirety of bb
	aa[7] = bb

	// for the last two indexes, uh. char + int. i don't really know what's going on here but end result is an int8!!
	for i := 8; i < 10; i++ {
		aa[i] = int8('a' + i)
	}

	// print out all of aa's contents
	fmt.Printf("%v\n", aa)

	// for all of aa, test if each index is a slice or not
	for _, v := range aa {
		fmt.Printf("%v\n", sliceTest(v))
	}

	fmt.Println()

	var testSlice = make([]interface{}, 10)
	copy(testSlice, aa)
	fmt.Println(sliceEquality(aa, testSlice)) // true
	fmt.Println(sliceEquality(aa, bb))        // false
}

// return true if the thing is a slice
func sliceTest(v interface{}) bool {
	rt := reflect.TypeOf(v)
	return rt.Kind() == reflect.Slice
}

// convert the thing passed to a slice
func sliceAssert(aa interface{}) []interface{} {
	return aa.([]interface{})
}

/**
* Checks to see if two slices are equal to each other (i.e., each element of
* the two slices are equal to each other)
 */
func sliceEquality(sliceA []interface{}, sliceB []interface{}) bool {
	// if the slices aren't the same length, then they can't be equal to each other
	if len(sliceA) != len(sliceB) {
		return false
	}

	// a recursive test for if the two slices are equal or not. inner function because no making people put count when declaring cause uh.
	// "encapsulation" or something like that.
	var equalityChecker func(sliceOne []interface{}, sliceTwo []interface{}, count int) bool

	equalityChecker = func(sliceOne []interface{}, sliceTwo []interface{}, count int) bool {
		// if it made it all the way through the func, the two slices are equal
		if count == len(sliceOne) {
			return true
		}

		// if the element is a slice itself,
		if sliceTest(sliceOne[count]) {
			// if slice2 doesn't have a slice at that position, they are not equal
			if !sliceTest(sliceTwo[count]) {
				return false
			}
			// run a comparison again
			debatablyEqual := equalityChecker(sliceAssert(sliceOne[count]), sliceAssert(sliceTwo[count]), 0)

			// if the inner slices are not equal, then the two outer slices are not equal
			if !debatablyEqual {
				return false
			}
			count = count + 1
		} else if sliceA[count] != sliceB[count] {
			// making this an else if because otherwise it would run for the ones that have inner slices
			// and that's No Good
			return false
		}

		// actually run the func now
		return equalityChecker(sliceOne, sliceTwo, count+1)
	}

	// run and return the equality checker
	return equalityChecker(sliceA, sliceB, 0)
}
