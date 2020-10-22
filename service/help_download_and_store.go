package service

import (
	"invest/utils"
	"net/http"
)

func (ds *DocStore) Download_and_store_file(r *http.Request) (map[string]interface{}, error) {
	//	parsing a file to get meta-data on a file (name, format, etc.),
	//	file in bytes, generated string (filename), err

	_, err := ds.Parse_file(r)
	if err != nil {
		return utils.ErrorInternalServerError, err
	}

	/*
		get format of the file
			* "image/png" -> .png
	 */
	if err = ds.OnlyParseFormatOfTheFile(); err != nil {
		return utils.ErrorInternalServerError, err
	}
	//return utils.ErrorInternalServerError, errors.New("new one")

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

