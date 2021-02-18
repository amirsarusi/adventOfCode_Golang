package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	twoNumSol(getIntSlice(bs), 2020)
	threeNumSol(getIntSlice(bs), 2020)
}

func getIntSlice(bs []byte) []int64 {
	is := make([]int64, len(bs))
	byteString := string(bs)
	splittedString := strings.Split(byteString, "\n")
	for i, num := range splittedString {
		i64, err := strconv.ParseInt(string(num), 10, 64)
		if err != nil {
			fmt.Println("Error : ", err)
		} else {
			is[i] = i64
		}
	}
	return is
}

func twoNumSol(is []int64, sumTarget int64) (product int64) {
	return findPair(0, sumTarget, is)
}

func threeNumSol(is []int64, sumTarget int64) (product int64) {
	for i := 0; i < len(is); i++ {
		difInvoice := sumTarget - is[i]
		found := findPair(int64(i), difInvoice, is)
		if found > 0 {
			fmt.Println(found * is[i])
			return found * is[i]
		}
	}
	return -1
}

func findPair(index int64, target int64, is []int64) int64 {
	m := make(map[int64]bool)
	for i := index; int(i) < len(is); i++ {
		currentVal := is[i]
		differenceVal := target - currentVal
		if m[differenceVal] {
			return currentVal * differenceVal
		}
		m[currentVal] = true
	}
	return -1
}
