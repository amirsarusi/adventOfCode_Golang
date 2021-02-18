package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type validatePassowrd interface {
	validatePassword(int64, int64, string, string) bool
}

type sledRental struct{}
type toboggan struct{}

func main() {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	ss := strings.Split(string(bs), "\n")
	t := toboggan{}
	fmt.Println("number of valid: ", validateDB(ss, t))
}

func validateDB(ss []string, vp validatePassowrd) int64 {
	var numOfValid int64
	numOfValid = 0
	for _, element := range ss {
		splitOne := strings.Split(element, " ")
		boundsUnparsed := splitOne[0]
		parsedBounds := strings.Split(boundsUnparsed, "-")
		//fmt.Println("parsedBounds: ", parsedBounds)
		min, err := strconv.ParseInt(parsedBounds[0], 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
		}
		max, err := strconv.ParseInt(parsedBounds[1], 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
		}
		char := string(splitOne[1][0])
		password := splitOne[2]
		//fmt.Println("min:", min, "max:", max, "char:", char, "password:", password)

		if vp.validatePassword(min, max, char, password) == true {
			//fmt.Println("got true for", element)
			numOfValid++
		}
	}
	return numOfValid
}

func (sr sledRental) validatePassword(min int64, max int64, char string, password string) bool {
	var numOfAppearences int64
	numOfAppearences = 0
	for i := 0; i < len(password); i++ {
		if string(password[i]) == char {
			numOfAppearences++
		}
	}
	if numOfAppearences < min || numOfAppearences > max {
		return false
	}
	return true
}

func (t toboggan) validatePassword(min int64, max int64, char string, password string) bool {
	flagMin := false
	flagMax := false
	if int(min) <= len(password) {
		flagMin = string(password[min-1]) == char
	}
	if int(max) <= len(password) {
		flagMax = string(password[max-1]) == char
	}
	return flagMin != flagMax
}
