package day15

import (
	"container/heap"
	"math"

	"github.com/daveseco7/advent-of-code-2021/util"
)

const filePath = "input1.txt"

type coordinate struct {
	x, y int
}

func (c *coordinate) neighbors() []coordinate {
	n := make([]coordinate, 4)
	n[0] = coordinate{c.x - 1, c.y}
	n[1] = coordinate{c.x + 1, c.y}
	n[2] = coordinate{c.x, c.y - 1}
	n[3] = coordinate{c.x, c.y + 1}

	return n
}

type best map[coordinate]int

type state struct {
	best  best
	input map[coordinate]int
}

func parseInput(lines []string, repeater int) (state, int, int) {
	s := state{
		input: make(map[coordinate]int),
		best:  make(best),
	}

	for ry := 0; ry < repeater; ry++ {
		offset := ry
		for rx := 0; rx < repeater; rx++ {
			for y := 0; y < len(lines); y++ {
				for x := 0; x < len(lines[y]); x++ {
					risk := util.MustAtoi(string(lines[y][x]))

					c := coordinate{
						x: x + (rx * len(lines[y])),
						y: y + (ry * len(lines)),
					}

					newRisk := risk + offset
					if newRisk > 9 {
						newRisk = newRisk % 9
					}

					s.input[c] = newRisk
					s.best[c] = math.MaxInt
				}
			}
			offset++
		}
	}

	return s, repeater*len(lines) - 1, repeater*len(lines[0]) - 1
}

func exe2(lines []string) int {
	s, maxX, maxY := parseInput(lines, 5)

	pq := make(PriorityQueue, 0)
	heap.Push(&pq, &Item{
		value:    coordinate{0, 0},
		priority: 0,
	})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		if item.priority < s.best[item.value] {
			s.best[item.value] = item.priority
			for _, n := range item.value.neighbors() {
				if risk, ok := s.input[n]; ok { // if exists and is within bounds
					heap.Push(&pq, &Item{
						value:    n,
						priority: item.priority + risk,
					})
				}
			}
		}
	}

	return s.best[coordinate{
		x: maxX,
		y: maxY,
	}]
}

func exe1(lines []string) int {
	s, maxX, maxY := parseInput(lines, 1)

	pq := make(PriorityQueue, 0)
	heap.Push(&pq, &Item{
		value:    coordinate{0, 0},
		priority: 0,
	})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		if item.priority < s.best[item.value] {
			s.best[item.value] = item.priority
			for _, n := range item.value.neighbors() {
				if risk, ok := s.input[n]; ok { // if exists and is within bounds
					heap.Push(&pq, &Item{
						value:    n,
						priority: item.priority + risk,
					})
				}
			}
		}
	}

	return s.best[coordinate{
		x: maxX,
		y: maxY,
	}]
}
