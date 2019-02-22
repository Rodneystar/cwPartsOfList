package cwpartsoflist

import (
	"testing"
)

func Test_getPairOfStrings_stringArrayandIndex0_stringFirstwordSeparate(t *testing.T) {
	input := []string{"never", "forget", "your", "towel"}
	actual := getPairOfStrings(input, 0)
	expected := "(never, forget your towel)"

	if actual != expected {
		t.Errorf("input: %s, actual: %s, expected %s \n", input, actual, expected)
	}
}

func Test_getPairOfStrings_stringArrayandIndex1_stringFirstAndSecondWordSeparate(t *testing.T) {
	input := []string{"never", "forget", "your", "towel"}
	actual := getPairOfStrings(input, 1)
	expected := "(never forget, your towel)"

	if actual != expected {
		t.Errorf("input: %s, actual: %s, expected %s \n", input, actual, expected)
	}
}

func Test_PartList(t *testing.T) {
	input := []string{"I", "wish", "I", "hadn't", "come"}
	actual := PartList(input)
	expected := "(I, wish I hadn't come)(I wish, I hadn't come)(I wish I, hadn't come)(I wish I hadn't, come)"
	if actual != expected {
		t.Errorf("input: %s, actual: %s, expected %s \n", input, actual, expected)
	}

}

func Test_PartListAgain(t *testing.T) {
	input := []string{"cdIw", "tzIy", "xDu", "rThG"}
	actual := PartList(input)
	expected := "(cdIw, tzIy xDu rThG)(cdIw tzIy, xDu rThG)(cdIw tzIy xDu, rThG)"
	if actual != expected {
		t.Errorf("input: %s, actual: %s, expected %s \n", input, actual, expected)
	}
}
