package match

import "math"

func calcDist(x, y, x2, y2 int) int {
	return int(math.Max(float64(x2-x), float64(y2-y)))
}
