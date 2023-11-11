package jsonmanagement

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var filePath = "storage/saves/"

func CreateJson(nameFile string) error {
	emptyData := make(map[string]interface{})
	emptyJSON, err := json.MarshalIndent(emptyData, "", "    ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filePath+nameFile+".json", emptyJSON, os.ModePerm); err != nil {
		return err
	}

	return nil
}
