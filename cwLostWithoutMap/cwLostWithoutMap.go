package cwlostwithoutmap

func double(y int) int {
	return y * 2
}
func Maps(x []int) []int {
	return Map(x, double)
}

func MapsSq(x []int) []int {
	return Map(x, func(n int) int { return n * n })
}

func Map(arr []int, fn func(n int) int) []int {
	var newArr = make([]int, len(arr))
	for i, n := range arr {
		newArr[i] = fn(n)
	}
	return newArr
}
