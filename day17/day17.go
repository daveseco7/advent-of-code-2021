package day17

import (
	"log"
	"regexp"
	"strconv"

	"github.com/daveseco7/advent-of-code-2021/util"
)

const filePath = "input1.txt"

const targetCoordinatesRegex = `target area: x=(-?\d+)..(-?\d+), y=(-?\d+)..(-?\d+)`

type target struct {
	minX, maxX int
	minY, maxY int
}

func (t *target) onTarget(p probe) bool {
	return p.x >= t.minX && p.x <= t.maxX && p.y >= t.minY && p.y <= t.maxY
}

func (t *target) simulate(x, y int) (vp probe, valid bool) {
	p := probe{dx: x, dy: y, x: 0, y: 0}
	for p.y >= t.minY && p.x <= t.maxX {
		p.step()
		if t.onTarget(p) {
			return p, true
		}
	}
	return vp, false
}

type probe struct {
	x, y   int
	dx, dy int
	maxY   int
}

func (p *probe) step() {
	p.dy--
	p.x += p.dx
	p.y += p.dy
	if p.dx > 0 {
		p.dx--
	} else {
		p.dx = 0
	}
	p.maxY = util.Max([]int{p.maxY, p.y})
}

func parseInput(lines []string) target {
	r := regexp.MustCompile(targetCoordinatesRegex)
	result := r.FindAllStringSubmatch(lines[0], -1)

	minX, err := strconv.Atoi(result[0][1])
	if err != nil {
		log.Fatal(" error parsing target coordinate min x")
	}
	maxX, err := strconv.Atoi(result[0][2])
	if err != nil {
		log.Fatal(" error parsing target coordinate max x")
	}
	minY, err := strconv.Atoi(result[0][3])
	if err != nil {
		log.Fatal(" error parsing target coordinate min y")
	}
	maxY, err := strconv.Atoi(result[0][4])
	if err != nil {
		log.Fatal(" error parsing target coordinate max y")
	}

	return target{
		minX: minX,
		maxX: maxX,
		minY: minY,
		maxY: maxY,
	}
}

func exe2(lines []string) int {
	target := parseInput(lines)

	valid := 0
	for x := 1; x <= target.maxX; x++ {
		for y := target.minY; y <= util.Abs(target.minY); y++ {
			if _, ok := target.simulate(x, y); ok {
				valid++
			}
		}
	}

	return valid
}

func exe1(lines []string) int {
	target := parseInput(lines)

	maxY := -1
	for x := 1; x <= target.maxX; x++ {
		for y := target.minY; y <= util.Abs(target.minY); y++ {
			if probe, ok := target.simulate(x, y); ok {
				maxY = util.Max([]int{maxY, probe.maxY})
			}
		}
	}

	return maxY
}
