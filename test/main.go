package main

import (
	"github.com/catmullet/TestSweet"
	"testing"
)

type TestStruct struct {
	Pickles string  `test:"name|good"`
	Nest    int     `test:"age"`
	Creepy  float64 `test:"price|good"`
	Child   ChildStruct
}

type ChildStruct struct {
	Hello string `test:"name|good"`
}

func TestTestSweet(t *testing.T) {
	ts := TestStruct{}
	ts.Child = ChildStruct{}

	TestSweet.FillStruct(&ts)
	TestSweet.FillStruct(&ts.Child)

	t.Error("Failed")
}

func main() {
	TestSweet.Suite.AddTest(TestTestSweet)
	TestSweet.Suite.RunAll()
}
