package easy

func HammingDistance(x int, y int) int {
	var diff = 0
	for x > 0 || y > 0 {
		xb, yb := 0, 0
		if x > 0 {
			xb = 1 & x
		}

		if y > 0 {
			yb = 1 & y
		}

		if yb != xb {
			diff++
		}

		x = x >> 1
		y = y >> 1
	}
	return diff
}
