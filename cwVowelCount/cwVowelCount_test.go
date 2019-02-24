package cwvowelcount

import "testing"

func Test_VowelCount(t *testing.T) {
	input := "father"
	actual := VowelCount(input)
	expected := 2
	if actual != expected {
		t.Errorf("input: %s, actual: %d, expected %d \n", input, actual, expected)
	}
}
