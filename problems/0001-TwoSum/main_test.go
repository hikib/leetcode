package twosum

import "testing"

func TestProblem(t *testing.T) {

	exampleTest := []struct {
		name   string
		input  []int
		target int
		output [2]int
	}{
		{"I", []int{2, 7, 11, 15}, 9, [2]int{0, 1}},
		{"II", []int{3, 2, 4}, 6, [2]int{1, 2}},
		{"III", []int{3, 3}, 6, [2]int{0, 1}},
	}

	for _, tt := range exampleTest {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.input, tt.target)
			if got != tt.output {
				t.Errorf("%s: got %d want %d", tt.name, got, tt.output)
			}
		})
	}
}
