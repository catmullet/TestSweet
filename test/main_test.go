package main

import (
	"fmt"
	. "github.com/catmullet/TestSweet"
	"github.com/labstack/gommon/log"
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

func HelloWorld() {
	fmt.Println("Hello World")
}

func Goodbye() {
	fmt.Println("Goodbye")
}

func init() {
	TestSweet.AddBootstapFunction(HelloWorld)
	TestSweet.AddTeardownFunction(Goodbye)
}

func TestTestSweet(t *testing.T) {
	teardown := TestSweet.Bootstrap(t)
	defer teardown(t)

	ts := TestStruct{}
	ts.Child = ChildStruct{}

	TestSweet.FillStruct(&ts)
	TestSweet.FillStruct(&ts.Child)

	log.Info(ts)
}
