package jsonmanagement

import "os"

func DeleteJson(nameFile string) error {
	return os.Remove(filePath + nameFile)
}
