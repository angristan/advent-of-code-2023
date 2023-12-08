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

	fmt.Printf("Part 2: %d\n", m.StepsCountToEndingZGhostMode())
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
	Directions       []Direction
	Nodes            map[string]Node
	EndingANodesKeys []string
}

var matchNodesRegex = regexp.MustCompile(`(\S+)\s\=\s\((\S+),\s(\S+)\)`)

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

		if node.Value[2] == 'A' {
			m.EndingANodesKeys = append(m.EndingANodesKeys, node.Value)
		}
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

func (m Map) StepsCountToEndingZGhostMode() int {
	iterationsNeededToEndZNode := make(map[string]int)

	for _, nodeKey := range m.EndingANodesKeys {
		currentNode := m.Nodes[nodeKey]
		iterationCount := 0

		for currentNode.Value[2] != 'Z' {
			if m.Directions[iterationCount%len(m.Directions)] == Left {
				currentNode = m.Nodes[currentNode.Left]
			} else {
				currentNode = m.Nodes[currentNode.Right]
			}

			iterationCount++
		}

		iterationsNeededToEndZNode[nodeKey] = iterationCount
	}

	// LCM of all iterationsValues needed to end Z nodes
	var iterationsValues []int
	for _, v := range iterationsNeededToEndZNode {
		iterationsValues = append(iterationsValues, v)
	}

	lcm := LCM(iterationsValues[0], iterationsValues[1], iterationsValues[2:]...)

	return lcm
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
