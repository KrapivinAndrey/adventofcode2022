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

type Dot struct {
	x, y int
}

func maximum(a, b int) int {
	result := 0
	if a > b {
		result = a
	} else {
		result = b
	}

	return result
}

func minimum(a, b int) int {
	result := 0
	if a < b {
		result = a
	} else {
		result = b
	}

	return result
}

func main() {
	var (
		bottom = -1
		left   = 10000
		right  = 0
	)
	rocks := readInput()
	cave := map[Dot]int{} // 0 - пусто, 1 - скала, 2 - песок

	for _, rock := range rocks {
		layer := strings.Split(rock, " -> ")
		prevX := -1
		prevY := -1
		for _, s := range layer {

			components := strings.Split(s, ",")
			x, _ := strconv.Atoi(components[0])
			y, _ := strconv.Atoi(components[1])

			bottom = maximum(bottom, y)
			left = minimum(left, x)
			right = maximum(right, x)
			if prevX == -1 && prevY == -1 {

				prevX = x
				prevY = y
				continue
			}

			for i := minimum(x, prevX); i <= maximum(x, prevX); i++ {
				for j := minimum(y, prevY); j <= maximum(y, prevY); j++ {
					cave[Dot{i, j}] = 1

				}
			}

			prevX = x
			prevY = y

		}

	}

	fmt.Println(cave)

}
