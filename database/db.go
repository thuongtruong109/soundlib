package database

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

func ReadMapFromDB[T any](path string) (map[string]*T, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data map[string]*T

	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func SaveMapToDB[T any](path string, data T) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func ReadSliceFromDB[T any](path string) ([]*T, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data []*T
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func SaveSliceToDB[T any](path string, data []*T) error {
	content, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	err = ioutil.WriteFile(path, content, 0644)
	if err != nil {
		return nil
	}
	return nil
}