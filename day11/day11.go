package day11

import (
	"log"
	"strconv"
)

const filePath = "input1.txt"

type grid [][]int

func (g *grid) step() (flashes int) {
	g.incrementAll()
	return g.flash()
}

func (g *grid) flash() int {
	for i := 0; i < len(*g); i++ {
		for j := 0; j < len((*g)[i]); j++ {
			if (*g)[i][j] >= 10 {
				(*g)[i][j] = 0
				g.incrementAdjacent(i, j)
				return 1 + g.flash()
			}
		}
	}
	return 0
}

func (g *grid) incrementAll() {
	for i := 0; i < len(*g); i++ {
		for j := 0; j < len((*g)[i]); j++ {
			(*g)[i][j]++
		}
	}
}

func (g *grid) incrementAdjacent(i, j int) {
	for i2 := i - 1; i2 <= i+1; i2++ {
		for j2 := j - 1; j2 <= j+1; j2++ {
			// ignore out of bounds
			if i2 >= len(*g) || i2 < 0 || j2 >= len((*g)[i2]) || j2 < 0 {
				continue
			}

			// ignore itself
			if i2 == i && j2 == j {
				continue
			}

			// ignore if octopus has already flash in this step
			if (*g)[i2][j2] == 0 {
				continue
			}

			(*g)[i2][j2]++
		}
	}
}

func (g *grid) isSynced() bool {
	for i := 0; i < len(*g); i++ {
		for j := 0; j < len((*g)[i]); j++ {
			if (*g)[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func parseInput(lines []string) grid {
	g := make(grid, 0)
	for _, line := range lines {
		energyLine := make([]int, 0)
		for _, r := range line {
			n, err := strconv.Atoi(string(r))
			if err != nil {
				log.Fatal("error parsing input, invalid number")
			}
			energyLine = append(energyLine, n)
		}
		g = append(g, energyLine)
	}
	return g
}

func exe2(lines []string) (step int) {
	g := parseInput(lines)

	for !g.isSynced() {
		step++
		g.step()
	}

	return step
}

func exe1(lines []string) (flashes int) {
	g := parseInput(lines)

	for i := 1; i <= 100; i++ {
		flashes += g.step()
	}

	return
}
