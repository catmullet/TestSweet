package TestSweet

import (
	"os"
	"path/filepath"
)

func getDataFile(name string) *os.File {
	filePath := filepath.Join("test_data", name)

	for i := 0; i < 5; i++ {
		jsonFile, err := os.Open(filePath)

		if err != nil {
			filePath = filepath.Join("..", filePath)
			continue
		}

		return jsonFile
	}

	return nil
}