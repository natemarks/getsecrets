package file

import (
	"os"
)

// WriteFile takes a filename and data and writes the data to the file
func WriteFile(filename, data string) (err error) {

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()
	f.WriteString(data)
	if err != nil {
		return err
	}

	f.Sync()
	return err
}
