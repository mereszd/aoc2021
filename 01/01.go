package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Solution for the first part
/*
func main() {

	dat, err := os.Open("input.txt")
	check(err)
	sc := bufio.NewScanner(dat)
	depths := make([]int, 0)

	for sc.Scan() {
		intVar, err := strconv.Atoi(sc.Text())
		check(err)
		depths = append(depths, intVar)
	}

	increases := 0

	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			increases = increases + 1
		}
	}

	fmt.Println(increases)

}
*/

func main() {

	dat, err := os.Open("input.txt")
	check(err)
	sc := bufio.NewScanner(dat)
	depths := make([]int, 0)

	for sc.Scan() {
		intVar, err := strconv.Atoi(sc.Text())
		check(err)
		depths = append(depths, intVar)
	}

	increases := 0

	for i := 3; i < len(depths); i++ {
		if depths[i] > depths[i-3] {
			increases += 1
		}
	}

	fmt.Println(increases)

}
