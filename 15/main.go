package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func abs(a int) int {
	result := a
	if a < 0 {
		result = -a
	}

	return result
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

func manhattan(a, b Dot) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
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

func readSensors() []Pair {
	var (
		cave []Pair
	)
	input := readInput()

	re, _ := regexp.Compile("(-?\\d+)")
	for _, s := range input {
		res := re.FindAllString(s, -1)
		var resInt []int
		for _, pair := range res {
			num, _ := strconv.Atoi(pair)
			resInt = append(resInt, num)
		}
		sensor := Dot{resInt[0], resInt[1]}
		beacon := Dot{resInt[2], resInt[3]}
		radius := manhattan(sensor, beacon)

		newPair := Pair{sensor, beacon, radius}
		cave = append(cave, newPair)
	}

	return cave
}

type Dot struct {
	x, y int
}

type Pair struct {
	sensor, beacon Dot
	radius         int
}

type Segment struct {
	x1, x2 int
}

func (s *Segment) len() int {
	return s.x2 - s.x1
}

func (a *Pair) into(x Dot) bool {

	return manhattan(a.sensor, x) <= a.radius && x != a.beacon

}

func slice(caves []Pair, y int) []Segment {
	var (
		slices []Segment
	)
	for _, cave := range caves {

		if abs(y-cave.sensor.y) > cave.radius {
			// Точка вне круга точно
			continue
		}

		x1 := cave.sensor.x - (cave.radius - abs(y-cave.sensor.y))
		x2 := cave.sensor.x + (cave.radius - abs(y-cave.sensor.y))

		slices = append(slices, Segment{x1, x2})

	}

	return slices

}

func combine(segments []Segment) []Segment {
	var result []Segment

	sort.Slice(segments, func(i, j int) bool {
		return segments[i].x1 < segments[j].x1
	})
	if len(segments) == 0 {
		return result
	}
	segment := segments[0]
	for _, other := range segments[1:] {
		if segment.x2+1 >= other.x1 {
			segment = Segment{segment.x1, maximum(other.x2, segment.x2)}
		} else {
			result = append(result, segment)
			segment = other
		}
	}

	result = append(result, segment)
	return result
}

func main() {
	var (
		caves []Pair
	)
	caves = readSensors()
	//notFree := 0

	slices := slice(caves, 2000000)

	sum := 0
	for _, s := range combine(slices) {
		sum += s.len()
	}
	fmt.Println(sum)

	for y := 0; y <= 4000000; y++ {
		all := combine(slice(caves, y))
		if len(all) == 2 {
			fmt.Println((all[0].x2+1)*4000000 + y)
		}

	}

}
