package main

import (
	"bufio"
	"io"
	"os"
)

func CreateNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)
	return nil
}

func ReadFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		message += string(line) + "\n"
	}

	return message, nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)
	return nil
}

func main() {
	// CreateNewFile("sample.log", "this is sample log")
	// result, _ := ReadFile("sample.log")
	// fmt.Println(result)
	addToFile("sample.log", "\nnew sample message")
}
