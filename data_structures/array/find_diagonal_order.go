package array

func findDiagonalOrder(mat [][]int) []int {
	var result []int

	lastYIndex := len(mat) - 1
	if lastYIndex < 0 {
		return result
	}

	lastXIndex := len(mat[0]) - 1

	if lastXIndex < 0 {
		return result
	}

	var x int
	var y int

	elementsCount := (lastXIndex + 1) * (lastYIndex + 1)

	for i := 0; i < elementsCount; i++ {
		value := mat[y][x]
		result = append(result, value)

		if (x+y)%2 == 0 {
			if y == 0 && x != lastXIndex {
				x++
			} else if x == lastXIndex && y != lastYIndex {
				y++
			} else if y != 0 && x != lastXIndex {
				y--
				x++
			}
		} else {
			if x == 0 && y != lastYIndex {
				y++
			} else if y == lastYIndex && x != lastXIndex {
				x++
			} else if x != 0 && y != lastYIndex {
				y++
				x--
			}
		}
	}

	return result
}
