/**
* An implementation of the "Baby Names" program in Go where, when the user
* gives a name or set of names, checks a "database" (a set of files with
* information on names for many years) and returns information about
* that name if it finds it as well as for names that contain that given name
* as a prefix.
* author @ayang
* Created: Dec 5, 2021
 */

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

type BabyName struct {
	Name   string
	Amount int
	Years  int
}

// ByName implements sort.Interface based on the Name field
type ByName []BabyName

func (name ByName) Len() int           { return len(name) }
func (name ByName) Less(i, j int) bool { return name[i].Name < name[j].Name }
func (name ByName) Swap(i, j int)      { name[i], name[j] = name[j], name[i] }

/*
	Modified version of the readHurricane code provided by Prof. Towell,
	changed so that it returns a string slice that contains each row of the
	excel sheet as an element, ideally

	@param url (the url to get the information from), maxLines (the amount of lines you want read)

	@return nameLst (a slice containing all information from the specified amount of lines), nil iff
	there was an error in reading the url
*/
func readName(url string, maxLines int) []string {
	resp, err := http.Get(url)
	// if there is a problem with the url, error message
	if err != nil {
		fmt.Println("Http problem", err)
		return nil
	}

	// read all the text in the url and then close
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	// if there is a problem with reading the text in the url, error message
	if err != nil {
		fmt.Println("Http problem2", err)
		return nil
	}

	// convert the text read into a string
	str := string(body)
	r := csv.NewReader(strings.NewReader(str))
	r.FieldsPerRecord = -1

	var nameLst []string

	// store the value for each line
	for i := 0; i < maxLines || maxLines <= 0; i++ {
		// for every line read it
		record, err := r.Read()

		// if there is no more input to read, stop
		if err == io.EOF {
			break
		}
		// if there is an error, print out the error and stop
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			break
		}
		name := strings.Join(record, ",")
		nameLst = append(nameLst, name)
		// return string???
	}
	return nameLst
}

/*
	Store names from a string (from a slice) into two separate slices based on whether the name is
	masculine or feminine usage

	@param: name (section from slice to sort), boyLst and girlLst (slices for masc/femme usage respectively)
	@return: updated boyLst and girlLst (bc Go is something something referenced based so I don't have fun :/)
*/
func nameStorer(name string, boyLst []BabyName, girlLst []BabyName) ([]BabyName, []BabyName) {
	nameLst := strings.Split(name, ",") // make the row into a slice of strings split at commas

	// get the amounts for the names. BEFORE anything else because I will use this A Lot
	boyAmount, err := strconv.Atoi(strings.TrimSpace(nameLst[2]))
	if err != nil {
		fmt.Fprintf(os.Stdout, "CONV PROB [[%v]]  %v\n", nameLst[2], err)
	}
	girlAmount, err := strconv.Atoi(strings.TrimSpace(nameLst[4]))
	if err != nil {
		fmt.Fprintf(os.Stdout, "CONV PROB [[%v]]  %v\n", nameLst[4], err)
	}

	// if the list is empty, just add the name. since boy & girl list are  the same length uh some questionable things are being done
	if len(boyLst) == 0 && len(girlLst) == 0 {
		// add the item to the list
		boyLst = append(boyLst, BabyName{nameLst[1], boyAmount, 1})
		girlLst = append(girlLst, BabyName{nameLst[3], girlAmount, 1})
		return boyLst, girlLst
	}

	// check to see if an instance of that name exists already
	boyIndx := -1
	girlIndx := -1
	for i := 0; i < len(boyLst); i++ {
		if strings.EqualFold(nameLst[1], boyLst[i].Name) {
			boyIndx = i // if there is an instance of the name in the list already, just update it.
			break
		}
	}
	for i := 0; i < len(girlLst); i++ {
		if strings.EqualFold(nameLst[3], girlLst[i].Name) {
			girlIndx = i // if there is an instance of the name in the list already, just update it
			break
		}
	}

	// all combinations of cases of already exists/doesn't already exist
	if boyIndx != -1 && girlIndx != -1 {
		// in the case where both already exist, just update them both
		boyLst[boyIndx].Amount = boyLst[boyIndx].Amount + boyAmount
		boyLst[boyIndx].Years = boyLst[boyIndx].Years + 1

		girlLst[girlIndx].Amount = girlLst[girlIndx].Amount + girlAmount
		girlLst[girlIndx].Years = girlLst[girlIndx].Years + 1

		sort.Sort(ByName(boyLst))
		sort.Sort(ByName(girlLst))

		return boyLst, girlLst
	} else if boyIndx != -1 && girlIndx == -1 {
		// in the case where boy already exists and girl doesn't update boy and add/sort girl
		boyLst[boyIndx].Amount = boyLst[boyIndx].Amount + boyAmount
		boyLst[boyIndx].Years = boyLst[boyIndx].Years + 1

		girlLst = append(girlLst, BabyName{nameLst[3], girlAmount, 1})

		sort.Sort(ByName(boyLst))
		sort.Sort(ByName(girlLst))

		return boyLst, girlLst
	} else if boyIndx == -1 && girlIndx != -1 {
		// in teh case where girl already exists and boy doesn't update girl and add/sort boy
		boyLst = append(boyLst, BabyName{nameLst[1], boyAmount, 1})

		girlLst[girlIndx].Amount = girlLst[girlIndx].Amount + girlAmount
		girlLst[girlIndx].Years = girlLst[girlIndx].Years + 1

		sort.Sort(ByName(boyLst))
		sort.Sort(ByName(girlLst))

		return boyLst, girlLst
	} else {
		// in the case where both girl and boy doesn't exist, add and sort for both
		boyLst = append(boyLst, BabyName{nameLst[1], boyAmount, 1})
		girlLst = append(girlLst, BabyName{nameLst[3], girlAmount, 1})

		sort.Sort(ByName(boyLst))
		sort.Sort(ByName(girlLst))

		return boyLst, girlLst
	}
}

/*
	Given a name and slices to search from, check to see if that given name (and any name that contains the given name
	as a prefix) exist in the list. If so, print out the name along with some statistics.

	@param: name (name you are looking for instances of), boyLst and girLst (lists you are searching for instances of "name")
*/
func nameSearch(name string, boyLst []BabyName, girlLst []BabyName) {
	var boyIndxs []int  // slice of indexes for the name to print information about for masc names
	var girlIndxs []int // slice of indexes for the name to print information about for femme names

	// check to see if the name itself is in the slice
	checkName := func(name string, lst []BabyName, lstModify []int) []int {
		for i := 0; i < len(lst); i++ {
			if strings.EqualFold(name, lst[i].Name) {
				lstModify = append(lstModify, i)
			}
		}
		return lstModify
	}

	girlIndxs = checkName(name, girlLst, girlIndxs)
	boyIndxs = checkName(name, boyLst, boyIndxs)

	// check for prefixes
	checkPrefix := func(name string, lst []BabyName, lstModify []int) []int {
		for i := 0; i < len(lst); i++ {
			// dodging the issue of "name you are comparing to is too small to subslice"
			if len(lst[i].Name) >= len(name) {
				// take a substring from the beginning of the size of the name you are comparing to
				prefName := lst[i].Name[0:len(name)]
				if strings.EqualFold(name, prefName) {
					// enter: a convoluted way to avoid duplicates where I don't just write a contains method
					hasName := false // assume it isn't in the ret list
					for j := 0; j < len(lstModify); j++ {
						if i == lstModify[j] {
							hasName = true // if it is in the ret list, change the traacker
						}
					}
					if !hasName {
						// if it is an prefix for the name from the slice, add it to the list
						// fmt.Println("in prefix checker")
						// fmt.Println(i)
						lstModify = append(lstModify, i)
					}
				}
			}
		}
		return lstModify
	}

	girlIndxs = checkPrefix(name, girlLst, girlIndxs)
	boyIndxs = checkPrefix(name, boyLst, boyIndxs)

	// get the total to calculate percentages with
	boyTotal := 0
	girlTotal := 0

	for i := 0; i < len(boyLst); i++ {
		boyTotal = boyTotal + boyLst[i].Amount
	}
	for i := 0; i < len(girlLst); i++ {
		girlTotal = girlTotal + girlLst[i].Amount
	}

	// return results
	fmt.Println(name) // print out the search results for the name

	// if the slice for the respective gender ISN'T empty, print things out. otherwise, print nothing out
	// for masc names:
	if len(boyIndxs) != 0 {
		// fmt.Println("*Snake voice* im in")
		for i := 0; i < len(boyIndxs); i++ {
			nameToPrint := boyLst[boyIndxs[i]] // so i don't have to write "boyLst[boyIndxs[i]]" a lot because it's cumberson and ugly :(

			// find percentage and rank here
			percent := (float64(nameToPrint.Amount) / float64(boyTotal)) * 100
			rank := boyIndxs[i] + 1

			// printing out the information
			fmt.Printf("%s: %s \t %s: %s \t %s: %d\n", "Name", nameToPrint.Name, "Gender", "Boy", "# of Years", nameToPrint.Years)
			fmt.Printf("%s: %d\n", "Alphabetical Position", rank)
			fmt.Printf("%s: %d\n", "Occurences", nameToPrint.Amount)
			fmt.Printf("%s %s %s: %f\n\n", "Percent of All", "Boys", "Names", percent)
		}
	}
	// for femme names:
	if len(girlIndxs) != 0 {
		for i := 0; i < len(girlIndxs); i++ {
			nameToPrint := girlLst[girlIndxs[i]] // so i don't have to write "boyLst[boyIndxs[i]]" a lot because it's cumberson and ugly :(

			// find percentage and rank here
			percent := (float64(nameToPrint.Amount) / float64(girlTotal)) * 100
			rank := girlIndxs[i] + 1

			// printing out the information
			fmt.Printf("%s: %s \t %s: %s \t %s: %d\n", "Name", nameToPrint.Name, "Gender", "Girl", "# of Years", nameToPrint.Years)
			fmt.Printf("%s: %d\n", "Alphabetical Position", rank)
			fmt.Printf("%s: %d\n", "Occurences", nameToPrint.Amount)
			fmt.Printf("%s %s %s: %f\n\n", "Percent of All", "Girl", "Names", percent)
		}
	}
}

func main() {
	var GirlNames []BabyName // list holding information on girl names
	var BoyNames []BabyName  // list holding information on boy names
	var NameQuery []string   // list of names to return results for

	// NameQuery := [4]string{"Aaron", "Devon", "marlen", "syd"}

	// if the user doesn't give any names to test for, yell at them to give name and then end early.
	if len(os.Args[1:]) < 1 {
		fmt.Println("Please enter name(s) to search for!")
		os.Exit(0)
	}
	// get all user inputs
	for i := 1; i < len(os.Args[1:])+1; i++ {
		NameQuery = append(NameQuery, os.Args[i])
	}

	// lsit of names. man, wish there was a way to collapse this!!
	filesLst := [10]string{"https://cs.brynmawr.edu/cs245/Data/names2000.csv",
		"https://cs.brynmawr.edu/cs245/Data/names2001.csv",
		"https://cs.brynmawr.edu/cs245/Data/names2002.csv",
		"https://cs.brynmawr.edu/cs245/Data/names2003.csv",
		"https://cs.brynmawr.edu/cs245/Data/names2004.csv",
		"https://cs.brynmawr.edu/cs245/Data/names2005.csv",
		"https://cs.brynmawr.edu/cs245/Data/names2006.csv",
		"https://cs.brynmawr.edu/cs245/Data/names2007.csv",
		"https://cs.brynmawr.edu/cs245/Data/names2008.csv",
		"https://cs.brynmawr.edu/cs245/Data/names2009.csv"}

	//list that holds the string for each name
	nameLst := make([][]string, 10)

	// read all the files
	for i := 0; i < len(filesLst); i++ {
		var url string
		url = filesLst[i]
		nameLst[i] = readName(url, 1000)
	}

	// put all the information from the names to their respective lists
	for i := 0; i < len(nameLst); i++ {
		for j := 0; j < len(nameLst[i]); j++ {
			// fmt.Println(nameLst[i][j])
			BoyNames, GirlNames = nameStorer(nameLst[i][j], BoyNames, GirlNames)
			// fmt.Println(BoyNames)
		}
	}

	// run the search for if that name (or if a name containing that name as a prefix) exists in the slices
	for i := 0; i < len(NameQuery); i++ {
		nameSearch(NameQuery[i], BoyNames, GirlNames)
	}
}
