package levelmanagement

import (
	"encoding/json"
	"io/ioutil"
)

type Level struct {
	Name  string                `json:"name"`
	Level [][]map[string]string `json:"level"`
}

func ReadLevel(levelName string) (string, [][]map[string]string, error) {
	filePath := "storage/levels/" + levelName + ".json"

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", nil, err
	}

	var level Level
	err = json.Unmarshal(data, &level)
	if err != nil {
		return "", nil, err
	}

	return level.Name, level.Level, nil
}
