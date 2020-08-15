package control

import (
	"errors"
	"mime/multipart"
	"regexp"
)

func Parse_the_format_of_the_file(handler *multipart.FileHeader) (string, error) {
	/*
		this part will obtain content-type: e.g. image/jpg
	*/
	compiled := regexp.MustCompile("/")
	var file_formats = compiled.Split(handler.Header.Get("Content-Type"), -1)
	var file_format string

	/*
		documents/pdf will be parsed as .pdf
		which will be used when naming an image file
		file_name.png
	*/
	if len(file_formats) != 2 {
		return "", errors.New("invalid file format")
	} else {
		file_format = "." + file_formats[1]
	}

	return file_format, nil
}
