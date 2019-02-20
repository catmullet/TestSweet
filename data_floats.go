package TestSweet

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

var fd *floatData

func init() {
	jsonFile := getDataFile("floats.json")

	var data floatData

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &data)

	fd = &data
}

func GetFloatValue(tag string) float64 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	value := 0.0

	dataTag := ParseTestSweetStructTag(tag)

	tagType := fd.Types[dataTag.Type]
	switch dataTag.DataFlag {
	case GOOD:
		if len(tagType.Good) > 0 {
			value = tagType.Good[r1.Intn(len(tagType.Good))]
		}
	case BAD:
		if len(tagType.Bad) > 0 {
			value = tagType.Bad[r1.Intn(len(tagType.Bad))]
		}
	}

	return value
}
