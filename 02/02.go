package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := os.Open("input.txt")
	check(err)
	sc := bufio.NewScanner(dat)
	instructions := make([]string, 0)

	for sc.Scan() {
		instructions = append(instructions, sc.Text())
	}

	final_forward := 0
	final_depth := 0

	for i, s := range instructions {
		_ = i
		split := strings.Split(s, " ")
		ins_value, err := strconv.Atoi(split[1])
		check(err)
		if split[0] == "forward" {
			final_forward = final_forward + ins_value
		} else if split[0] == "up" {
			final_depth = final_depth - ins_value
		} else if split[0] == "down" {
			final_depth = final_depth + ins_value
		}
	}

	solution := final_forward * final_depth

	fmt.Println(solution)
}
