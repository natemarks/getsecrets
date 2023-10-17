package file

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// WriteFile takes a filename and data and writes the data to the file
func WriteFile(filename, data string) {

	f, err := os.Create(filename)
	check(err)

	defer f.Close()
	f.WriteString(data)
	check(err)

	f.Sync()
	check(err)
}
