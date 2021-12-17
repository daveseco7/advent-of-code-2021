package day15

import (
	"testing"

	"github.com/daveseco7/advent-of-code-2021/util"
)

func TestExe1(t *testing.T) {
	const expected = 755

	lineArray, err := util.ReadLines(filePath)
	if err != nil {
		t.Fatal(err)
	}

	ans := exe1(lineArray)
	if ans != expected {
		t.Errorf("invalid response %d", ans)
	}
}

func TestExe2(t *testing.T) {
	const expected = 3016

	lineArray, err := util.ReadLines(filePath)
	if err != nil {
		t.Fatal(err)
	}

	ans := exe2(lineArray)
	if ans != expected {
		t.Errorf("invalid response %d", ans)
	}
}
