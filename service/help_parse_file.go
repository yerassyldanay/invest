package service

import (
	"invest/utils"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

/*
	Does:
		* ParseMultipartForm
		* download a file
		* generate random string (length is 50 characters)
		* read a file as bytes
 */
func (ds *DocStore) Parse_file(r *http.Request) (*multipart.FileHeader, error) {

	// creating a buffer of size 10 * 2^20 ~= 50 megabyte
	if err := r.ParseMultipartForm(50 << 20); err != nil {
		return &multipart.FileHeader{}, err
	}

	// parse file using the field name "uploadFile"
	var file, handler, err = r.FormFile("uploadFile")
	if err != nil {
		return &multipart.FileHeader{}, err
	}
	defer file.Close()

	// generate a name for the file
	var file_name = utils.Generate_Random_String(50)

	// converting file into bytes
	fileInBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return &multipart.FileHeader{}, err
	}

	ds.Filename = file_name
	ds.ContentBytes = fileInBytes

	return handler, nil
}
