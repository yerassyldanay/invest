package service

import (
	"errors"
	"fmt"
	"github.com/h2non/filetype"
	"strings"
)

func (ds *DocStore) OnlyParseFormatOfTheFile() (error) {
	//detectedType := http.DetectContentType(ds.ContentBytes)

	kind, err := filetype.Match(ds.ContentBytes)
	fmt.Println(kind)
	switch {
	case err != nil:
		return err
	case kind.Extension == "unknown" || kind.MIME.Value == "":
		splits := strings.Split(ds.RawFileName, ".")
		if len(splits) < 1 {
			return errors.New("invalid file extension")
		}

		ds.Format = "." + splits[len(splits) - 1]
		return nil
	}

	ds.Format = "." + kind.Extension
	return nil
}
