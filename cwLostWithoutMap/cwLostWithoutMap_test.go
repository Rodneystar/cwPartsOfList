package cwlostwithoutmap

import "testing"

func Test_Maps(t *testing.T) {
	expected := []int{2, 4, 6}
	actual := Maps([]int{1, 2, 3})

	if !sliceEqual(expected, actual) {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}

func Test_MapsSquared(t *testing.T) {
	expected := []int{1, 4, 9}
	actual := MapsSq([]int{1, 2, 3})

	if !sliceEqual(expected, actual) {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, obj := range a {
		if b[i] != obj {
			return false
		}
	}
	return true
}
func Test_map(t *testing.T) {
	expected := []int{2, 4, 6}
	actual := Map([]int{1, 2, 3},
		func(n int) int { return n * 2 })

	if !sliceEqual(expected, actual) {
		t.Errorf("expected: %d, actual %d\n", expected, actual)
	}
}
