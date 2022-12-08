package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func isSeen(forest []string, x int, y int, dx int, dy int) bool {

	if x == 0 || y == 0 || x == len(forest[0])-1 || y == len(forest)-1 {
		return true
	}
	myHeight, _ := strconv.Atoi(string(forest[x][y]))
	for (x > 0 && x < len(forest[0])-1) && (y > 0 && y < len(forest)-1) {
		x += dx
		y += dy
		newHeight, _ := strconv.Atoi(string(forest[x][y]))
		if newHeight >= myHeight {
			return false
		}
	}
	return true
}

func viewingDistance(forest []string, x int, y int, dx int, dy int) int {
	result := 0
	myHeight, _ := strconv.Atoi(string(forest[x][y]))
	for (x > 0 && x < len(forest[0])-1) && (y > 0 && y < len(forest)-1) {
		x += dx
		y += dy
		result++
		newHeight, _ := strconv.Atoi(string(forest[x][y]))
		if newHeight >= myHeight {
			break
		}
	}

	return result
}

func main() {
	forest := readInput()
	sum := 0
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[i]); j++ {
			if isSeen(forest, i, j, 1, 0) || isSeen(forest, i, j, -1, 0) || isSeen(forest, i, j, 0, 1) || isSeen(forest, i, j, 0, -1) {
				sum++
			}
		}
	}
	fmt.Println(sum)

	maximum := 0
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[i]); j++ {
			distance := viewingDistance(forest, i, j, 1, 0) * viewingDistance(forest, i, j, -1, 0) * viewingDistance(forest, i, j, 0, 1) * viewingDistance(forest, i, j, 0, -1)

			if distance >= maximum {
				maximum = distance
			}
		}
	}

	fmt.Println(maximum)
}
