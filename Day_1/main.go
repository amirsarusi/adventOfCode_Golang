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
	is := getIntSlice(bs)
	for i := 0; i < len(is); i++ {
		for j := 1; j < len(is); j++ {
			if is[i]+is[j] == 2020 {
				fmt.Println("values are :", is[i], is[j], "when multiplied ", is[i]*is[j])
				os.Exit(0)
			}
		}
	}

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
