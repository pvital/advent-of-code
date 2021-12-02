// Copyright © 2021 Paulo Vital <paulo@vital.eng.br>
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Advent of Code 2021 - Day 01
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFileLines(filename string) ([]string, error) {
    file, err := os.Open(filename)  
    if err != nil {
        return nil, err  
    }
	defer file.Close()
  
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

	var data []string  
    for scanner.Scan() {
        data = append(data, scanner.Text())
    }

	return data, nil
}


func calculateIncreases(data []string) (int) {
	count := 0

	last, err := strconv.Atoi(data[0])
	if err != nil {
		fmt.Printf("ERROR converting '%s' to integer.\n", data[0])
	}

	for _, i := range data[1:] {
		curr, err := strconv.Atoi(i)
		if err != nil {
        	fmt.Printf("ERROR converting '%s' to integer.\n", i)
    	}

		if curr > last {
			count++
		}
		last = curr
	}
	return count
}

func main() {
	year := 2021
	day := 1
	fmt.Printf("Advent of Code %d - Day %d\n", year, day)

	inputfile := ""
	if len(os.Args) > 1 {
		inputfile = os.Args[1] 
	} else { 
		inputfile = "input.txt"
	}
	fmt.Printf("Reading file %s\n", inputfile)

	data, err := ReadFileLines(inputfile)
	if err != nil {
		panic(err)
	}

	increase_count := calculateIncreases(data)

	fmt.Printf("Number of increase measurements: %d\n", increase_count)
}