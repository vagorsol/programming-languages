/**
* A set of structures intended to store information about flights.
* Featuring experimentations with structures.
* @author ayang
* Created: September 17, 2021
 */
package main

import "fmt"

// a structure to hold the date information of a flight
type Date struct {
	year  int
	month int
	day   int
}

// a structure that holds all the information about a flight
type Flight struct {
	date        Date
	deptime     int    // deptime: 24 hrs clock without colon (ex: 2355 == 11:55pm)
	arrtime     int    // same as arrtime
	carrier     string // airline ID
	flightnum   int
	arrdelay    int    // negative number indicates early arrival
	depdelay    int    // negative number indicates early departure
	origin      string // if !PHL, then is an arrival
	destination string // if !PHL, is a departure
	distance    int    // miles
	cancelled   bool   // true iff flight was cancelled
}

// String() method for the struct Date
func (d Date) String() string {
	return fmt.Sprintf("%d %d %d", d.year, d.month, d.day)
}

// String() method for the struct Flight
func (f Flight) String() string {
	return fmt.Sprintf("%s %d %d %s %d %d %d %s %s %d %t", (f.date).String(), f.deptime, f.arrtime, f.carrier, f.flightnum, f.arrdelay, f.depdelay,
		f.origin, f.destination, f.distance, f.cancelled)
}

func main() {
	// defining flights for Part 1
	var flightDate = Date{2008, 1, 3}
	var a = Flight{flightDate, 1734, 1941, "WN", 23, 36, 44, "JAX", "PHL", 742, false}
	var b = Flight{flightDate, 712, 926, "WN", 1232, 11, 12, "JAX", "PHL", 742, false}
	var c = Flight{flightDate, 1127, 1856, "WN", 1285, 21, 42, "LAS", "PHL", 2176, true}
	d := c
	d.cancelled = !c.cancelled

	// defining additional flights for Part 2
	var flightDate2 = Date{2008, 12, 12}
	var flightDate3 = Date{2008, 12, 13}
	var flightDate4 = Date{2008, 12, 5}
	var flightDate5 = Date{2008, 12, 6}
	var flightDate6 = Date{2008, 12, 7}
	var flightDate7 = Date{2008, 12, 22}

	var e = Flight{flightDate2, 1834, 2149, "DL", 1687, 86, 69, "PHL", "SLC", 1926, true}
	var f = Flight{flightDate2, 944, 1555, "DL", 1776, -10, -1, "SLC", "PHL", 1926, true}
	var g = Flight{flightDate3, 1221, 1413, "DL", 692, -9, -4, "ATL", "PHL", 665, true}
	var h = Flight{flightDate4, 1523, 1804, "AA", 1768, -21, 3, "ORD", "PHL", 678, true}
	var i = Flight{flightDate5, 1628, 1933, "AA", 1768, -7, -2, "ORD", "PHL", 678, true}
	var j = Flight{flightDate6, 1520, 1804, "AA", 1768, -21, 0, "ORD", "PHL", 678, true}
	var k = Flight{flightDate7, 545, 839, "CO", 1677, 19, 10, "PHL", "IAH", 1324, true}
	var l = Flight{flightDate7, 1920, 2314, "CO", 1776, -16, 0, "IAH", "PHL", 1324, true}
	var m = Flight{flightDate7, 723, 952, "CO", 1777, 3, 18, "PHL", "IAH", 1324, true}

	// a slice that holds all of the flights
	var aslice = []Flight{a, b, c, e, f, g, h, i, j, k, l, m}

	// subslice of the first 5 elements of aslice
	// var subAslice []Flight
	var subAslice = aslice[:5] // make subAslice the first 5 components of aslice
	// fmt.Printf("%v\n", subAslice)

	// subslice of the last 5 elements of aslice
	// var subBslice []Flight
	var subBslice = aslice[len(aslice)-5:]
	fmt.Printf("%v\n", subBslice)

	// subslice of the first 3 and last 2 elements of aslice
	var subCslice = aslice[:3]
	var subDslice = aslice[len(aslice)-2:]
	subEslice := append(subCslice, subDslice...)
	fmt.Printf("%v\n", subEslice)

	// Question 2: equality of slice and it's subslice
	fmt.Printf("%v", aslice[0])
	subAslice[0].cancelled = true
	fmt.Printf("%v", subAslice[0])
	fmt.Printf("%v", aslice[0])
	// checking equality
	if aslice[0] == subAslice[0] {
		fmt.Println("\nequal!")
	} else {
		fmt.Println("\nunequal!")
	}

	// checking to see if any of flights A, B, and C are equal to each other
	if a == b {
		fmt.Printf("Flights A and B are the same!")
	} else if a == c {
		fmt.Printf("Flights A and C are the same!")
	} else if b == c {
		fmt.Printf("Flights B and C are the same!")
	} else {
		fmt.Printf("None of these flights are the same...")
	}

	// checking to see if C and D are equal
	fmt.Println(c)
	fmt.Println(d)
	if c == d {
		fmt.Printf("Flights C and D are the same!")
	} else {
		fmt.Printf("Flight C and D are not the same..\n")
	}

	// testing the String() functions
	// fmt.Println(c)
	// fmt.Println(Flight(c))
	// fmt.Printf("%+v", c)
	// fmt.Println(flightDate)
}
