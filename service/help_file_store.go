package service

import (
	"os"
	"path/filepath"
)

/*
	store document on hard disk
 */
func (d *DocStore) Store_document() (error) {
	/*
		creating a file within the indicated directory
	*/
	var tempFileDir = d.Filename + d.Format
	var a = "." + filepath.Join(d.Directory, filepath.Base(tempFileDir))
	newDocument, err := os.OpenFile(a, os.O_EXCL|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	/*
		writing bytes into the file (which has been created)
	*/
	_, err = newDocument.Write(d.ContentBytes)
	if err != nil {
		return err
	}

	return nil
}
