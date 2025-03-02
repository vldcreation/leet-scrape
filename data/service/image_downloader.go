package service

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/vldcreation/leet-scrape/v2/internal/errors"
)

type ImageDownloader struct{}

func NewImageDownloader() *ImageDownloader {
	return &ImageDownloader{}
}

func (d *ImageDownloader) DownloadImageFromUrl(fName string, path string, url string) error {
	response, e := http.Get(url)
	if e != nil {
		return errors.FailedToDownloadImage
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(response.Body)

	//open a file for writing
	err := os.MkdirAll(path, 0644)
	if err != nil {
		return errors.FailedToCreateDirectory
	}
	file, err := os.Create(filepath.Join(path, fName))
	if err != nil {
		return errors.FailedToWriteImage
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return errors.FailedToWriteImage
	}

	return nil
}
