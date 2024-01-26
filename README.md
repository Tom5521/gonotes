# GoNotes

A note program for the linux terminal!

## Features

- Support for temporary notes
- Configuration of the editor to be used

## Usage/Examples

You can see how to use it running `gonotes --help`

```bash
Usage: main [--open OPEN] [--delete DELETE] [--temporal] [--list] <command> [<args>]

Options:
  --open OPEN, -o OPEN   Open a file for editing or reading
  --delete DELETE, -d DELETE
                         Deletes a file
  --temporal, -t         Specifies whether the operation will be done in the temporary or constant directory.
  --list, -l             List all files, whether temporary or not
  --help, -h             display this help and exit

Commands:
  new                    Create a new text file
  config                 Configure some program variables
```

## Installation

Install my-project with go

```bash
go install github.com/Tom5521/GoNotes/cmd/gonotes@latest
```

Or you can copy the [binaries](https://github.com/Tom5521/GoNotes/releases/latest) for your OS into the path

## License

[MIT](https://choosealicense.com/licenses/mit/)
