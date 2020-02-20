package reporter

import (
	"encoding/csv"
	"os"
	"path/filepath"
)

func csvExport(header []string,data [][]string,path string,name string) error {
	file, err := os.Create(filepath.Join(path,name))
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(header); err != nil {
		return err
	}
	for _, value := range data {
		if err := writer.Write(value); err != nil {
			return err // let's return errors if necessary, rather than having a one-size-fits-all error handler
		}
	}
	return nil
}
