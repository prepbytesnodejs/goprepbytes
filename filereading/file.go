package filereading

import (
	"fmt"
	"os"
)

func ReadFile(path string) (string, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		// That  error is there in the system

		return "", err
	}

	return string(data), err

}

func WriteFile(data string, name string) error {

	fmt.Println()

	data2 := []byte(data)

	return os.WriteFile(name, data2, 0666)

}
