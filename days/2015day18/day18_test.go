package day18_test

import (
	_ "embed"
	"testing"

	day18 "github.com/mleone10/advent-of-code-2023/days/2015day18"
	"github.com/mleone10/advent-of-code-2023/internal/assert"
)

//go:embed test_input.txt
var testInput string

//go:embed input.txt
var input string

var tcs = []struct {
	input           string
	steps           int
	stepsPartTwo    int
	initialOn       int
	onAfterOneStep  int
	expectedPartOne int
	expectedPartTwo int
}{
	{
		testInput,
		4,
		5,
		15,
		11,
		4,
		17,
	},
	{
		input,
		100,
		100,
		4934,
		2871,
		768,
		781,
	},
}

func TestInitialState(t *testing.T) {
	for _, tc := range tcs {
		lg := day18.NewLightGrid(tc.input)
		assert.Equals(t, tc.initialOn, lg.NumOn())
	}
}

func TestOneStep(t *testing.T) {
	for _, tc := range tcs {
		lg := day18.NewLightGrid(tc.input)
		lg.Step()
		assert.Equals(t, tc.onAfterOneStep, lg.NumOn())
	}
}

func TestSolvePartOne(t *testing.T) {
	for _, tc := range tcs {
		lg := day18.NewLightGrid(tc.input)
		lg.StepN(tc.steps)
		assert.Equals(t, tc.expectedPartOne, lg.NumOn())
	}
}

func TestSolvePartTwo(t *testing.T) {
	for _, tc := range tcs {
		lg := day18.NewLightGrid(tc.input)
		lg.CornersStuckOn()
		lg.StepN(tc.stepsPartTwo)
		assert.Equals(t, tc.expectedPartTwo, lg.NumOn())
	}
}
