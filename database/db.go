package database

import (
	"encoding/json"
	"io/ioutil"
)

type database struct{}

func NewDatabase() Database {
	return &database{}
}

func (db *database) ReadFromDB(path string) (interface{}, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data interface{}

	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (db *database) SaveToDB(path string, data interface{}) error {
	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, content, 0644)
	if err != nil {
		return err
	}

	return nil
}
