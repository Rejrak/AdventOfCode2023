package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


const (
	day5 = "puzzle.txt"
)

func readAlmanac(strGames []string) ([]int, map[string][][]int){
	seeds := []int{}
	maps := map[string][][]int{}

	for _, line := range strGames{
		if strings.Contains(line, "seeds:"){
			seeds = seedsMe(strings.ReplaceAll(line, "seeds: ", ""))
			continue
		}

		almanac := strings.Split(line, "\n")
		name := strings.Split(almanac[0], " ")[0]
		
		for i := 1; i < len(almanac); i++{
			maps[name] = append(maps[name], seedsMe(almanac[i]))
		} 
	}
	return seeds, maps
}

func part1(strGames []string){
	pt1 := 0
	seeds := []int{}
	maps := map[string][][]int{}

	seeds, maps = readAlmanac(strGames)

	for _, seed := range seeds{
		ts := calcScore(seed, maps)
		if pt1 == 0{
			pt1 = ts
			continue
		}
		if ts < pt1 {
			pt1 = ts 
		}
	}

	fmt.Println(pt1)
}

func part2(strGames []string){
	pt2 := 0

	seeds := []int{}
	maps := map[string][][]int{}

	seeds, maps = readAlmanac(strGames)
	// oh no
	for i := 0; i < len(seeds); i+=2{
		// oh fugg
		for j := seeds[i]; j < seeds[i] + seeds[i+1]; j++{
			// pls don't
			ts := calcScore(j, maps)
			if pt2 == 0{
				pt2 = ts
				continue
			}
			if ts < pt2 {
				pt2 = ts 
			}
		}
	}
	fmt.Println(pt2)
}

func main(){
	byteMap, _ := os.ReadFile(day5)
	puzzle := string(byteMap)
	strGames := strings.Split(puzzle, "\n\n")
	part1(strGames)
	part2(strGames)
}

func calcDest(source int, mapping [][]int) int{
	for _, m := range mapping{
		if m[1] <= source && source < m[1] + m[2]{
			return m[0] + (source - m[1])
		}
	}
	return source
}

func calcScore (seed int, maps map[string][][]int) int{
	soil := calcDest(seed, maps["seed-to-soil"])
	fertilizer := calcDest(soil, maps["soil-to-fertilizer"])
	water := calcDest(fertilizer, maps["fertilizer-to-water"])
	light := calcDest(water, maps["water-to-light"])
	temperature := calcDest(light, maps["light-to-temperature"])
	humidity := calcDest(temperature, maps["temperature-to-humidity"])
	location := calcDest(humidity, maps["humidity-to-location"])
	return location
}


func seedsMe(s string) []int{
	vals := []int{}
	for _, v := range strings.Split(s, " "){
		i, err := strconv.Atoi(v)
		if err != nil{
			// should not but
			panic(err)
		}
		vals = append(vals, i)
	}
	return vals
}
