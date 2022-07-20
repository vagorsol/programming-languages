/**
* A function that can sort a struct by integers or by floats.
* Probably not done in a very efficient or elegant way.
* author @ayang
* Created: September 21, 2021
* Additional Note: I don't know what/how to "cast your slice from the
* original type to then new type" means. So I just made another slice instead.
 */

package main

import (
	"fmt"
	"sort"
)

type S struct {
	Number  int
	Decimal float64
	Words   string
}

// slices to store the instances of S struct
type IntSlice []S
type Float64Slice []S

// return the length of the slice for IntSlice
func (p IntSlice) Len() int {
	return len(p)
}

// return the length of the slice for Float64Slice
func (p Float64Slice) Len() int {
	return len(p)
}

// swap two values in the slice for IntSlice
func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// swap the values in the slice for Float64Slice
func (p Float64Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// check to see if the first int is less than the second one
func (p IntSlice) Less(i, j int) bool {
	return p[i].Number < p[j].Number
}

// check to see if the first float is less than the second one
func (p Float64Slice) Less(i, j int) bool {
	return p[i].Decimal < p[j].Decimal
}

func main() {
	a := S{1, 2.2, "help"}
	b := S{5, 6.7, "me"}
	c := S{4, 2.3, "please"}
	d := S{2, 4.5, "oh"}
	e := S{6, 7.8, "god"}
	f := S{3, 6.9, "i"}
	g := S{8, 9.8, "don't"}
	h := S{7, 1.4, "know"}
	i := S{0, 8.2, "anything"}
	j := S{9, 3.4, "aaaaaaaaaaaaaa"}

	// slice to be sorted by integer value
	aa := IntSlice{a, b, c, d, e, f, g, h, i, j}
	// slice to be sorted by float64 value
	bb := Float64Slice{a, b, c, d, e, f, g, h, i, j}

	// sort the structures by their int values (in descending order)
	sort.Slice(aa, func(i, j int) bool {
		return aa[i].Number > aa[j].Number
	})
	fmt.Println(aa)

	// sort the structures by their float64 values (in ascending order)
	sort.Slice(bb, func(i, j int) bool {
		return bb[i].Decimal > bb[j].Decimal
	})
	fmt.Println(bb)
}
