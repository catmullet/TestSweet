package TestSweet

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

var id *IntData

func init() {
	jsonFile, err := os.Open("data/ints.json")

	if err != nil {
		panic(err)
	}

	var data IntData

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &data)

	id = &data
}

func GetIntValue(tag string) int64 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	value := 0

	dataTag := ParseTestSweetStructTag(tag)

	tagType := id.Types[dataTag.Type]
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

	return int64(value)
}
