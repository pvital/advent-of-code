// Copyright © 2021 Paulo Vital <paulo@vital.eng.br>
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Advent of Code 2021 - Day 02
// Drive!

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type submarinePosition struct {
	horizontal int
	depth int
}


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


func submarineDrive(data []string) (*submarinePosition, error) {
	submarine_pos := &submarinePosition{}
	submarine_pos.horizontal = 0
	submarine_pos.depth = 0

	for _, line := range data {
		directions := strings.Split(line, " ")

		unit, err := strconv.Atoi(directions[1])
		if err != nil {
			return nil, err
		}

		switch directions[0] {
		case "forward":
			submarine_pos.horizontal += unit
		case "down":
			submarine_pos.depth += unit
		case "up":
			submarine_pos.depth -= unit
		default:
			fmt.Printf("Submarine has stopped.\n")
		}
	}

	return submarine_pos, nil
}


func main() {
	year := 2021
	day := 2
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

	yellow_submarine, err := submarineDrive(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Final horizontal position: %d\n", yellow_submarine.horizontal)
	fmt.Printf("Final depth position: %d\n", yellow_submarine.depth)
	fmt.Printf(
		"Final position: %d\n", 
		yellow_submarine.horizontal * yellow_submarine.depth,
	)
}