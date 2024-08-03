package utils

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"practice/data"
)

func ReadDataBase() ([]data.User, error) {
	var users []data.User
	dataPath := "data/storage.json"

	file, err := os.OpenFile(dataPath, os.O_RDWR, 0644)
	if err != nil {
		return users, err
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		return users, err
	}

	err = json.Unmarshal(fileData, &users)
	if err != nil {
		return users, err
	}

	if len(users) == 0 {
		return users, errors.New("user not found")
	}
	return users, nil
}
