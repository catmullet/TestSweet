package main

import (
	"fmt"
	"github.com/catmullet/TestSweet/TestSweet"
)

type TestStruct struct {
	Pickles string  `TestSweet:"name|good"`
	Nest    int     `TestSweet:"age"`
	Creepy  float64 `TestSweet:"price|good"`
	Child   ChildStruct
}

type ChildStruct struct {
	Hello string `TestSweet:"name|good"`
}

func main() {
	ts := TestStruct{}
	ts.Child = ChildStruct{}

	TestSweet.FillStruct(&ts)
	TestSweet.FillStruct(&ts.Child)

	fmt.Println(ts)
}

func TestThis() {

}
