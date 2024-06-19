package search

import (
	"strconv"

	"github.com/Tom5521/gonotes/internal/db"
	"github.com/Tom5521/gonotes/internal/options"
)

func Deep(arg string) (file db.File, err error) {
	switch {
	case options.ByID:
		var id int
		id, err = strconv.Atoi(arg)
		if err != nil {
			return
		}
		file, err = ByID(uint(id))
	case options.ByName:
		file, err = ByName(arg)
	case options.CustomID >= 0 || options.CustomName != "":
		file, err = ByID(uint(options.CustomID))
		if err != nil {
			file, err = ByName(options.CustomName)
			if err != nil {
				file, err = ByName(arg)
			}
		}
	default:
		file, err = ByName(arg)
		if err != nil {
			var id int
			id, err = strconv.Atoi(arg)
			if err != nil {
				return file, ErrNotFound
			}
			file, err = ByID(uint(id))
		}

	}
	if (options.CustomID != -1 || options.CustomName != "") &&
		(file.ID != uint(options.CustomID) || file.Name != options.CustomName) {
		err = ErrNotFound
	}
	if options.Temporal && !file.Temporal {
		err = ErrNotFound
	}
	if options.Normal && file.Temporal {
		err = ErrNotFound
	}

	return
}
