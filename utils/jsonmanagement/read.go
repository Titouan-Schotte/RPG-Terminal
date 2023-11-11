package jsonmanagement

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadJsonCore(filePath string) (map[string]interface{}, error) {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if len(fileContent) == 0 {
		data = make(map[string]interface{})
	} else {
		if err := json.Unmarshal(fileContent, &data); err != nil {
			return nil, err
		}
	}

	return data, nil
}
func ReadJson(jsonIn string) map[string]interface{} {
	data, errRead := ReadJsonCore(filePath + jsonIn + ".json")
	if errRead != nil {
		fmt.Println("Erreur lors de la lecture du JSON :", errRead)
		return nil
	}
	return data
}
