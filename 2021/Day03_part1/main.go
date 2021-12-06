// Copyright © 2021 Paulo Vital <paulo@vital.eng.br>
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Advent of Code 2021 - Day 03
// Binary Diagnostic

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type diagnosticReport struct {
	gamma_rate int64
	epsilon_rate int64
	power_consumption int64
}


type bitCount struct {
	count_0 int
	count_1 int
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


func calculateCommonBits(data []string) ([]bitCount) {
	// Initialize a slice of bitCount with capacity = length of data
	common_bits := make([]bitCount, 0, len(data[0]))
	for i := 0; i < len(data[0]); i++ {
		counting := &bitCount{}
		counting.count_0 = 0
		counting.count_1 = 0
		common_bits = append(common_bits, *counting)
	}

	// Calculate how many zeros and ones each bit has
	for _, line := range data {
		for i, bit := range strings.Split(line, "") {			
			counting := common_bits[i]
			if bit == "0" {
				counting.count_0++
			} else {
				counting.count_1++
			}
			common_bits[i] = counting
		}
	}

	return common_bits
}


func generateDiagnosticReport(data []string) (*diagnosticReport, error) {
	dr := &diagnosticReport{}
	dr.gamma_rate = 0
	dr.epsilon_rate = 0
	dr.power_consumption = dr.gamma_rate * dr.epsilon_rate

	common_bits := calculateCommonBits(data)

	// Check the most common bit to define the gamma rate and the least common
	// bit to calculate the epsilon rate
	gamma_binary := ""
	epsilon_binary := ""
	for _, bit_count := range common_bits {
		if bit_count.count_0 > bit_count.count_1 {
			gamma_binary += "0"
			epsilon_binary += "1"
		} else {
			gamma_binary += "1"
			epsilon_binary += "0"
		}
	}

	v, err := strconv.ParseInt(gamma_binary, 2, 64)
	if err != nil {
		return nil, err
	} else {
		dr.gamma_rate = v
	}

	v, err = strconv.ParseInt(epsilon_binary, 2, 64)
	if err != nil {
		return nil, err
	} else {
		dr.epsilon_rate = v
	}
	dr.power_consumption = dr.gamma_rate * dr.epsilon_rate

	return dr, nil
}


func main() {
	year := 2021
	day := 3
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

	yellow_submarine, err := generateDiagnosticReport(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Gama rate: %d\n", yellow_submarine.gamma_rate)
	fmt.Printf("Epsilon rate: %d\n", yellow_submarine.epsilon_rate)
	fmt.Printf("Power consumption: %d\n", yellow_submarine.power_consumption)
}