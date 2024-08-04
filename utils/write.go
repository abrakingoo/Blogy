package utils

import (
	"encoding/json"
	"os"
	"practice/data"
)

func WriteToDataBase(users []data.User, newUser data.User) (bool, error) {
	filepath := "data/storage.json"

	users = append(users, newUser)

	marshaledUsers, err := json.MarshalIndent(users, "", " ")

	if err != nil {
		return false, err
	}

	err = os.WriteFile(filepath, marshaledUsers, 0644)

	if err != nil {
		return false, err
	}

	return true, nil
}
