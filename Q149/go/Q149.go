package Q149

import "math"

type Formula struct {
	slope_top     int
	slope_bot     int
	intercept_top int
	intercept_bot int
}

func gcd(a, b int) int {
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}

	if b != 0 {
		return gcd(b, a%b)
	} else {
		return a
	}
}

func getFormula(point1 []int, point2 []int) Formula {
	y := point2[1] - point1[1]
	x := point2[0] - point1[0]
	var d int
	if x != 0 && y != 0 {
		g := gcd(x, y)
		y /= g
		x /= g
		d := point1[1]*x - point1[0]*y

		if x < 0 {
			x *= -1
			y *= -1
			d *= -1
		}
	} else if x == 0 {
		y = math.MaxInt
		d = point1[0]
	} else {
		x = math.MaxInt
		d = point1[1]
	}

	return Formula{
		slope_top:     y,
		slope_bot:     x,
		intercept_top: d,
		intercept_bot: x,
	}
}

func maxPoints(points [][]int) int {
	points_len := len(points)

	graph := make([][]Formula, points_len)
	for i := 0; i < points_len; i++ {
		graph[i] = make([]Formula, points_len)
	}

	formulaMap := map[Formula]int{}
	for i := 0; i < points_len; i++ {
		for j := i + 1; j < points_len; j++ {
			formula := getFormula(points[i], points[j])
			formulaMap[formula]++
		}
	}

	max := 0
	for _, value := range formulaMap {
		if max < value {
			max = value
		}
	}

	if max == 0 {
		return 1
	} else {
		max *= 2
		i := 1
		for i = 1; i < 300; i++ {
			if max == i*(i+1) {
				break
			}
		}

		return i + 1
	}
}

// C(4,2)
// 4!
// 2!2!

// C(x,2)
// x!
// (x-2)!2!
