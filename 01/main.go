package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	var (
		sum = 0
	)
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	elves := []int{1}
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			elves = append(elves, sum)
			sum = 0
		} else {
			intValue := 0
			fmt.Sscan(s, &intValue)
			sum = sum + intValue
		}
	}
	elves = append(elves, sum)
	sort.Ints(elves)
	//first
	fmt.Println(elves[len(elves)-1])
	//second
	fmt.Println(elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3])

}
