package day10

import "sort"

const filePath = "input1.txt"

var (
	matching = map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}

	scoreCorrupted = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	scoreIncomplete = map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
)

type stack []string

func (s *stack) matches(str string) bool {
	if len(*s) != 0 {
		return str == matching[(*s)[len(*s)-1]]
	}
	return false
}

func isClose(s string) bool {
	_, ok := scoreCorrupted[s]
	return ok
}

func isCorrupted(line string) (points int, ok bool) {
	stk := make(stack, 0, len(line))
	for _, r := range line {
		s := string(r)
		if !isClose(s) {
			stk = append(stk, s)
			continue
		}
		if stk.matches(s) {
			stk = stk[:len(stk)-1]
			continue
		}
		return scoreCorrupted[s], true
	}

	return 0, false
}

func isIncomplete(line string) (points int, ok bool) {
	stk := make(stack, 0, len(line))
	for _, r := range line {
		s := string(r)
		if !isClose(s) {
			stk = append(stk, s)
			continue
		}
		if stk.matches(s) {
			stk = stk[:len(stk)-1]
			continue
		}
	}

	for i := len(stk) - 1; i >= 0; i-- {
		matchingClose := matching[stk[i]]
		points = points*5 + scoreIncomplete[matchingClose]
	}

	return points, points != 0
}

func exe2(lines []string) int {
	scores := make([]int, 0)
	for _, line := range lines {
		if _, ok := isCorrupted(line); ok {
			continue
		}

		if s, ok := isIncomplete(line); ok {
			scores = append(scores, s)
		}
	}

	sort.SliceStable(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})

	return scores[len(scores)/2]
}

func exe1(lines []string) (points int) {
	for _, line := range lines {
		s, ok := isCorrupted(line)
		if ok {
			points += s
		}
	}

	return points
}
