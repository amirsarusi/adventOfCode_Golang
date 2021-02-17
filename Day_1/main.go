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
	//twoNumSol(getIntSlice(bs), 2020)
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
	for i := 0; i < len(is); i++ {
		for j := 1; j < len(is); j++ {
			if is[i]+is[j] == sumTarget {
				fmt.Println("values are :", is[i], is[j], "when multiplied ", is[i]*is[j])
				return is[i] * is[j]
			}
		}
	}
	return -1
}

func threeNumSol(is []int64, sumTarget int64) (product int64) {

	for i := 0; i < len(is); i++ {
		for j := i + 1; j < len(is); j++ {
			for k := j + 1; k < len(is); k++ {
				if is[i]+is[j]+is[k] == sumTarget && is[i] != 0 && is[j] != 0 && is[k] != 0 {
					fmt.Println("values are :", is[i], is[j], is[k], "when multiplied ", is[i]*is[j]*is[k])
					return is[i] * is[j] * is[k]
				}
			}
		}
	}
	return -1
}
