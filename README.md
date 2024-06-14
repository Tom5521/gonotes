# GoNotes

A note program for the linux terminal!

## Features

- Support for temporary notes
- Configuration of the editor to be used

## Usage/Examples

You can see how to use it running `gonotes --help`

```bash
A CLI that allows you to manipulate and manage notes from your terminal using your favorite editor.

Usage:
  gonotes [command]

Available Commands:
  cat               Performs golang\'s equivalent of "cat" from unix shells to the file.
  completion        Generate the autocompletion script for the specified shell
  config            Manages the possible program configurations.
  delete            Deletes a file by its name.
  gen-markdown-tree Generate markdown documentation into the gived directory.
  help              Help about any command
  licence           Prints the current program license.
  list              Lists all files detected in normal and temporal storage.
  new               Create a new note.
  open              Open a existent file
  print-settings
  search            Look for a note specifying specific patterns.

Flags:
      --by-id           Specifies whether to search by id.
      --by-name         Specifies whether to search by name.
      --editor string   Specifies the editor to use. (default "nvim")
  -h, --help            help for gonotes
      --id int          Specifies the id to be searched for with the argument. (default -1)
      --name string     Specifies the name to be searched for with the argument.
      --normal          Perform the operation on a file that is specifically located in a normal directory.
      --tmp             Perform the operation on a file that is specifically located in the temporary directory.
      --type string     Specifies the file type. (default ".txt")

Use "gonotes [command] --help" for more information about a command.
```

## Installation

Install my-project with go

```bash
go install -v github.com/Tom5521/gonotes@latest
```

Or you can copy the [binaries](https://github.com/Tom5521/GoNotes/releases/latest) for your OS into the path

## License

[MIT](https://choosealicense.com/licenses/mit/)
