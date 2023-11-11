package jsonmanagement

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func Store(data map[string]interface{}, nameFile string) error {
	updatedJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filePath+nameFile+".json", updatedJSON, os.ModePerm); err != nil {
		return err
	}

	return nil
}
