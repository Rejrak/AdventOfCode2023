package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	puzzle = "puzzle.txt"
)

// func distManhattan(a, b [2]int) int {
// 	return abs(b[0]-a[0]) + abs(b[1]-a[1])
// }

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }

func readPuzzle() []string {
	bytes, _ := os.ReadFile(puzzle)
	puzzleLines := strings.Split(string(bytes), "\n")
	return puzzleLines
}

func expandedRow(lines []string) map[int]bool {
	expandedRow := map[int]bool{}
	for y, row := range lines {
		expanded := true
		for _, v := range row {
			if v == '#' {
				expanded = false
				break
			}
		}
		expandedRow[y] = expanded
	}
	return expandedRow
}
func expandedCol(lines []string) map[int]bool {
	expandedCol := map[int]bool{}
	for x := range lines[0] {
		expanded := true
		for _, row := range lines {
			if row[x] == '#' {
				expanded = false
				break
			}
		}
		expandedCol[x] = expanded
	}
	return expandedCol
}

func main() {
	lines := readPuzzle()

	var coords [][2]int
	for i, line := range lines {
		for j, ch := range line {
			if ch == '#' {
				coords = append(coords, [2]int{i, j})
			}
		}
	}
	expandedRow := expandedRow(lines)
	expandedCol := expandedCol(lines)

	var part1 int64 = 0
	var part2 int64 = 0
	for idx1, coord1 := range coords {
		for idx2, coord2 := range coords {
			if idx1 < idx2 {
				d1, d2 := dist(coord1, coord2, expandedRow, expandedCol)
				part1 += d1
				part2 += d2
			}
		}
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}

func dist(a, b [2]int, emptyRow, emptyCol map[int]bool) (int64, int64) {
	var part1 int64 = 0
	var part2 int64 = 0
	for i := a[0]; i != b[0]; {
		if i < b[0] {
			i++
		} else {
			i--
		}
		part1++
		part2++
		if emptyRow[i] {
			part1++
			part2 += int64(math.Pow(10, 6)) - 1

		}
	}
	for j := a[1]; j != b[1]; {
		if j < b[1] {
			j++
		} else {
			j--
		}
		part1++
		part2++
		if emptyCol[j] {
			part1++
			part2 += int64(math.Pow(10, 6)) - 1
		}
	}
	return part1, part2
}
