package cwnbateams

import (
	"fmt"
	"testing"
)

func Test_nba_cup(t *testing.T) {
	r := `Los Angeles Clippers 104 Dallas Mavericks 88`
	expected := "Los Angeles Clippers:W=1;D=0;L=1;Scored=204;Conceded=208;Points=3"
	actual := nba_cup(r, "Los Angeles Clippers")

	if expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}
func Test_parseResults(t *testing.T) {
	var rO = `Los Angeles Clippers 104 Dallas Mavericks 88`

	expectedTeams, expectedScores :=
		[]string{"Los Angeles Clippers", "Dallas Mavericks"},
		[]int{104, 88}
	actualTeams, actualScores := parseResults(rO)

	if !intSliceEqual(expectedScores, actualScores) || !stringSliceEqual(expectedTeams, actualTeams) {
		t.Errorf("expectedScores: %d, actual: %d, expectedTeams: %s, actual: %s", expectedScores, actualScores, expectedTeams, actualTeams)
	}
}

func Test_parseGame_returnsCorrectMap(t *testing.T) {
	var str = "what whatters 100 why whyers 101"
	actual := parseGame(str)
	expected := map[string]int{
		"what whatters": 100,
		"why whyers":    101,
	}
	if !mapEquals(actual, expected) {
		fmt.Printf("actual: %v, expected: %v\n", actual, expected)
	}

}

func Test_getGameScores(t *testing.T) {
	var str = "what whatters 100 why whyers 101"
	actual := getGameScores(str)
	expected := [2]string{
		"what whatters 100",
		"why whyers 101",
	}
	if !stringSliceEqual(actual[:], expected[:]) {
		t.Errorf("actual: %v, expected: %v\n", actual, expected)
	}
}

func Test_mapEquals(t *testing.T) {
	first := map[string]int{
		"one": 1,
		"two": 2,
	}
	second := map[string]int{
		"one": 1,
		"two": 2,
	}

	if !mapEquals(first, second) {
		t.Errorf("first: %v, second: %v", first, second)
	}

	second["two"] = 3

	if mapEquals(first, second) {
		t.Errorf("first: %v, second: %v", first, second)
	}
}

func mapEquals(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for key, value := range a {
		if b[key] != value {
			return false
		}
	}
	return true
}

func intSliceEqual(a, b []int) bool {
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

func stringSliceEqual(a, b []string) bool {
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
