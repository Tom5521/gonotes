package search

import (
	"fmt"

	"github.com/Tom5521/gonotes/internal/db"
	"github.com/Tom5521/gonotes/internal/options"
	"github.com/Tom5521/slicelib"
)

func Deep(arg string) (file db.File, err error) {
	var files slicelib.Slice[db.File]

	// Filter 1: File location.
	switch {
	case options.Temporal:
		files = db.TemporalFiles()
	case options.Normal:
		files = db.NormalFiles()
	default:
		files = db.Files
	}
	compare := func(file db.File) (db.File, error) {
		if options.Filetype != file.Type {
			err = fmt.Errorf(
				"types do not match, type requested: %s | type found: %s",
				options.Filetype,
				file.Type,
			)
		}
		if options.CustomName != "" && options.CustomName != file.Name {
			err = fmt.Errorf(
				"names do not match, type requested: %s | type found: %s",
				options.CustomName,
				file.Name,
			)
		}
		if options.CustomID != -1 && options.CustomID != int(file.ID) {
			err = fmt.Errorf(
				"IDs do not match, type requested: %v | type found: %v",
				options.CustomID,
				file.ID,
			)
		}

		if err != nil {
			return db.File{}, err
		}
		return file, nil
	}

	// Try to get file by the name.
	fileByName, errByName := byName(files, arg)

	// Try to get file by the ID
	fileByID, errByID := byID(files, arg)

	foundByName := errByName == nil
	foundByID := errByID == nil

	switch {
	case options.ByName && foundByName:
		return compare(fileByName)
	case options.ByID && foundByID:
		return compare(fileByID)
	case !foundByName && foundByID:
		return compare(fileByID)
	case !foundByID && foundByName:
		return compare(fileByName)
	case !foundByName && !foundByID:
		return db.File{}, ErrNotFound
	}

	return
}
