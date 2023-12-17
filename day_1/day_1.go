package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	day1 = "puzzle.txt"
	digits = "123456789"
	zero   = '0'
)

func calibrationValues(codes []string, r *strings.Replacer) (result int) {
	for _, code := range codes {
		code = r.Replace(r.Replace(code))

		idxFirstDigit := strings.IndexAny(code, digits)
		firstDigit := int(code[idxFirstDigit] - zero)

		idxLastDigit := strings.LastIndexAny(code, digits)
		lastDigit := int(code[idxLastDigit] - zero)
		
		result += (10 * firstDigit) + lastDigit
	}
	return
}

func main() {
	bytesCode, _ := os.ReadFile(day1)
	codes := strings.Fields(string(bytesCode))
	
	r0 := strings.NewReplacer()
	r1 := strings.NewReplacer(
		"one", "o1e", "two", "t2o", "three", "t3e",
		"four", "f4r", "five", "f5e", "six", "s6x",
		"seven", "s7n", "eight", "e8t", "nine", "n9e",
	)
	
	fmt.Println(calibrationValues(codes, r0))
	fmt.Println(calibrationValues(codes, r1))
}
