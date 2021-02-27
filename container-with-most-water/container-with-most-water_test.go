package container_with_most_water

import (
	"testing"
)

func TestTheContainerWithMostWater(t *testing.T) {
	expects := make([]int, 0)
	testCases := make([][]int, 0)

	expects = append(expects, []int{
		1,
		16,
		2,
		36,
		70,
	}...)
	testCases = append(testCases, [][]int{
		{1, 1},
		{4, 3, 2, 1, 4},
		{1, 2, 1},
		{2, 3, 10, 5, 7, 8, 9},
		{1, 3, 10, 5, 7, 8, 9, 19 ,4, 20},
	}...)

	for i, testCase := range testCases {
		if expects[i] != MaxArea(testCase) {
			t.Errorf(`Error!! result should equal to %d, but it has %d`, expects[i], MaxArea(testCase))
		}
	}
}
