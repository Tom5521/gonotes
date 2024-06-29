//go:build exp
// +build exp

package search

import (
	"errors"
	"strconv"

	"github.com/Tom5521/gonotes/internal/db"
	"github.com/Tom5521/gonotes/internal/options"
	"github.com/Tom5521/slicelib"
)

func Deep(arg string) (f db.File, err error) {

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

	hasCustomName := options.CustomName != ""
	hasCustomID := options.CustomID != -1

	f2 := files.Copy()

	// Filter 2: By name.
	f2.Filter(func(f db.File) bool {
		return f.Name == arg
	})
	// Filter 3: By ID.
	if f2.IsEmpty() {
		f2 = files
		var id uint64
		id, err = strconv.ParseUint(arg, 10, strconv.IntSize)
		if err != nil {
			return
		}
		f2.Filter(func(f db.File) bool {
			return uint64(f.ID) == id
		})
		if f2.IsEmpty() {
			return f, errors.New("name and ID could not be found")
		}
	}

	if f2.Len() == 1 {
		return f2.Elem(0), nil
	}

	// Filter 3: By filetype
	f3 := files.Copy()
	f3.Filter(func(f db.File) bool {
		return f.Type == options.Filetype
	})
	if f3.IsEmpty() {
		return f, errors.New("could not filter by file type")
	}

	return
}
