/**
* A program that computs how long it will take someone to get a given number
* of pictures from candy wrapers.
* Author: ayang
* Created: September 27, 2021
 */

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/**
* A function that tests how long it takes to get all unique pictures from
* a given number of total unique pictures.
* Something about the loop/slice logic is wrong. I can't be bothered to spend more time
* trying to fix it right now.
 */
func collect(cardNum int) int {
	cardCollection := make([]int, cardNum) // a slice to store all the cards you've aquired
	sum := 0                               // number of cards you have to purchase
	var hasAll bool = true                 // the status of whether you have all unique cards yet (reverse because please for while loop work)

	// a loop that buys and checks cards until you have all the cards
	for hasAll {
		sum++

		addCard := randCard(cardNum)
		fmt.Printf("%s %d\n", "Your random card is:", addCard)

		if len(cardCollection) == 0 {
			// if there are no cards in the collection, just add the first card in
			cardCollection = append(cardCollection, addCard)
		} else if contains(cardCollection, addCard) == false {
			//if this is a new card, add it to the collection
			cardCollection = append(cardCollection, addCard)
			fmt.Println("Add!")
			// check to see if you have all the cards now
			if len(cardCollection) == 35 {
				hasAll = false
			}
		}
	}
	return sum
}

/**
* A function that returns a random number
* from 0 to (given number of cards) - 1
 */
func randCard(cardNum int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(cardNum) // the range for getting random numbers
}

/**
* Checks to see if a given integer is present in the slice. If the slice has it, return true. else, false
* function copied from freshman.tech
* I could convert the for loop to using int i increment but I don't fee like it
 */
func contains(intArr []int, num int) bool {
	for _, v := range intArr {
		if v == num {
			return true
		}
	}
	return false
}

func main() {
	numsCard, err1 := strconv.Atoi(os.Args[1]) // number of cards to collect

	// if the user does not give a number, exit program
	if err1 != nil {
		fmt.Printf("ERR: %v\n", err1)
		os.Exit(3)
	}

	runs := 1 // number of trials
	sum := 0

	// run the trials
	for i := 0; i < runs; i++ {
		sum = sum + collect(numsCard)
	}

	var output float64 = float64(sum / runs) // average amount of cards you need to get
	fmt.Printf("%s %d %s %f\n", "Average Number of Candies bought to get", numsCard, "Cards: ", output)
}
