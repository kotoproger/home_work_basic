package search

func BinarySearch(list []int, found int) int {
	var indexShift int
	for len(list) > 0 {
		mediumIndex := (len(list) - len(list)%2) / 2

		switch {
		case list[mediumIndex] == found:
			return indexShift + mediumIndex
		case list[mediumIndex] > found:
			list = list[0:(mediumIndex)]
		default:
			list = list[(mediumIndex + 1):]
			indexShift += mediumIndex + 1
		}
	}
	return -1
}
