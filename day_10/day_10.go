package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

const (
	puzzle = "puzzle.txt"
)

func readPuzzle() [][]string {
	bytes, _ := os.ReadFile(puzzle)
	puzzleLines := strings.Split(string(bytes), "\n")

	var puzzle [][]string
	for _, line := range puzzleLines {
		puzzle = append(puzzle, strings.Split(line, ""))
	}

	return puzzle
}

func findStart(puzzle [][]string) (int, int) {
	for y, row := range puzzle {
		for x, tile := range row {
			if tile == `S` {
				return x, y
			}
		}
	}
	return -1, -1
}

func main() {
	puzzle := readPuzzle()
	startX, startY := findStart(puzzle)

	var wg sync.WaitGroup
	var mu sync.Mutex
	paths := make([][][]int, 4)

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(index int, deltaX, deltaY int) {
			defer wg.Done()
			visited := make(map[string]bool)
			paths[index] = explore(puzzle, startX, startY, startX+deltaX, startY+deltaY, visited, &mu)
		}(i, (i-2)%2, (1-i)%2)
	}

	wg.Wait()

	var longestLoop [][]int
	for _, path := range paths {
		if len(path) > 0 && path[len(path)-1][0] == startX && path[len(path)-1][1] == startY {
			if len(path) > len(longestLoop) {
				longestLoop = path
			}
		}
	}

	part1 := len(longestLoop) / 2
	part2 := countTile(longestLoop, puzzle)

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}

func explore(puzzle [][]string, fromX, fromY, toX, toY int, visited map[string]bool, mu *sync.Mutex) [][]int {
	locationKey := fmt.Sprintf("%d,%d", toX, toY)

	mu.Lock()
	if _, ok := visited[locationKey]; ok || toX < 0 || toY < 0 || toX >= len(puzzle[0]) || toY >= len(puzzle) || puzzle[toY][toX] == `.` {
		mu.Unlock()
		return nil
	}
	visited[locationKey] = true
	mu.Unlock()

	if puzzle[toY][toX] == `S` && (toX != fromX || toY != fromY) {
		return [][]int{{toX, toY}}
	}

	nextSteps := getNextSteps(puzzle, toX, toY)
	var path [][]int
	for _, nextStep := range nextSteps {
		nextX, nextY := nextStep[0], nextStep[1]
		if nextX != fromX || nextY != fromY {
			subPath := explore(puzzle, toX, toY, nextX, nextY, visited, mu)
			if subPath != nil {
				path = append([][]int{{toX, toY}}, subPath...)
				break
			}
		}
	}

	mu.Lock()
	delete(visited, locationKey)
	mu.Unlock()

	return path
}

func getNextSteps(puzzle [][]string, x, y int) [][]int {
	var steps [][]int
	switch puzzle[y][x] {
	case `|`:
		steps = append(steps, []int{x, y - 1}, []int{x, y + 1})
	case `-`:
		steps = append(steps, []int{x - 1, y}, []int{x + 1, y})
	case `L`:
		steps = append(steps, []int{x + 1, y}, []int{x, y - 1})
	case `J`:
		steps = append(steps, []int{x - 1, y}, []int{x, y - 1})
	case `7`:
		steps = append(steps, []int{x - 1, y}, []int{x, y + 1})
	case `F`:
		steps = append(steps, []int{x + 1, y}, []int{x, y + 1})
	case `S`:
		steps = append(steps, []int{x, y - 1}, []int{x + 1, y}, []int{x, y + 1}, []int{x - 1, y})
	}
	return steps
}

func countTile(loop [][]int, puzzle [][]string) int {
	loopMap := createLoopMap(loop)
	sX, sY := findStart(puzzle)
	puzzle[sY][sX] = `|` // should not, but still ...

	markOutsideLoopCells(puzzle, loopMap)

	containedCells := findContainedCells(puzzle, loopMap)
	return len(containedCells)
}

func createLoopMap(loop [][]int) map[string]bool {
	loopMap := make(map[string]bool)
	for _, point := range loop {
		key := fmt.Sprintf("%d,%d", point[1], point[0])
		loopMap[key] = true
	}
	return loopMap
}

func markOutsideLoopCells(puzzle [][]string, loopMap map[string]bool) {
	for y, row := range puzzle {
		for x := range row {
			if loopMap[fmt.Sprintf("%d,%d", y, x)] {
				continue
			}
			puzzle[y][x] = `.`
		}
	}
}

func findContainedCells(puzzle [][]string, loopMap map[string]bool) map[string]bool {
	contained := make(map[string]bool)
	for y, row := range puzzle {
		up := false
		inLoop := false
		for x, cell := range row {
			switch cell {
			case "|":
				up = false
				inLoop = !inLoop
			case "F", "L":
				up = cell == "L"
			case "7", "J":
				if cell == "J" != up {
					inLoop = !inLoop
				}
				up = false
			}

			if inLoop && !loopMap[fmt.Sprintf("%d,%d", y, x)] {
				contained[fmt.Sprintf("%d,%d", y, x)] = true
			}
		}
	}
	return contained
}
