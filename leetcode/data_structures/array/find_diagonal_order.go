// https://leetcode.com/explore/learn/card/array-and-string/202/introduction-to-2d-array/1167/

package array

func findDiagonalOrder(mat [][]int) []int {
	result := []int{}

	yLength := len(mat)
	if yLength == 0 {
		return result
	}

	xLength := len(mat[0])
	elementsCount := yLength * xLength
	result = make([]int, elementsCount)
	lastYIndex, lastXIndex := yLength-1, xLength-1

	x, y := 0, 0

	for i := 0; i < elementsCount; i++ {
		result[i] = mat[y][x]

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
