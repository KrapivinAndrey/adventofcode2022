package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type File struct {
	name   string
	weight int
}

type Directory struct {
	name           string
	subdirectories map[string]*Directory
	files          []*File
	parent         *Directory
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

func (d *Directory) totalWeight() int {
	result := 0
	for _, file := range d.files {
		result += file.weight
	}

	for _, dir := range d.subdirectories {
		result += dir.totalWeight()
	}

	return result
}

func main() {
	var allDirectories []*Directory
	commands := readInput()
	root := new(Directory)
	root.name = "root"
	root.subdirectories = make(map[string]*Directory)
	allDirectories = append(allDirectories, root)

	currentDirectory := &root

	i := 0
	for i < len(commands) {
		if commands[i] == "$ cd /" {
			currentDirectory = &root
			i++
		} else if commands[i] == "$ cd .." {
			currentDirectory = &((*currentDirectory).parent)
			i++
		} else if commands[i] == "$ ls" {
			i++
			for (i < len(commands)) && (string(commands[i][0]) != "$") {
				components := strings.Split(commands[i], " ")
				if components[0] == "dir" {
					dir := new(Directory)
					dir.name = components[1]
					dir.parent = *currentDirectory
					dir.subdirectories = make(map[string]*Directory)
					allDirectories = append(allDirectories, dir)
					(*currentDirectory).subdirectories[components[1]] = dir

				} else {
					weight, _ := strconv.Atoi(components[0])
					file := File{components[1], weight}
					(*currentDirectory).files = append((*currentDirectory).files, &file)
				}
				i++
			}
		} else {
			components := strings.Split(commands[i], " ")
			newDir := (*currentDirectory).subdirectories[components[2]]
			currentDirectory = &newDir
			i++
		}

	}

	answer := 0
	for _, dir := range allDirectories {
		dirWeight := dir.totalWeight()
		if dirWeight < 100000 {
			answer += dirWeight
		}
	}

	fmt.Print("Первая часть. Ответ ")
	fmt.Println(answer)

	totalSpace := 70000000
	needSpace := 30000000
	freeSpace := totalSpace - root.totalWeight()

	minimal := root.totalWeight()
	for _, dir := range allDirectories {
		weight := dir.totalWeight()
		if freeSpace+weight >= needSpace && weight < minimal {
			minimal = weight
		}
	}

	fmt.Print("Вторая часть. Ответ ")
	fmt.Println(minimal)

}
