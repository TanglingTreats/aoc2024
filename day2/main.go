package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	_ "slices"
	"strconv"
	_ "strings"
)

func main() {

	fi, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(fi)

	//fmt.Println(problem1(scanner))
	fmt.Println(problem2(scanner))

}

func problem1(scanner *bufio.Scanner) int {
	safety := 0
	for scanner.Scan() {
		line := scanner.Text()

		pattern := "\\s+"
		regex := regexp.MustCompile(pattern)
		result := regex.Split(line, -1)

		// If negative, direction is subtraction. Else addition
		dir := 0

		// Tracks trueness within a report
		isTrue := 0

	Audit:
		for i, res := range result {
			if i+1 == len(result) {
				break Audit
			}

			currNum, err := strconv.Atoi(res)
			if err != nil {
				panic(err)
			}
			nextNum, err := strconv.Atoi(result[i+1])
			if err != nil {
				panic(err)
			}

			op := 0
			if currNum > nextNum {
				if dir == 1 {
					break Audit
				}
				dir = -1

				op = currNum - nextNum

			} else if currNum < nextNum {
				if dir == -1 {
					break Audit
				}
				dir = 1

				op = nextNum - currNum

			} else {
				// When numbers are equal
				break Audit
			}
			if op >= 1 && op <= 3 {
				isTrue += 1
			}
		}

		//fmt.Println("----------")
		//fmt.Printf("%v\n", result)
		//fmt.Println("isTrue: ", isTrue)
		//fmt.Println("Result: ", len(result))

		// Increment safety report when isTrue is the same as number of reports - 1
		if isTrue == len(result)-1 {
			//fmt.Println("Safe")
			safety += 1

		} else {
			//fmt.Println("Unsafe")
		}

		// Reset
		dir = 0
		isTrue = 0
	}

	return safety
}

func problem2(scanner *bufio.Scanner) int {
	diffs := make([]int, 0)
	dirs := make([]int, 0)
	safety := 0
	for scanner.Scan() {
		line := scanner.Text()

		pattern := "\\s+"
		regex := regexp.MustCompile(pattern)
		result := regex.Split(line, -1)

		// If negative, direction is subtraction. Else addition
		dir := 0

		// Tracks trueness within a report
		isTrue := 0

	Audit:
		for i, res := range result {
			if i+1 == len(result) {
				break Audit
			}

			currNum, err := strconv.Atoi(res)
			if err != nil {
				panic(err)
			}
			nextNum, err := strconv.Atoi(result[i+1])
			if err != nil {
				panic(err)
			}

			op := 0
			if currNum > nextNum {
				if dir == 1 {
					break Audit
				}
				dir = -1

				op = currNum - nextNum

			} else if currNum < nextNum {
				if dir == -1 {
					break Audit
				}
				dir = 1

				op = nextNum - currNum

			} else {
				// When numbers are equal
				break Audit
			}
			if op >= 1 && op <= 3 {
				isTrue += 1
			}
		}

		//fmt.Println("----------")
		//fmt.Printf("%v\n", result)
		//fmt.Println("isTrue: ", isTrue)
		//fmt.Println("Result: ", len(result))

		// Increment safety report when isTrue is the same as number of reports - 1
		if isTrue == len(result)-1 {
			//fmt.Println("Safe")
			safety += 1

		} else {
			//fmt.Println("Unsafe")
		}

		// Reset
		dir = 0
		isTrue = 0
	}

	return safety
}
