package database

import (
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type database struct{}

func NewDatabase() Database {
	return &database{}
}

func (db *database) checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *database) ReadFromDB(path string) ([]interface{}, error) {
	err := db.checkFile(path)
    if err != nil {
        return nil, err
    }

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(content) == 0 {
		return nil, nil
	}

	var data []interface{}

	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (db *database) SaveInsertToDB(path string, data interface{}) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}
	
	return nil
}

func (db *database) SaveToDB(path string, data []interface{}) error {
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
