package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
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

func maximum(a, b int) int {
	var result int
	if a > b {
		result = a
	} else {
		result = b
	}

	return result
}

func toList(x interface{}) []interface{} {
	var elem []interface{}

	rt := reflect.TypeOf(x)

	if rt.Kind() == reflect.Slice {
		elem = x.([]interface{})
	} else {
		elem = append(elem, x)
	}

	return elem

}

func equals(a, b interface{}) bool {

	rtA := reflect.TypeOf(a)
	rtB := reflect.TypeOf(b)

	if rtA.Kind() == reflect.Float64 && rtB.Kind() == reflect.Float64 {
		return a == b
	} else if rtA.Kind() == reflect.Slice && rtB.Kind() == reflect.Slice {
		if len(a.([]interface{})) != len(b.([]interface{})) {
			return false
		}

		for i, e := range a.([]interface{}) {
			if !equals(e, b.([]interface{})[i]) {
				return false
			}

		}
		return true
	}

	return equals(toList(a), toList(b))

}

func less(a, b interface{}) bool {
	rtA := reflect.TypeOf(a)
	rtB := reflect.TypeOf(b)

	if rtA.Kind() == reflect.Float64 && rtB.Kind() == reflect.Float64 {
		return a.(float64) < b.(float64)
	} else if rtA.Kind() == reflect.Slice && rtB.Kind() == reflect.Slice {
		elemA := a.([]interface{})
		elemB := b.([]interface{})
		i := 0
		for i < maximum(len(elemA), len(elemB)) {
			if i >= len(elemA) {
				return true
			}

			if i >= len(elemB) {
				return false
			}

			if equals(elemA[i], elemB[i]) {
				i++
				continue
			} else {
				return less(elemA[i], elemB[i])
			}
		}
	}

	return less(toList(a), toList(b))

}
func main() {

	var (
		signal1, signal2, signal, first, second interface{}
		mySignals                               []interface{}
		sum                                     int
	)
	i := 0
	signals := readInput()
	for i < len(signals) {
		json.Unmarshal([]byte(signals[i]), &signal1)
		json.Unmarshal([]byte(signals[i+1]), &signal2)

		if less(signal1, signal2) {
			sum += (i / 3) + 1
		}
		i += 3
	}

	fmt.Println(sum)

	// похер, сделаю пызырек поздно уже

	for _, line := range readInput() {
		if line == "" {
			continue
		}
		json.Unmarshal([]byte(line), &signal)
		mySignals = append(mySignals, signal)
	}

	json.Unmarshal([]byte("[[2]]"), &first)
	mySignals = append(mySignals, first)
	json.Unmarshal([]byte("[[6]]"), &second)
	mySignals = append(mySignals, second)

	flag := true
	for flag {
		flag = false
		for i := 0; i < len(mySignals)-1; i++ {

			if !less(mySignals[i], mySignals[i+1]) {
				mySignals[i], mySignals[i+1] = mySignals[i+1], mySignals[i]
				flag = true
			}
		}
	}

	for i, elem := range mySignals {
		if equals(elem, first) || equals(elem, second) {
			fmt.Println(i + 1)
		}
	}

}
