package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() []string {
	var result []string
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		result = append(result, s)
	}

	return result
}

func strToIntArr(s string, sep string) []int {
	strs := strings.Split(s, sep)
	res := make([]int, len(strs))
	for i := range res {
		res[i], _ = strconv.Atoi(strs[i])
	}
	return res
}

func fullInclude(first string, second string) bool {

	a := strToIntArr(first, "-")
	b := strToIntArr(second, "-")

	one := (a[0] <= b[0] && b[0] <= a[1]) && (a[0] <= b[1] && b[1] <= a[1])
	two := (b[0] <= a[0] && a[0] <= b[1]) && (b[0] <= a[1] && a[1] <= b[1])
	result := one || two
	return result

}

func overlap(first string, second string) bool {

	a := strToIntArr(first, "-")
	b := strToIntArr(second, "-")

	one := b[0] <= a[0] && a[0] <= b[1] && b[1] <= a[1]
	two := a[0] <= b[0] && b[0] <= a[1] && a[1] <= b[1]
	result := one || two || fullInclude(first, second)
	return result

}

func main() {

	input := readInput()

	sum1 := 0
	sum2 := 0

	for _, elves := range input {
		e := strings.Split(elves, ",")
		if fullInclude(e[0], e[1]) {
			sum1++
		}

		if overlap(e[0], e[1]) {
			sum2++
		}
	}

	fmt.Println(sum1)
	fmt.Println(sum2)

}
