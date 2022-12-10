package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sign(x int) int {
	if x == 0 {
		return 0
	} else if x < 0 {
		return -1
	}
	return 1
}

type Dot struct {
	x int
	y int
}

func (a *Dot) isNear(b Dot) bool {
	return Abs(a.x-b.x) <= 1 && Abs(a.y-b.y) <= 1
}

func (a *Dot) isOnLine(b Dot) bool {
	return a.x == b.x || a.y == b.y
}

func (a *Dot) moveDot(dx int, dy int) Dot {
	return Dot{a.x + dx, a.y + dy}
}

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

func main() {

	moves := readInput()
	head := Dot{0, 0}
	tail := Dot{0, 0}
	prev := map[Dot]bool{}

	for _, move := range moves {
		components := strings.Split(move, " ")
		dx := 0
		dy := 0
		switch components[0] {
		case "U":
			dy = 1
		case "D":
			dy = -1
		case "L":
			dx = -1
		case "R":
			dx = 1
		}
		steps, _ := strconv.Atoi(components[1])
		for i := 0; i < steps; i++ {
			head.x += dx
			head.y += dy

			if !head.isNear(tail) {
				tail.x = head.x - dx
				tail.y = head.y - dy
			}
			prev[tail] = true

		}
	}
	fmt.Println(len(prev))

	prev = map[Dot]bool{}

	rope := [10]Dot{}
	for _, move := range moves {
		components := strings.Split(move, " ")
		dx := 0
		dy := 0
		switch components[0] {
		case "U":
			dy = 1
		case "D":
			dy = -1
		case "L":
			dx = -1
		case "R":
			dx = 1
		}
		steps, _ := strconv.Atoi(components[1])
		for step := 0; step < steps; step++ {
			rope[9] = rope[9].moveDot(dx, dy)

			for i := 8; i >= 0; i-- {
				if rope[i].isNear(rope[i+1]) {
					break
				}

				if rope[i].isOnLine(rope[i+1]) {
					rope[i].x += Sign(rope[i+1].x - rope[i].x)
					rope[i].y += Sign(rope[i+1].y - rope[i].y)
				} else {
					ur := rope[i].moveDot(1, 1)
					ul := rope[i].moveDot(-1, 1)
					dr := rope[i].moveDot(1, -1)
					dl := rope[i].moveDot(-1, -1)

					switch {
					case ur.isNear(rope[i+1]):
						rope[i] = ur
					case ul.isNear(rope[i+1]):
						rope[i] = ul
					case dr.isNear(rope[i+1]):
						rope[i] = dr
					case dl.isNear(rope[i+1]):
						rope[i] = dl
					}

				}
			}

			prev[rope[0]] = true
		}
	}
	fmt.Println(len(prev))
}
