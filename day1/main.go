package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
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

	firstList := make([]int, 0)
	secondList := make([]int, 0)

	freqList := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		println(line)
		pattern := "\\s+"
		regex := regexp.MustCompile(pattern)
		result := regex.Split(line, -1)

		firstInt, err := strconv.Atoi(result[0])
		if err != nil {
			panic(err)
		}
		secondInt, err := strconv.Atoi(result[1])
		if err != nil {
			panic(err)
		}

		freqList[secondInt] += 1

		firstList = append(firstList, firstInt)
		secondList = append(secondList, secondInt)
	}

	slices.Sort(firstList)
	slices.Sort(secondList)

	var distance int
	for i, num1 := range firstList {
		if num1 < secondList[i] {
			distance += secondList[i] - num1
		} else {
			distance += num1 - secondList[i]
		}
	}

	fmt.Println(distance)

	fmt.Println(similarity(firstList, freqList))
}

func similarity(left []int, freq map[int]int) int {
	similarity := 0

	for _, num := range left {
		similarity += (num * freq[num])
	}

	return similarity
}
