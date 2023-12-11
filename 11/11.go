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
	fmt.Printf("Part 1: %d\n", image.SumShortestPathBetweenAllGalaxies(2))

	fmt.Printf("Part 2: %d\n", image.SumShortestPathBetweenAllGalaxies(1000000))
}

type Pixel string

const (
	s Pixel = "."
	G Pixel = "#"
)

type Image [][]Pixel

type Coords struct {
	X int
	Y int
}

type Pair struct {
	Start Coords
	End   Coords
}

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

func ComputeShortestBetween(start Coords, end Coords) int {
	deltaX := math.Abs(float64(start.X - end.X))
	deltaY := math.Abs(float64(start.Y - end.Y))

	return int(deltaX + deltaY)
}

func (image Image) GetGalaxies() []Coords {
	galaxies := []Coords{}

	for y, row := range image {
		for x, pixel := range row {
			if pixel == G {
				galaxies = append(galaxies, Coords{x, y})
			}
		}
	}

	return galaxies
}

func GetPairsOfGalaxies(galaxies []Coords) []Pair {
	pairsOfGalaxies := []Pair{}

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			pairsOfGalaxies = append(pairsOfGalaxies, Pair{galaxies[i], galaxies[j]})
		}
	}

	return pairsOfGalaxies
}

func (image Image) SumShortestPathBetweenAllGalaxies(expensionFactor int) int {
	galaxies := image.GetGalaxies()
	expendedRowsIndexes := image.ExpandedRowsIndexes()
	expendedColumnsIndexes := image.ExpandedColumnsIndexes()
	pairsOfGalaxies := GetPairsOfGalaxies(galaxies)

	sum := 0
	for _, pair := range pairsOfGalaxies {
		sum += ComputeShortestBetween(pair.Start, pair.End)

		maxY := slices.Max([]int{pair.Start.Y, pair.End.Y})
		minY := slices.Min([]int{pair.Start.Y, pair.End.Y})
		maxX := slices.Max([]int{pair.Start.X, pair.End.X})
		minX := slices.Min([]int{pair.Start.X, pair.End.X})

		// Take expension into account
		traversedRows := []int{}
		for _, rowIndex := range expendedRowsIndexes {
			if rowIndex >= minY && rowIndex <= maxY {
				traversedRows = append(traversedRows, rowIndex)
			}
		}

		traversedColumns := []int{}
		for _, columnIndex := range expendedColumnsIndexes {
			if columnIndex >= minX && columnIndex <= maxX {
				traversedColumns = append(traversedColumns, columnIndex)
			}
		}

		sum += (len(traversedRows) + len(traversedColumns)) * (expensionFactor - 1)
	}

	return sum
}

func (image Image) ExpandedRowsIndexes() []int {
	expandedRowsIndexes := []int{}

	for i, row := range image {
		if !slices.Contains(row, G) {
			expandedRowsIndexes = append(expandedRowsIndexes, i)
		}
	}

	return expandedRowsIndexes
}

func (image Image) ExpandedColumnsIndexes() []int {
	expandedColumnsIndexes := []int{}

	galaxyOnColumn := map[int]bool{}

	for _, row := range image {
		for i, pixel := range row {
			if pixel == G {
				galaxyOnColumn[i] = true
			}
		}
	}

	for i := 0; i < len(image[0]); i++ {
		if !galaxyOnColumn[i] {
			expandedColumnsIndexes = append(expandedColumnsIndexes, i)
		}
	}

	return expandedColumnsIndexes
}
