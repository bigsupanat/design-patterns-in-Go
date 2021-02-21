package main

import "testing"

type DummyDatabase struct {
	dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3,
		}
	}
	return d.dummyData[name]
}

func TestGetTotalPopulationEx(t *testing.T) {
	names := []string{"alpha", "beta", "gamma"}
	actual := GetTotalPopulationEx(&DummyDatabase{}, names)
	expected := 6
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
