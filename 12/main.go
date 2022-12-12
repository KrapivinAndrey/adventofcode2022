package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type position struct {
	x int
	y int
}

func getPosition(maps []string, x int, y int) byte {
	letter := maps[y][x]
	if letter == ("S"[0]) {
		return "a"[0]
	} else if letter == ("E"[0]) {
		return "z"[0]
	} else {
		return letter
	}
}

func getNeighbors(maps []string, ver position) []position {

	var result []position

	width := len(maps[0])
	height := len(maps)
	current := getPosition(maps, ver.x, ver.y)

	if ver.x > 0 {
		left := getPosition(maps, ver.x-1, ver.y)
		if left-1 <= current {
			result = append(result, position{ver.x - 1, ver.y})
		}
	}

	if ver.x < width-1 {
		right := getPosition(maps, ver.x+1, ver.y)
		if right-1 <= current {
			result = append(result, position{ver.x + 1, ver.y})
		}
	}

	if ver.y > 0 {
		up := getPosition(maps, ver.x, ver.y-1)
		if up-1 <= current {
			result = append(result, position{ver.x, ver.y - 1})
		}
	}

	if ver.y < height-1 {
		down := getPosition(maps, ver.x, ver.y+1)
		if down-1 <= current {
			result = append(result, position{ver.x, ver.y + 1})
		}
	}

	return result
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

func bfs(maps []string, start position, finish position) int {
	distance := map[position]int{}
	visited := map[position]bool{}
	q := []position{start}
	u := position{0, 0}
	visited[start] = true
	distance[start] = 0
	for len(q) > 0 {
		u, q = q[0], q[1:]
		neighbors := getNeighbors(maps, u)

		for _, v := range neighbors {
			if !visited[v] {

				distance[v] = distance[u] + 1
				visited[v] = true
				q = append(q, v)

			}
		}
	}

	return distance[finish]
}

func main() {

	mountains := readInput()

	start := position{}
	finish := position{}

	// Ищем старт и финиш

	for i, s := range mountains {

		if strings.Contains(s, "S") {
			start = position{strings.Index(s, "S"), i}
		}

		if strings.Contains(s, "E") {
			finish = position{strings.Index(s, "E"), i}
		}
	}

	fmt.Println(bfs(mountains, start, finish))

	minimal := 100000
	for j, layer := range mountains {
		for i, point := range layer {

			if string(point) == "a" {
				dist := bfs(mountains, position{i, j}, finish)
				if dist != 0 && dist < minimal {
					minimal = dist
				}
			}

		}
	}

	fmt.Println(minimal)
}
