package main

import (
	"encoding/json"
	"io"
	"os"
)

type FileValues struct {
	ApiKey      string `json:"api_key"`
	MaxHandlers int    `json:"max_handlers"`
}

func GetFileValues(fileName string) (data FileValues, err error) {
	fileValues, err := os.Open(fileName)
	if err != nil {
		return data, err
	}
	defer func() { _ = fileValues.Close() }()

	bytesData, err := io.ReadAll(fileValues)
	if err != nil {
		return data, err
	}

	if err = json.Unmarshal(bytesData, &data); err != nil {
		return data, err
	}

	return data, nil
}
