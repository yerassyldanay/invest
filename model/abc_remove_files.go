package model

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// files, after requesting an analysis, are left
// this is to remove those files periodically
func Remove_files_left_after_analysis_periodically(ctx context.Context) {

	removeFilesHelper := func() {
		directoryPath := "../documents/analysis/"

		err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {

			if path == directoryPath {
				return nil
			}

			if _err := os.Remove(directoryPath + info.Name()); _err != nil {
				fmt.Println("[analysis/file] could not remove file: ", _err)
			} else {
				fmt.Println("[analysis/file] removed a file: ", directoryPath + info.Name())
			}
			return nil
		})

		if err != nil {
			fmt.Println("could not go through files in a directory")
		}
	}
	
	for {
		select {
		case <- time.Tick(time.Second):
			removeFilesHelper()
		case <- ctx.Done():
			return
		}
	}
}
