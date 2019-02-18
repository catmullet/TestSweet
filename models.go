package TestSweet

import (
	"regexp"
	"testing"
)

var Suite *TestSweet

func init() {
	Suite = new(TestSweet)
}

type TestSweet struct {
	SuiteOfTests []func(t *testing.T)
}

func (t *TestSweet) AddTest(test func(t *testing.T)) {
	t.SuiteOfTests = append(t.SuiteOfTests, test)
}

func (t *TestSweet) RunAll() {
	for _, val := range t.SuiteOfTests {
		testing.Main()

	}
}

type StringData struct {
	Types map[string]struct {
		Good []string `json:"good"`
		Bad  []string `json:"bad"`
	} `json:"types"`
}

type IntData struct {
	Types map[string]struct {
		Good []int `json:"good"`
		Bad  []int `json:"bad"`
	} `json:"types"`
}

type FloatData struct {
	Types map[string]struct {
		Good []float64 `json:"good"`
		Bad  []float64 `json:"bad"`
	} `json:"types"`
}
