package pather

import (
	"path/filepath"
)

/*
	GetFilePath:
		provide as /path/to/somewhere
 */
func GetFilePath(directory string) (string, error) {
	return filepath.Abs("." + directory)
}