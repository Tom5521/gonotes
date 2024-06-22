//go:build !old
// +build !old

package search

import (
	"errors"
	"strconv"

	"github.com/Tom5521/gonotes/internal/db"
	"github.com/Tom5521/gonotes/internal/options"
)

func Deep(arg string) (file db.File, err error) {
	var files []db.File

	// Filter 1: File location.
	switch {
	case options.Temporal:
		files = db.TemporalFiles()
	case options.Normal:
		files = db.NormalFiles()
	default:
		files = db.Files
	}

	unfilteredFiles := files

	var found bool

	var nameIsProvided = options.CustomName != ""
	var idIsProvided = options.CustomID >= 0

	var argIsName = !options.ByID || !nameIsProvided && idIsProvided || options.ByName

	// Filter 2: Name.
	if argIsName {
		name := arg
		if nameIsProvided {
			name = options.CustomName
		}
		files, found = filterByName(name, files)
		if !found {
			files = unfilteredFiles
		}
	}

	// Filter 3: ID.
	if !found || !argIsName {
		var id = uint64(options.CustomID)
		if !idIsProvided {
			id, err = atoui(arg)
			if err != nil {
				return
			}
		}
		files, found = filterByID(id, files)
	}
	if !found {
		return file, errors.New("file name/id not found")
	}
	if len(files) == 1 {
		return files[0], nil
	}

	// Filter 4: File type.
	files, found = filterByType(options.Filetype, files)
	if len(files) == 1 {
		return files[0], nil
	}
	if !found {
		return file, ErrNotFound
	}

	return
}

func filterByName(name string, files []db.File) (filteredFiles []db.File, found bool) {
	for _, f := range files {
		if f.Name == name {
			filteredFiles = append(filteredFiles, f)
		}
	}
	found = len(filteredFiles) == 1
	return
}

func filterByID(id uint64, files []db.File) (filteredFiles []db.File, found bool) {
	for _, f := range files {
		if uint64(f.ID) == id {
			filteredFiles = append(filteredFiles, f)
		}
	}
	found = len(filteredFiles) == 1
	return
}

func filterByType(t string, files []db.File) (filteredFiles []db.File, found bool) {
	for _, f := range files {
		if f.Type == t {
			filteredFiles = append(filteredFiles, f)
		}
	}
	found = len(filteredFiles) == 1
	return
}

func atoui(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, strconv.IntSize)
}

/*
func Deep(arg string) (file db.File, err error) {
	var files []db.File

	// Filter 1: File location.
	switch {
	case options.Temporal:
		files = db.TemporalFiles()
	case options.Normal:
		files = db.NormalFiles()
	default:
		files = db.Files
	}

	var argIsName bool
	if !options.ByID || options.CustomID != -1 || options.ByName {
		argIsName = true
	}

	// Filter 2: File name.
	if argIsName {
		files = func() (newFiles []db.File) {
			for _, f := range files {
				if f.Name == arg {
					newFiles = append(newFiles, f)
				}
			}
			return
		}()
	}

	// Filter 3: ID.
	if !argIsName {
		files = func() (newFiles []db.File) {
			var id uint64
			if options.CustomID != -1 {
				id = uint64(options.CustomID)
			} else {
				id, err = strconv.ParseUint(arg, 10, strconv.IntSize)
				if err != nil {
					return files
				}
			}
			for _, f := range files {
				if uint64(f.ID) == id {
					newFiles = append(newFiles, f)
				}
			}
			return
		}()
		if err != nil {
			return
		}
		if len(files) == 0 {
			err = errors.New("ID not found")
		}
	}

	if len(files) > 1 {
		// Filter 4: File type.
		files = func() (newFiles []db.File) {
			for _, f := range files {
				if f.Type == options.Filetype {
					newFiles = append(newFiles, f)
				}
			}
			return
		}()
		if len(files) == 0 {
			err = errors.New("")
		}
	} else if files[0].Type != options.Filetype {
		err = fmt.Errorf(
			"types do not match, type requested: %s | type found: %s",
			options.Filetype,
			files[0].Type,
		)
	}

	if err != nil {
		return
	}

	file = files[0]

	/*
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
	* /
	return
}

*/
