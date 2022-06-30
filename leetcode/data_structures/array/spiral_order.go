// https://leetcode.com/explore/learn/card/array-and-string/202/introduction-to-2d-array/1168/

package array

const (
	Right = "right"
	Down  = "down"
	Left  = "left"
	Up    = "up"
)

func spiralOrder(matrix [][]int) []int {
	result := []int{}

	lengthY := len(matrix)
	if lengthY == 0 {
		return result
	}

	lengthX := len(matrix[0])

	elementsCount := lengthX * lengthY
	result = make([]int, 0, elementsCount)
	x, y := 0, 0
	leftBorder, topBorder, rightBorder, bottomBorder := 0, 0, lengthX-1, lengthY-1
	direction := Right

	for i := 0; i < elementsCount; i++ {
		result = append(result, matrix[y][x])

		switch direction {
		case Right:
			if x == rightBorder {
				topBorder++
				direction = Down
				y++
			} else {
				x++
			}
		case Down:
			if y == bottomBorder {
				rightBorder--
				direction = Left
				x--
			} else {
				y++
			}
		case Left:
			if x == leftBorder {
				bottomBorder--
				direction = Up
				y--
			} else {
				x--
			}
		case Up:
			if y == topBorder {
				leftBorder++
				direction = Right
				x++
			} else {
				y--
			}
		}
	}

	return result
}
