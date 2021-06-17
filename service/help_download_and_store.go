package service

import (
	"github.com/yerassyldanay/invest/utils/errormsg"
	"net/http"
)

func (ds *DocStore) Download_and_store_file(r *http.Request) (map[string]interface{}, error) {
	//	parsing a file to get meta-data on a file (name, format, etc.),
	//	file in bytes, generated string (filename), err

	_, err := ds.Parse_file(r)
	if err != nil {
		return errormsg.ErrorInternalServerError, err
	}

	/*
		get format of the file
			* "image/png" -> .png
	 */
	if err = ds.OnlyParseFormatOfTheFile(); err != nil {
		return errormsg.ErrorInternalServerError, err
	}
	//return utils.ErrorInternalServerError, errors.New("new one")

	/*
		Storing a document on the directory
	*/
	ds.Directory = "/documents/docs/"
	err = ds.Store_document()

	if err != nil {
		return errormsg.ErrorInternalServerError, err
	}

	return errormsg.NoErrorFineEverthingOk, nil
}

