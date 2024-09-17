package search

func BinarySearch(list []int, found int) int {
	if len(list) == 0 {
		return -1
	}

	var left int
	right := len(list) - 1

	for (right - left) > 0 {
		sum := left + right
		mediumIndex := (sum - (sum % 2)) / 2

		switch {
		case list[mediumIndex] == found:
			return mediumIndex
		case list[mediumIndex] > found:
			right = mediumIndex - 1
		default:
			left = mediumIndex + 1
		}
	}
	if list[left] == found {
		return left
	}
	return -1
}
