package service

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type FileWriter struct{}

func NewFileWriter() *FileWriter {
	return &FileWriter{}
}

func (fw *FileWriter) WriteDataToFile(fName string, path string, data string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	f, err := os.OpenFile(filepath.Join(path, fName), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	_, err = f.WriteString(data)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Created file " + fName + " successfully")
	return nil
}
