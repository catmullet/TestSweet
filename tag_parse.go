package TestSweet

import (
	"strings"
)

type DataFlag string

const BAD = "bad"
const GOOD = "good"

type Tag struct {
	Type     string
	DataFlag DataFlag
}

func ParseTestSweetStructTag(tagValue string) Tag {
	dataSlice := strings.Split(tagValue, "|")

	tag := Tag{}

	if len(dataSlice) > 0 {
		tag.Type = dataSlice[0]
		tag.DataFlag = GOOD
	}
	if len(dataSlice) > 1 {
		switch dataSlice[1] {
		case BAD:
			tag.DataFlag = BAD
			break
		default:
			tag.DataFlag = GOOD
		}
	}
	return tag
}
