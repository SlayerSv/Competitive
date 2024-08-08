package algs

// returns index of the first element that is not less than x
func leftBinSearch(stops []int, x int) int {
	l, r := -1, len(stops)
	for r-l > 1 {
		mid := (r + l) / 2
		if x <= stops[mid] {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

// returns index of the first element that is greater than x
func rightBinSearch(stops []int, x int) int {
	l, r := -1, len(stops)
	for r-l > 1 {
		mid := (r + l) / 2
		if x < stops[mid] {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}
