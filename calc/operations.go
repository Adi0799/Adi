package calc

func Calc(x, y int) (sum int, diff int, div float64) {
	sum = x + y
	diff = x - y
	div = float64(x) / float64(y)
	return
}
