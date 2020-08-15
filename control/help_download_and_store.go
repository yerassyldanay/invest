package control

import (
	"invest/utils"
	"mime/multipart"
	"net/http"
)

func (ds *DocStore) Download_and_store_file(r *http.Request) (map[string]interface{}, error) {
	/*
		Parsing a file to get meta-data on a file (name, format, etc.), file in bytes, generated string (filename), err
	*/

	var err error
	var handler *multipart.FileHeader

	handler, err = ds.Parse_file(r)
	if err != nil {
		return utils.ErrorInternalServerError, err
	}

	ds.Format, err = Parse_the_format_of_the_file(handler)
	if err != nil {
		return utils.ErrorInternalServerError, err
	}

	/*
		Storing a document on the directory
	*/
	ds.Directory = "/documents/docs/"
	err = ds.Store_document()

	if err != nil {
		return utils.ErrorInternalServerError, err
	}

	return utils.NoErrorFineEverthingOk, nil
}

