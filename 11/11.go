package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/angristan/advent-of-code-2023/utils"
)

func main() {
	input := utils.ParseInput("input.txt")

	image := ConvertRawInputToImage(input)
	image = image.ExpandUniverse()
	fmt.Printf("Part 1: %d\n", image.SumShortestPathBetweenAllGalaxies())
}

type Pixel string

const (
	s Pixel = "."
	G Pixel = "#"
)

type Image [][]Pixel

func ConvertRawInputToImage(rawInput []string) Image {
	image := make(Image, 0)

	for _, line := range rawInput {
		row := make([]Pixel, 0)
		for _, pixel := range line {
			row = append(row, Pixel(string(pixel)))
		}
		image = append(image, row)
	}

	return image
}

func (image Image) ExpandUniverse() Image {
	newImage := make(Image, 0)

	for _, row := range image {
		if !slices.Contains(row, G) {
			newRow := make([]Pixel, 0)
			for i := 0; i < len(row); i++ {
				newRow = append(newRow, s)
			}
			newImage = append(newImage, newRow)
		}
		newImage = append(newImage, row)
	}

	galaxyOnColumn := map[int]bool{}

	for _, row := range newImage {
		for i, pixel := range row {
			if pixel == G {
				galaxyOnColumn[i] = true
			}
		}
	}

	newImage2 := make(Image, 0)
	for _, row := range newImage {
		newRow := make([]Pixel, 0)
		for i := 0; i < len(row); i++ {
			if !galaxyOnColumn[i] {
				newRow = append(newRow, s)
			}
			newRow = append(newRow, row[i])
		}
		newImage2 = append(newImage2, newRow)
	}

	return newImage2
}

type Coords struct {
	X int
	Y int
}

func ComputeShortestBetween(start Coords, end Coords) int {
	deltaX := math.Abs(float64(start.X - end.X))
	deltaY := math.Abs(float64(start.Y - end.Y))

	return int(deltaX + deltaY)
}

type Pair struct {
	Start Coords
	End   Coords
}

func (image Image) SumShortestPathBetweenAllGalaxies() int {
	galaxies := []Coords{}

	for y, row := range image {
		for x, pixel := range row {
			if pixel == G {
				galaxies = append(galaxies, Coords{x, y})
			}
		}
	}

	pairsOfGalaxies := []Pair{}

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			pairsOfGalaxies = append(pairsOfGalaxies, Pair{galaxies[i], galaxies[j]})
		}
	}

	sum := 0
	for _, pair := range pairsOfGalaxies {
		sum += ComputeShortestBetween(pair.Start, pair.End)
	}

	return sum
}
