package control

import (
	"invest/utils"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func (ds *DocStore) Parse_file(r *http.Request) (*multipart.FileHeader, error) {
	/*
		Creating a buffer of size 10 * 2^20
	*/
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return &multipart.FileHeader{}, err
	}

	/*
		parse file using the field name "uploadFile"
	*/
	var file, handler, err = r.FormFile("uploadFile")
	if err != nil {
		return &multipart.FileHeader{}, err
	}
	defer file.Close()

	var file_name = utils.Generate_Random_String(50)

	/*
		converting file into bytes
	*/
	fileInBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return &multipart.FileHeader{}, err
	}

	ds.Filename = file_name
	ds.ContentBytes = fileInBytes

	return handler, nil
}
