package file

import (
	"os"
)

func ReadFile(inputfile string) string {
	data, err := os.ReadFile(inputfile)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func WriteFile(outputfile string, data string) {
	err := os.WriteFile(outputfile, []byte(data), 0644)
	if err != nil {
		panic(err)
	}
}
