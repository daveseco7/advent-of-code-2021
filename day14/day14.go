package day14

import (
	"math"
	"strings"
)

const filePath = "input1.txt"

type pairs map[string]int

func (p *pairs) merge(other pairs) {
	for k, v := range other {
		(*p)[k] += v
	}
}

type letters map[string]int

type insertions map[string]string

type state struct {
	pairs      pairs
	letters    letters
	insertions insertions
}

func (s *state) step() {
	newPairs := make(pairs)

	for pair, insertion := range s.insertions {
		// match found, insertion of 2 new pairs needed

		if counter, ok := s.pairs[pair]; ok {
			// delete all matches to merge with the rest of the pairs
			delete(s.pairs, pair)

			// increment counter for new char
			s.letters[insertion] += counter

			// increment counter for new pairs
			newPairs[pair[:1]+insertion] += counter
			newPairs[insertion+pair[1:]] += counter
		}
	}

	s.pairs.merge(newPairs)
}

func parseInput(lines []string) state {
	ins := make(insertions)
	for i := 2; i < len(lines); i++ {
		s := strings.Split(lines[i], " -> ")
		ins[s[0]] = s[1]
	}

	p := make(pairs)
	for i := 0; i < len(lines[0])-1; i++ {
		left := lines[0][i : i+1]
		right := lines[0][i+1 : i+2]
		p[left+right]++
	}

	l := make(letters)
	for i := 0; i < len(lines[0]); i++ {
		l[lines[0][i:i+1]]++
	}

	return state{
		pairs:      p,
		letters:    l,
		insertions: ins,
	}
}

func exe2(lines []string) int {
	state := parseInput(lines)

	for i := 0; i < 40; i++ {
		state.step()
	}

	max, min := 0, math.MaxInt
	for _, v := range state.letters {
		if max < v {
			max = v
		}

		if min > v {
			min = v
		}
	}

	return max - min
}

func exe1(lines []string) int {
	state := parseInput(lines)

	for i := 0; i < 10; i++ {
		state.step()
	}

	max, min := 0, math.MaxInt
	for _, v := range state.letters {
		if max < v {
			max = v
		}

		if min > v {
			min = v
		}
	}

	return max - min
}
