package main

import (
	"fmt"
	"regexp"

	"github.com/angristan/advent-of-code-2023/utils"
)

func main() {
	input := utils.ParseInput("input.txt")

	m := ConvertRawInputToMap(input)
	fmt.Printf("Part 1: %d\n", m.StepsCountToZZZ())
}

type Direction string

const (
	Left  Direction = "L"
	Right Direction = "R"
)

type Node struct {
	Value string
	Left  string
	Right string
}

type Map struct {
	Directions []Direction
	Nodes      map[string]Node
}

var matchNodesRegex = regexp.MustCompile(`([A-Z]+) = \(([A-Z]+), ([A-Z]+)\)`)

func ConvertRawInputToMap(input []string) Map {
	m := Map{}

	m.Directions = make([]Direction, len(input[0]))
	for i, v := range input[0] {
		m.Directions[i] = Direction(string(v))
	}

	m.Nodes = make(map[string]Node)

	for _, v := range input[2:] {
		matches := matchNodesRegex.FindStringSubmatch(v)
		node := Node{
			Value: matches[1],
			Left:  matches[2],
			Right: matches[3],
		}
		m.Nodes[node.Value] = node
	}

	return m
}

func (m Map) StepsCountToZZZ() int {
	count := 0
	currentNode := m.Nodes["AAA"]

	for currentNode.Value != "ZZZ" {
		if m.Directions[count%len(m.Directions)] == Left {
			currentNode = m.Nodes[currentNode.Left]
		} else {
			currentNode = m.Nodes[currentNode.Right]
		}

		count++
	}

	return count
}
