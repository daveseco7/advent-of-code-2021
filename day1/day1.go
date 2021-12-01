package day1

import (
	"fmt"

	"github.com/daveseco7/advent-of-code-2021/util"
)

const filePath = "/Users/dave/go/src/github.com/daveseco7/advent-of-code-2021/day1/input1.txt"

func exe1(lines []int) (asc int) {
	for i := 0; i < len(lines)-1; i++ {
		if lines[i+1]-lines[i] > 0 {
			asc++
		}
	}

	return asc
}

func exe2(lines []int) (asc int) {
	newInput := make([]int, 0)
	for i := 0; i < len(lines)-3; i += 2 {
		threeMeasurementA := lines[i] + lines[i+1] + lines[i+2]
		threeMeasurementB := lines[i+1] + lines[i+2] + lines[i+3]
		newInput = append(newInput, threeMeasurementA)
		newInput = append(newInput, threeMeasurementB)
	}

	for i := 0; i < len(newInput)-1; i++ {
		if newInput[i+1]-newInput[i] > 0 {
			asc++
		}
	}

	return asc
}

func Exec() {
	lineArray, err := util.ReadLinesAsInt(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Println(exe1(lineArray))
	fmt.Println(exe2(lineArray))
}
