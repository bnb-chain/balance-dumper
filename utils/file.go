package utils

import (
	"encoding/csv"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func RemoveAllExcept(dir string, exceptName string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		if name != exceptName {
			err = os.RemoveAll(filepath.Join(dir, name))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func RemoveContents(dir string) error {
	return RemoveAllExcept(dir, "")
}

func ReadCSV(fileName string) (records [][]string, err error) {

	bts, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	reader := csv.NewReader(strings.NewReader(string(bts)))
	records, err = reader.ReadAll()

	return
}
