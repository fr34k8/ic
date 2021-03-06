package main

import (
	"io/ioutil"
	"os"
)

func Create(filename string) error {
	f, err := os.Create(filename)
	defer f.Close()

	return err
}

func Write(filename string, content []byte) {
	err := ioutil.WriteFile(filename, content, 0600)
	CheckError(err)
}

func Append(filename string, content []byte) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	CheckError(err)

	defer f.Close()

	_, err = f.Write(content)

	return err
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	} else {
		return false
	}
}

func Read(filename string) string {
	s, err := ioutil.ReadFile(filename)
	CheckError(err)

	return string(s)
}
