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

func strength(cycle int, x int) int {
	if (cycle-20)%40 == 0 {
		fmt.Println(x, x*cycle)
		return x * cycle
	} else {
		return 0
	}
}

func checkCycle(cycle int, x int, row *string) {
	if cycle%40 == 0 {
		fmt.Println(*row)
		*row = ""
	} else {
		pixel := cycle%40 - 1
		if x-1 <= pixel && pixel <= x+1 {
			*row += "#"
		} else {
			*row += "."
		}
	}
}

func main() {
	commands := readInput()
	cycle := 1
	x := 1
	total := 0
	for _, instruction := range commands {
		if instruction == "noop" {
			cycle++
			total += strength(cycle, x)
		} else {
			components := strings.Split(instruction, " ")
			value, _ := strconv.Atoi(components[1])
			cycle++
			total += strength(cycle, x)
			x += value
			cycle++
			total += strength(cycle, x)
		}
	}

	fmt.Println(total)

	fmt.Println("Part two")

	row := ""
	x = 1
	cycle = 1

	for _, instruction := range commands {
		checkCycle(cycle, x, &row)
		if instruction == "noop" {
			cycle++
		} else {
			components := strings.Split(instruction, " ")

			cycle++
			checkCycle(cycle, x, &row)

			value, _ := strconv.Atoi(components[1])
			cycle++
			x += value

		}
	}
	checkCycle(cycle, x, &row)

}
