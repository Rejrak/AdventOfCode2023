package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// exploring go trees

const (
	puzzle = "puzzle.txt"
)

type TreeNode struct {
	Value  int
	Children []*TreeNode
}

func buildTree(data map[int][]int, rootKey int) *TreeNode {
	root := &TreeNode{Value: rootKey}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		children, ok := data[currentNode.Value]
		if !ok {
			continue
		}

		for _, childValue := range children {
			childNode := &TreeNode{Value: childValue}
			currentNode.Children = append(currentNode.Children, childNode)
			queue = append(queue, childNode)
		}
	}

	return root
}

func countAllNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	scratchCards := 1 // Conta il nodo corrente

	for _, child := range root.Children {
		scratchCards += countAllNodes(child)
	}

	return scratchCards
}

func main() {
	bytesCards, _ := os.ReadFile(puzzle)
	cards := strings.Split(string(bytesCards), "\n")
	re := regexp.MustCompile(`\d+`)
	points := 0
	winningCards := map[int][]int{}
	
	for i := 0; i < len(cards); i++{
		winningCards[i+1]= []int{}
	} 
	
	scratchCards := 0
	for _, card := range cards {
		cardNumber, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(strings.Split(card, ": ")[0], "d")[1], " ", ""))
		cardValues := strings.Split(strings.Split(card, ": ")[1], "|") // extract cardValues my number | winnin numbers
		patternMatching := ""
		
		myNumbers := re.FindAllStringSubmatch(cardValues[0], -1)
		for i := range myNumbers {
			patternMatching += `\b` + myNumbers[i][0] + `\b`
			if i >= 0 && i < len(myNumbers) - 1 {
				patternMatching += "|"
			}
		}
		
		matches := func (text string, pattern string) int {
			re := regexp.MustCompile(pattern)
			matches := re.FindAllString(text, -1)
			return len(matches)
		}(cardValues[1], patternMatching)

		
		k := 0
		var copies []int = make([]int, matches)
		for i := cardNumber; i <  cardNumber + matches; i++{
			copies[k] = i + 1
			k++
		}
		winningCards[cardNumber] = copies
		points += int(math.Pow(2, float64(matches-1)))
	}

	for cardNumber := range winningCards{
		root := buildTree(winningCards, cardNumber)
		scratchCards += countAllNodes(root)
	}
	fmt.Println(points)
	fmt.Println(scratchCards)
}

