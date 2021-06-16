package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var Document_download = func(w http.ResponseWriter, r *http.Request) {
	fname := "Document_download"

	fileName := mux.Vars(r)["file"]

	var filePath, err = filepath.Abs("./documents/docs/" + fileName)
	//fileName = "/documents/docs/" + fileName

	if err != nil {
		OnlyReturnInvalidParametersError(w, r, err.Error(), fname, "abs")
		return
	}

	fmt.Println(fname, filePath)

	// check whether this file exists
	if file, err := os.Stat(filePath); err != nil {
		OnlyReturnInvalidParametersError(w, r, err.Error(), fname, "stat")
		return
	} else {
		fmt.Println("file: ", file)
	}

	// read file
	reading, err := os.Open(filePath)
	if err != nil {
		OnlyReturnInvalidParametersError(w, r, err.Error(), fname, "open")
		return
	}
	defer reading.Close()

	// set status
	w.Header().Set("Content-Disposition", "attachment; filename=" + fileName)
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	w.WriteHeader(http.StatusOK)

	// send file
	_, err = io.Copy(w, reading)
	if err != nil {
		OnlyReturnInvalidParametersError(w, r, err.Error(), fname, "copy")
		return
	}
}
