/**
* A program demonstrating JSON use in Go
* Author: @ayang
* Created: October 4, 2021
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type DOB struct {
	Year  int
	Month string
	Day   int
}

type Student struct {
	FirstName          string
	LastName           string
	ID                 string
	Age                string
	DateOfBirth        DOB
	FavoriteRealNumber int
}

type jm map[string]interface{}
type js []jm

func main() {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Sucessfully Opened users.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonFile.Close()

	var student []Student
	json.Unmarshal(byteValue, &student)
	fmt.Printf("Students: %+v\n\n", student)

	// fmt.Printf("Name: %s, Age: %s", student.FirstName, student.Age)
	/*var res js
	json.Unmarshal(byteValue, &res)
	fmt.Println(res)*/

	data, _ := json.Marshal(student)
	fmt.Println(string(data))
}
