package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	day6 = "puzzle.txt"
)

var digitsRe = regexp.MustCompile(`\d+`)

func main(){
	byteMap, _ := os.ReadFile(day6)
	puzzle := string(byteMap)
	lines := strings.Split(puzzle, "\n")
	
	times := digitsRe.FindAllString(lines[0], -1)
	distances := digitsRe.FindAllString(lines[1], -1)
	
	pt1 := part1(times, distances)
	pt2 := part2(times, distances)

	fmt.Printf("Part 1: %d\n",pt1)
	fmt.Printf("Part 2: %d\n",pt2)
	
}

func part1(times []string, distances []string) (score int){
	score = 1
	lenRecords := len(distances) // or len(times)


	for i := 0; i < lenRecords; i++{
		t, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(distances[i])
		count := 0
		for ht := 1; ht < t; ht++{
			if ht * (t - ht) > d{
				count ++
			}
		}
		score = score * count
	}
	return score

}

func part2(times []string, distances []string) (score int){
	score = 1

	ts := strings.Join(times, "")
	ds := strings.Join(distances, "")

	t, _ := strconv.Atoi(ts)
	d, _ := strconv.Atoi(ds)
	count := 0
	for ht := 1; ht < t; ht++{
		if ht * (t - ht) > d{
			count ++
		}
	}
	
	score = score * count
	return score

}