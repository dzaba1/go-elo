package main

import (
	"log"
	"os"
)

func FileExists(file string) bool {
	_, error := os.Stat(file)

	if os.IsNotExist(error) {
		return false
	} else {
		return true
	}
}

func DeleteFileSafe(file string) {
	if FileExists(file) {
		err := os.Remove(file)
		if err != nil {
			log.Printf("Error removing file %s", file)
		}
	}
}

func WriteAllFile(file string, content string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func ReadAllFile(file string) (string, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return "", nil
	}

	return string(b), nil
}
