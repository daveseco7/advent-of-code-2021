package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const filePath = "input1.txt"

type direction string

const (
	forward direction = "forward"
	down    direction = "down"
	up      direction = "up"
)

type position struct {
	horizontal int
	depth      int
	aim        int
}

func (p *position) product() int {
	return p.horizontal * p.depth
}

func (p *position) move(dir direction, value int) error {
	switch dir {
	case forward:
		p.horizontal += value
		p.depth += p.aim * value
	case down:
		p.aim += value
	case up:
		p.aim -= value
	default:
		return fmt.Errorf("invalid move direction: %s", dir)
	}
	return nil
}

func parseMoveOperation(line string) (direction, int, error) {
	slice := strings.Split(line, " ")
	if len(slice) != 2 {
		return "", 0, fmt.Errorf(" invalid move direction: %s", line)
	}
	n, err := strconv.Atoi(slice[1])
	if err != nil {
		return "", 0, fmt.Errorf(" invalid move intensity: %s", slice[1])
	}

	return direction(slice[0]), n, nil
}

func exe(lines []string) int {
	p := position{0, 0, 0}

	for _, l := range lines {
		dir, intensity, err := parseMoveOperation(l)
		if err != nil {
			log.Fatalf("error in exer1: %v", err)
		}
		if err := p.move(dir, intensity); err != nil {
			log.Fatalf("error in exer1: %v", err)
		}
	}

	return p.product()
}
