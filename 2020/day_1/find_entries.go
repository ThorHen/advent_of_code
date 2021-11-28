package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getInputs(filePath string) []int {
	// Get entries as byte slice
	data, err := os.ReadFile(filePath)
	check(err)

	// Convert to string and split string into slice on newline delimiter
	dataStr := string(data)
	strSlice := strings.Split(dataStr, "\n")

	// Convert string slice to int slice
	intSlice := make([]int, len(strSlice))
	for i, elem := range strSlice {
		intSlice[i], err = strconv.Atoi(elem)
		check(err)
	}

	return intSlice
}

func findElementsThatAddTo(elements []int, result int) int {
	for i, elem0 := range elements {
		for j := i + 1; j < len(elements); j++ {
			for k := j + 1; k < len(elements); k++ {
				if elem0+elements[j]+elements[k] == 2020 {
					return elem0 * elements[j] * elements[k]
				}
			}

		}
	}

	return 0
}

func main() {
	inputs := getInputs("entries.txt")
	fmt.Println(findElementsThatAddTo(inputs, 2020))
}
