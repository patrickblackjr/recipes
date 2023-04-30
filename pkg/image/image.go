package image

import (
	"errors"
	"io"
	"net/http"
	"os"
)

func DownloadFile(URL string, fileName string) error {
	res, err := http.Get(URL)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New("failed to fetch image. Response code is not 200")
	}

	// Create an empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the downloaded bytes to the file
	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}

	return nil
}
