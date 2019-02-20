package TestSweet

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

var sd *stringData

func init() {
	jsonFile := getDataFile("strings.json")

	var data stringData

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &data)

	sd = &data
}

func GetStringValue(tag string) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	value := ""

	dataTag := ParseTestSweetStructTag(tag)

	tagType := sd.Types[dataTag.Type]
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
