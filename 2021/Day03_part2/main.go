// Copyright © 2021 Paulo Vital <paulo@vital.eng.br>
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Advent of Code 2021 - Day 03
// Binary Diagnostic - Part 2

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type diagnosticReport struct {
	oxygen_rate int64
	co2_rate int64
	life_support int64
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

func getMostCommonValue(bit_count bitCount) (string) {
	if bit_count.count_0 > bit_count.count_1 {
		return "0"
	} else {
		return "1"
	}
}

func getLeastCommonValue(bit_count bitCount) (string) {
	if bit_count.count_0 > bit_count.count_1 {
		return "1"
	} else {
		return "0"
	}
}


func getOxygenRating(data []string, bit_pos int) ([]string) {
	common_bits := calculateCommonBits(data)
	mcv := getMostCommonValue(common_bits[bit_pos])

	// Generate a slice with the strings matching the mcv
	mcvs := make([]string, 0, len(data))
	for _, val := range data {
		bits := strings.Split(val, "")
		if bits[bit_pos] == mcv {
			mcvs = append(mcvs, val)
		}
	}
	return mcvs
}

func getCO2Rating(data []string, bit_pos int) ([]string) {
	common_bits := calculateCommonBits(data)
	lcv := getLeastCommonValue(common_bits[bit_pos])

	// Generate a slice with the strings matching the mcv
	lcvs := make([]string, 0, len(data))
	for _, val := range data {
		bits := strings.Split(val, "")
		if bits[bit_pos] == lcv {
			lcvs = append(lcvs, val)
		}
	}
	return lcvs
}

func generateDiagnosticReport(data []string) (*diagnosticReport, error) {
	dr := &diagnosticReport{}
	dr.oxygen_rate = 0
	dr.co2_rate = 0
	dr.life_support = dr.oxygen_rate * dr.co2_rate

	// Calculate the oxygen generator rating
	oxygen_rate := make([]string, len(data))
	copy(oxygen_rate, data)

	for i := 0; len(oxygen_rate) > 1; i++ {
		oxygen_rate = getOxygenRating(oxygen_rate, i)
	}

	output, err := strconv.ParseInt(oxygen_rate[0], 2, 64)
	if err != nil {
		return nil, err
	}

	dr.oxygen_rate = output

	// Calculate the CO2 scrubber rating
	co2_rate := make([]string, len(data))
	copy(co2_rate, data)

	for i := 0; len(co2_rate) > 1; i++ {
		co2_rate = getCO2Rating(co2_rate, i)
	}

	output, err = strconv.ParseInt(co2_rate[0], 2, 64)
	if err != nil {
		return nil, err
	}
 	dr.co2_rate = output

	// Calculate the life support rating
	dr.life_support = dr.oxygen_rate * dr.co2_rate

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

	fmt.Printf("Oxygen generator rating: %d\n", yellow_submarine.oxygen_rate)
	fmt.Printf("CO2 scrubber rating: %d\n", yellow_submarine.co2_rate)
	fmt.Printf("Life support rating: %d\n", yellow_submarine.life_support)
}