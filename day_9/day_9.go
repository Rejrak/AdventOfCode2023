package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	day9 = "puzzle.txt"
)

func readAndParsePuzzle() [][]int {
	byteMap, _ := os.ReadFile(day9)
	puzzle := string(byteMap)

	var history []int
	var histories [][]int
	lines := strings.Split(puzzle, "\n")

	for _, line := range lines {
		strVals := strings.Split(line, " ")
		for val := range strVals {
			i, _ := strconv.Atoi(strVals[val])
			history = append(history, i)
		}
		histories = append(histories, history)
		history = []int{}
	}

	return histories
}

func sequenceOfDifference(history []int) []int {
	var deltas []int = make([]int, len(history)-1)
	for i := 1; i < len(history); i++ {
		delta := history[i] - history[i-1]
		deltas[i-1] = delta
	}
	return deltas
}

func predictionMap(predictions *[][]int, deltas []int, level int) {
	noChanges := true
	for _, delta := range deltas {
		if delta != 0 {
			noChanges = false
			break
		}
	}

	if noChanges {
		return
	}

	deltas = sequenceOfDifference((*predictions)[level])
	*predictions = append(*predictions, deltas)
	predictionMap(predictions, deltas, level+1)
}

func main() {
	histories := readAndParsePuzzle()
	pt1 := calcScorePt1(histories)
	pt2 := calcScorePt2(histories)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", pt1, pt2)
}

func extrapolateBackward(predictions [][]int) int {
	lastLevel := len(predictions) - 1

	shiftedValue := []int{0}
	shiftedValue = append(shiftedValue, (predictions)[lastLevel]...) // predictions[lastLevel]...)

	predictions[lastLevel] = shiftedValue

	for i := lastLevel - 1; i >= 0; i-- {
		nextValue := predictions[i][0]
		currentValue := predictions[i+1][0]
		shiftedValue = []int{nextValue - currentValue}
		shiftedValue = append(shiftedValue, predictions[i]...)
		predictions[i] = shiftedValue
	}
	return predictions[0][0]
}

func makePrediction(predictions [][]int) int {
	lastLevel := len(predictions) - 1
	predictions[lastLevel] = append(predictions[lastLevel], 0)
	for i := lastLevel - 1; i >= 0; i-- {
		nextValue := predictions[i][len(predictions[i])-1]
		currentValue := predictions[i+1][len(predictions[i+1])-1]
		predictions[i] = append(predictions[i], nextValue+currentValue)
	}
	return predictions[0][len(predictions[0])-1]
}

func calcScorePt1(histories [][]int) int {
	var predictedValue int = 0
	var predictions [][]int

	for _, history := range histories {
		predictions = append(predictions, history)
		predictionMap(&predictions, sequenceOfDifference(history), 0)
		predictedValue += makePrediction(predictions)
		predictions = [][]int{}
	}
	return predictedValue
}

func calcScorePt2(histories [][]int) int {
	var extrapolatedValue int = 0
	var predictions [][]int

	for _, history := range histories {
		predictions = append(predictions, history)
		predictionMap(&predictions, sequenceOfDifference(history), 0)
		extrapolatedValue += extrapolateBackward(predictions)
		predictions = [][]int{}
	}

	return extrapolatedValue
}
