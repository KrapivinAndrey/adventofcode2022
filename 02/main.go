package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func isWinner(a string, b string) int {
	rockWin := a == "C" && b == "X"
	scissorsWin := a == "A" && b == "Y"
	paperWin := a == "B" && b == "Z"
	if a == "A" && b == "X" || a == "B" && b == "Y" || a == "C" && b == "Z" {
		return 3
	} else if rockWin || paperWin || scissorsWin {
		return 6
	} else {
		return 0
	}

}

func cost(a string) int {
	res := 0
	switch a {
	case "X", "A":
		res = 1
	case "Y", "B":
		res = 2
	case "Z", "C":
		res = 3
	}
	return res
}

func point(a string) int {
	res := 0
	switch a {
	case "X":
		res = 0
	case "Y":
		res = 3
	case "Z":
		res = 6

	}
	return res
}

func my_move(a string, b string) string {
	if b == "Y" {
		return a
	}
	if b == "X" {
		switch a {
		case "A":
			return "C"
		case "B":
			return "A"
		case "C":
			return "B"

		}
	}

	if b == "Z" {
		switch a {
		case "A":
			return "B"
		case "B":
			return "C"
		case "C":
			return "A"

		}
	}

	return "0"

}

func main() {

	var (
		sum1 = 0
		sum2 = 0
	)
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		move := strings.Split(s, " ")
		sum1 += isWinner(move[0], move[1]) + cost(move[1])
		sum2 += point(move[1]) + cost(my_move(move[0], move[1]))
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
}
