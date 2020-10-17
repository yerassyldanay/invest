package service

import (
	"errors"
	"mime"
	"net/http"
)

func (ds *DocStore) OnlyParseFormatOfTheFile() (error) {
	detectedType := http.DetectContentType(ds.ContentBytes)

	// get the extension
	formatsOfFile, err := mime.ExtensionsByType(detectedType)
	switch {
	case err != nil:
		return err
	case len(formatsOfFile) < 1:
		return errors.New("len of formats list is 0")
	}

	ds.Format = formatsOfFile[0]
	return nil
}
