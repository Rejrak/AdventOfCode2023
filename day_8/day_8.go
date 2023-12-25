package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	day8  = "puzzle.txt"
	test2 = "test2.txt"
)

var directionMap = map[string]int{
	"L": 0,
	"R": 1,
}

var coords map[string][]string = map[string][]string{}

type node struct {
	Name      string
	leftNode  *node
	rightNode *node
	z         bool
}

func readAndParsePuzzle() (string, map[string][]string) {
	byteMap, _ := os.ReadFile(day8)
	puzzle := string(byteMap)	
	directions := strings.Split(puzzle, "\n")[0]
	strCoords := strings.Split(puzzle, "\n")[2:]

	
	//pt1
	for _, s := range strCoords {
		position := s[0:3]
		left := s[7:10]
		right := s[12:15]
		coords[position] = append(coords[position], left, right)
	}

	return directions, coords
}

func initializeNodesPt2(coords map[string][]string) map[string]*node {
	nodes := map[string]*node{}

	for k := range coords {
		nodes[k] = &node{Name: k, z: strings.Contains(k, `Z`)}
	}
	for k := range nodes {
		nodes[k].leftNode = nodes[coords[k][0]]
		nodes[k].rightNode = nodes[coords[k][1]]
	}

	return nodes
}

func foundActiveNodesPt2(nodes map[string]*node) []*node {

	activeNodes := []*node{}
	for k, v := range nodes {
		if strings.Contains(k, `A`) {
			activeNodes = append(activeNodes, v)
		}
	}
	return activeNodes
}

func main() {
	pt1 := 0
	directions, coords := readAndParsePuzzle()
	
	//pt1
	pt1 = followDirection(0, directions, `AAA`, 0)

	//pt2
	nodes := initializeNodesPt2(coords)
	activeNodes := foundActiveNodesPt2(nodes)

	iterations := []int{}
	for _, node := range activeNodes{
		iterations = append(iterations, followDirection(0, directions, node.Name, 0))
	}
	lcm := iterations[0]
	for _, n := range iterations {
		lcm = lcm * n / gcd(lcm, n)
	}
	pt2 := lcm
	
	fmt.Printf("Part 1: %d\nPart 2: %d\n", pt1, pt2)
}

func gcd(n1 int, n2 int) int {
	if n2 != 0 {
		return gcd(n2, n1%n2)
	} else {
		return n1
	}
}

func followDirection(step int, direction string, humanCoord string, cnt int) int {
	if strings.Contains(humanCoord, `Z`) {
		return cnt
	}

	if step%len(direction) == 0 {
		step = 0
	}

	nextStep := directionMap[string(direction[step])]
	nextHumanCoord := coords[humanCoord][nextStep]

	return followDirection(step+1, direction, nextHumanCoord, cnt+1)
}
