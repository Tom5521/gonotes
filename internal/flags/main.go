package flags

import (
	flag "github.com/spf13/pflag"
)

var (
	Help       = flag.Bool("help", false, "Print this message")
	Type       = flag.String("type", "txt", "Specifies the file type, by default txt")
	Delete     = flag.String("del", "", "Deletes a note by specifying its id or name, e.g. gonotes --del 12/note-name")
	Open       = flag.String("open", "", "Edit a note clarifying its name, i.e. gonotes --open [note-name] (you can also put its id)")
	Temporal   = flag.Bool("tmp", false, "The note will be saved in /tmp, i.e., it will be deleted on reboot.")
	Config     = flag.Bool("config", false, "It is used to clarify that you want to edit a configuration.")
	SetDefault = flag.String("set-editor", "nano", "Set the default editor.")
	New        = flag.String("new", "", "Create a new note, usage: gonotes --new [note name], the type of the file is specified with --type")
	Log        = flag.Bool("log", false, "Prints a list of the notes with their name and id")
)
