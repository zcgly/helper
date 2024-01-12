package helper

import (
	"encoding/json"
	"io"
	"os"
)

func ReadJsonFile(filename string, v any) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	bs, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(bs, v)
}
