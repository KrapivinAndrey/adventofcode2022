package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func checkDiff(a string) bool {
	smap := map[byte]int{}

	for i := 0; i < len(a); i++ {
		(smap[a[i]])++

	}

	for _, v := range smap {

		if v != 1 {
			return false
		}
	}

	return true
}

func main() {

	signal := readInput()[0]
	for i := 0; i < len(signal); i++ {
		part := string([]rune(signal)[i : i+4])
		if checkDiff(part) {
			fmt.Println(i + 4)
			break
		}
	}
}
