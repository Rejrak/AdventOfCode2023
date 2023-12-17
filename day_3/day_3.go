package main

import (
	"fmt"
	"image"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const (
	day3 = "puzzle.txt"
)

func main() {
	byteSchema, _ := os.ReadFile(day3)
	engineSchema := strings.Fields(string(byteSchema))
	re := regexp.MustCompile(`\d+`)
	
	parts := map[image.Point][]int{}
	partsNumberSum := 0
	gearRatio := 0

	symbolsGrid := map[image.Point]rune{}
	for y, row := range engineSchema {
		for x, col := range row {
			isValidSymbol := col != '.' && !unicode.IsDigit(col)
			if isValidSymbol {
				symbolsGrid[image.Point{x, y}] = col
			}
		}
	}

	neighbours := []image.Point{
		{-1, -1}, {-1, 0},
		{-1, 1}, {0, -1},
		{0, 1}, {1, -1},
		{1, 0}, {1, 1},
	}
	
	for y, row := range engineSchema {
		for _, partIndexes := range re.FindAllStringIndex(row, -1) {
			square := map[image.Point]struct{}{}

			for x := partIndexes[0]; x < partIndexes[1]; x++ {
				for _, d := range neighbours {
					p := image.Point{X: x, Y:y}
					sp := p.Add(d)
					square[sp] = struct{}{}
				}
			}
			
			partNumber, _ := strconv.Atoi(row[partIndexes[0]:partIndexes[1]])			
			for n := range square{
				if _, ok := symbolsGrid[n]; ok{
					parts[n] = append(parts[n], partNumber)
					partsNumberSum += partNumber
				}
			}
		}
	}

	for p, ns := range parts{
		if symbolsGrid[p] == '*' && len(ns) == 2{
			gearRatio += ns[0] * ns[1]
		}
	}

	fmt.Println(partsNumberSum)
	fmt.Println(gearRatio)
}