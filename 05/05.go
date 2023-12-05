package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/angristan/advent-of-code-2023/utils"
)

func main() {
	input := utils.ParseInput("input.txt")

	cards := ConvertInputToAlmanac(input)
	part1Score := cards.GetLowestLocationNumber()
	fmt.Printf("Part 1: %d\n", part1Score)

}

type Range struct {
	DestinationIndex int
	SourceIndex      int
	RangeLength      int
}

type Map struct {
	Ranges []Range
}

type Seed int

type Almanac struct {
	Maps  []Map
	Seeds []Seed
}

var numberRegex = regexp.MustCompile(`\d+`)

func ConvertInputToAlmanac(input []string) Almanac {
	seeds := []Seed{}

	seedsString := numberRegex.FindAllString(input[0], -1)

	for _, seedString := range seedsString {
		seed, err := strconv.Atoi(seedString)
		if err != nil {
			panic(err)
		}

		seeds = append(seeds, Seed(seed))
	}

	maps := []Map{}

	currentMap := Map{}
	for _, line := range input[2:] {
		if strings.Contains(line, "map") || line == "" {
			if len(currentMap.Ranges) > 0 {
				maps = append(maps, currentMap)
				currentMap = Map{}
			}
			continue
		}

		indices := numberRegex.FindAllString(line, -1)
		destinationIndex, err := strconv.Atoi(indices[0])
		if err != nil {
			panic(err)
		}

		sourceIndex, err := strconv.Atoi(indices[1])
		if err != nil {
			panic(err)
		}

		rangeLength, err := strconv.Atoi(indices[2])
		if err != nil {
			panic(err)
		}

		currentMap.Ranges = append(currentMap.Ranges, Range{
			DestinationIndex: destinationIndex,
			SourceIndex:      sourceIndex,
			RangeLength:      rangeLength,
		})
	}

	maps = append(maps, currentMap)

	return Almanac{
		Maps:  maps,
		Seeds: seeds,
	}
}

func (almanac Almanac) GetSeedsLocations() []int {
	locations := make([]int, 0)

	for _, seed := range almanac.Seeds {
		nextIndex := int(seed)
	nextMap:
		for _, m := range almanac.Maps {
			for _, r := range m.Ranges {
				if nextIndex >= r.SourceIndex && nextIndex < r.SourceIndex+r.RangeLength {
					nextIndex = r.DestinationIndex + int(nextIndex) - r.SourceIndex
					continue nextMap
				}
			}
		}

		location := nextIndex
		locations = append(locations, location)
	}

	return locations
}

func (almanac Almanac) GetLowestLocationNumber() int {
	locations := almanac.GetSeedsLocations()
	slices.Sort(locations)

	return locations[0]
}
