package main

func main() {
	maxArea([]int{2, 3, 10, 5, 7, 8, 9})
}

func maxArea(height []int) int {
	maxContainer := 0
	for i, iH := range height {
		for j, jH := range height {
			if i == j {
				continue
			}
			if maxContainer < GetLowest(iH, jH)*AbsMinus(i, j) {
				maxContainer = GetLowest(iH, jH) * AbsMinus(i, j)
			}
		}
	}

	return maxContainer
}

func GetLowest(val int, otherVal int) int {
	if val <= otherVal {
		return val
	} else {
		return otherVal
	}
}

func AbsMinus(val int, additional int) int {
	if val >= additional {
		return val - additional
	} else {
		return additional - val
	}
}
