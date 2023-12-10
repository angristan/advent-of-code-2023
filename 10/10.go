package main

import (
	"fmt"
	"math"

	"github.com/angristan/advent-of-code-2023/utils"
)

func main() {
	input := utils.ParseInput("input.txt")

	grid := ConvertRawInputToSurfacePipes(input)
	fmt.Printf("Part 1: %d\n", grid.GetFurthestPipeFromStartStepsCount())

	enclosedTiles := grid.GetEnclosedTiles()
	fmt.Printf("Part 2: %d\n", len(enclosedTiles))
}

type TileType string

const (
	PipeVertical   TileType = "|"
	PipeHorizontal TileType = "-"
	PipeBendL      TileType = "L"
	PipeBendJ      TileType = "J"
	PipeBend7      TileType = "7"
	PipeBendF      TileType = "F"
	PipeGround     TileType = "."
	Start          TileType = "S"
)

type Tile struct {
	Type           TileType
	CoordX, CoordY int
}

// https://www.youtube.com/watch?v=4-J4duzP8Ng
type TheGrid [][]Tile

func ConvertRawInputToSurfacePipes(rawInput []string) TheGrid {
	surfacePipes := make(TheGrid, len(rawInput))
	for y, row := range rawInput {
		surfacePipes[y] = make([]Tile, len(row))
		for x, pipe := range row {
			if len(string(pipe)) != 1 {
				panic("Invalid pipe type")
			}
			surfacePipes[y][x] = Tile{
				Type:   TileType(string(pipe)),
				CoordX: x,
				CoordY: y,
			}
		}
	}

	return surfacePipes
}

func (sp TheGrid) GetStartPipe() Tile {
	for _, row := range sp {
		for _, pipe := range row {
			if pipe.Type == Start {
				return pipe
			}
		}
	}

	panic("No start pipe found")
}

func (grid TheGrid) GetAdjacentTiles(tile Tile) []Tile {
	adjacentTiles := []Tile{}

	if tile.CoordX > 0 {
		adjacentTiles = append(adjacentTiles, grid[tile.CoordY][tile.CoordX-1])
	}
	if tile.CoordX < len(grid[0])-1 {
		adjacentTiles = append(adjacentTiles, grid[tile.CoordY][tile.CoordX+1])
	}
	if tile.CoordY > 0 {
		adjacentTiles = append(adjacentTiles, grid[tile.CoordY-1][tile.CoordX])
	}
	if tile.CoordY < len(grid)-1 {
		adjacentTiles = append(adjacentTiles, grid[tile.CoordY+1][tile.CoordX])
	}

	return adjacentTiles
}

func (grid TheGrid) GetAdjacentPipes(tile Tile) []Tile {
	adjacentPipes := []Tile{}

	for _, adjacentTile := range grid.GetAdjacentTiles(tile) {
		if adjacentTile.Type != PipeGround {
			adjacentPipes = append(adjacentPipes, adjacentTile)
		}
	}

	return adjacentPipes
}

func (grid TheGrid) GetConnectedPipes(tile Tile) []Tile {
	connectedPipes := []Tile{}

	for _, adjacentPipe := range grid.GetAdjacentPipes(tile) {
		switch adjacentPipe.CoordX {
		case tile.CoordX - 1:
			switch tile.Type {
			case Start, PipeHorizontal, PipeBendJ, PipeBend7:
				switch adjacentPipe.Type {
				case PipeHorizontal, PipeBendF, PipeBendL:
					connectedPipes = append(connectedPipes, adjacentPipe)
				}
			}
		case tile.CoordX + 1:
			switch tile.Type {
			case Start, PipeHorizontal, PipeBendL, PipeBendF:
				switch adjacentPipe.Type {
				case PipeHorizontal, PipeBendJ, PipeBend7:
					connectedPipes = append(connectedPipes, adjacentPipe)
				}
			}
		}

		switch adjacentPipe.CoordY {
		case tile.CoordY - 1:
			switch tile.Type {
			case Start, PipeVertical, PipeBendL, PipeBendJ:
				switch adjacentPipe.Type {
				case PipeVertical, PipeBendF, PipeBend7:
					connectedPipes = append(connectedPipes, adjacentPipe)
				}
			}
		case tile.CoordY + 1:
			switch tile.Type {
			case Start, PipeVertical, PipeBend7, PipeBendF:
				switch adjacentPipe.Type {
				case PipeVertical, PipeBendL, PipeBendJ:
					connectedPipes = append(connectedPipes, adjacentPipe)
				}
			}
		}
	}

	switch len(connectedPipes) {
	case 1:
		connectedPipes = append(connectedPipes, grid.GetStartPipe())
	case 2:
	default:
		panic("Invalid number of connected pipes")
	}

	return connectedPipes
}

func (grid TheGrid) GetFurthestPipeFromStartStepsCount() int {
	steps := 0

	lastPipe := Tile{}
	var currentPipe Tile
	nextPipe := grid.GetStartPipe()
	for {
		currentPipe = nextPipe
		connectedPipes := grid.GetConnectedPipes(currentPipe)

		if connectedPipes[0] != lastPipe {
			nextPipe = connectedPipes[0]
		} else {
			nextPipe = connectedPipes[1]
		}
		lastPipe = currentPipe

		steps++

		if nextPipe.Type == Start {
			break
		}
	}

	return int(math.Round(float64(steps) / 2))
}

type Coord struct {
	X, Y int
}

func (grid TheGrid) GetLoopTiles() map[Coord]Tile {
	loopMap := map[Coord]Tile{}

	lastPipe := Tile{}
	var currentPipe Tile
	nextPipe := grid.GetStartPipe()
	for {
		currentPipe = nextPipe
		connectedPipes := grid.GetConnectedPipes(currentPipe)

		if connectedPipes[0] != lastPipe {
			nextPipe = connectedPipes[0]
		} else {
			nextPipe = connectedPipes[1]
		}
		lastPipe = currentPipe

		loopMap[Coord{X: currentPipe.CoordX, Y: currentPipe.CoordY}] = currentPipe

		if nextPipe.Type == Start {
			break
		}
	}

	return loopMap
}

func (grid TheGrid) GetEnclosedTiles() []Tile {
	loopMap := grid.GetLoopTiles()
	enclosedTiles := []Tile{}

	for y, row := range grid {
		for x, tile := range row {
			// If already in the loop: skip
			if _, ok := loopMap[Coord{X: x, Y: y}]; ok {
				continue
			}

			// If touches side of grid: outside
			if x == 0 || x == len(grid[0])-1 || y == 0 || y == len(grid)-1 {
				continue
			}

			// Take all tiles to the neareast side of the grid
			var tilesToSide []Tile
			if x < len(grid)/2 {
				tilesToSide = grid[y][:x]
			} else {
				tilesToSide = grid[y][x+1:]
			}

			// Drop all PipeHorizontal
			var tilesToSideWithoutPipeHorizontal []Tile
			for _, tileToSide := range tilesToSide {
				if tileToSide.Type != PipeHorizontal {
					tilesToSideWithoutPipeHorizontal = append(tilesToSideWithoutPipeHorizontal, tileToSide)
				}
			}

			// Drop all tiles not belonging to the loop
			var loopTilesToSide []Tile
			for _, tileToSide := range tilesToSideWithoutPipeHorizontal {
				if _, ok := loopMap[Coord{X: tileToSide.CoordX, Y: tileToSide.CoordY}]; ok {
					loopTilesToSide = append(loopTilesToSide, tileToSide)
				}
			}

			// Replace L+7 and F+J by |
			var layeredLoopTilesToSide []Tile
			for i, tileToSide := range loopTilesToSide {
				if i < len(loopTilesToSide)-1 {
					if tileToSide.Type == PipeBendL && loopTilesToSide[i+1].Type == PipeBend7 {
						layeredLoopTilesToSide = append(layeredLoopTilesToSide, Tile{
							Type: PipeVertical,
						})
						continue
					}

					if tileToSide.Type == PipeBendF && loopTilesToSide[i+1].Type == PipeBendJ {
						layeredLoopTilesToSide = append(layeredLoopTilesToSide, Tile{
							Type: PipeVertical,
						})
						continue
					}
				}

				if i >= 1 {
					// If replaced by | in previous iteration: skip
					if loopTilesToSide[i-1].Type == PipeBendL && tileToSide.Type == PipeBend7 {
						continue
					}
					if loopTilesToSide[i-1].Type == PipeBendF && tileToSide.Type == PipeBendJ {
						continue
					}
				}

				layeredLoopTilesToSide = append(layeredLoopTilesToSide, tileToSide)
			}

			//TODO: S, but it works without it for some reason 🥸

			// Count the number of layers of the loop until the side
			// If odd: inside
			// If even: outside
			if len(layeredLoopTilesToSide)%2 == 1 {
				enclosedTiles = append(enclosedTiles, tile)
			}
		}
	}

	return enclosedTiles
}
