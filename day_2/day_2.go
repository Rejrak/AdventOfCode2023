package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	puzzle = "puzzle.txt"
)

func main(){
	bytesGames, _ := os.ReadFile(puzzle)
	strGames := strings.Split(string(bytesGames), "\n")
		
	possibleGames := 0
	cubesPower := 0

	re := regexp.MustCompile(`(\d+) (\w+)`)
	
	for i, s := range strGames {
		cubes := map[string]int{}

		for _, m := range re.FindAllStringSubmatch(s, -1) {
			n, _ := strconv.Atoi(m[1])
			color := m[2]
			cubes[color] = max(cubes[color], n)
		}

		checkPossibleGames := cubes["red"] <= 12 && cubes["green"] <= 13 && cubes["blue"] <= 14 
		if checkPossibleGames {
			possibleGames += i + 1
		}

		cubesPower += cubes["red"] * cubes["green"] * cubes["blue"]
	}
	fmt.Printf("Possible Games => %d\n  Cubes Power  => %d\n", possibleGames, cubesPower)
}