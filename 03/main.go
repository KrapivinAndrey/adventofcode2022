package main

import (
	"fmt"
	"strings"
)

func inArray(element string, arr []string) bool {

	var result bool = false
	for _, x := range arr {
		if x == element {
			result = true
			break
		}
	}

	return result

}

func cost(elem string) int {
	var priority int
	if elem >= "a" {
		priority = int(elem[0]) - int("a"[0]) + 1
	} else {
		priority = int(elem[0]) - int("A"[0]) + 27
	}

	return priority
}

func checkRucksack(in string) []string {
	var result []string
	l := len(in)
	left := in[0 : l/2]
	right := in[l/2 : l]

	for i := 0; i < l/2; i++ {
		a := string(left[i])
		if strings.Contains(right, a) && !inArray(a, result) {

			result = append(result, a)

		}
	}

	return result

}

func checkLabel(rucksacks []string) string {

	var result string
	for i := 0; i < len(rucksacks[0]); i++ {
		a := string(rucksacks[0][i])

		if strings.Contains(rucksacks[1], a) && strings.Contains(rucksacks[2], a) {
			result = a
			break
		}

	}

	return result

}
func main() {

	tasks := readInput()

	var sum int
	for _, rucksack := range tasks {

		diff := checkRucksack(rucksack)
		for _, elem := range diff {
			sum += cost(elem)
		}
	}

	fmt.Println(sum)

	sum = 0
	for i, _ := range tasks {

		if i > 0 && i%3 == 2 {
			elem := checkLabel(tasks[i-2 : i+1])
			sum += cost(elem)
		}
	}

	fmt.Println(sum)
}
