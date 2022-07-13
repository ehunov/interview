package tasks

func intersect(firstStart int, firstEnd int, secondStart int, secondEnd int) []int {
	intersection := make([]int, 2)

	if secondStart < firstStart {
		return intersect(secondStart, secondEnd, firstStart, firstEnd)
	}

	if firstEnd >= secondStart {
		intersection[0] = max(firstStart, min(firstEnd, secondStart))
		intersection[1] = min(firstEnd, secondEnd)
	}

	return intersection
}

func min(a int, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a int, b int) int {
	if a < b {
		return b
	}

	return a
}
