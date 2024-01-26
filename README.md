# GoNotes

A note program for the linux terminal!

## Features

- Support for temporary notes
- Configuration of the editor to be used

## Usage/Examples

You can see how to use it running `gonotes --help`

```bash
Usage: main [--open FILE] [--del FILE] [--tmp] [--list] [--print FILE] <command> [<args>]

Options:
  --open FILE, -o FILE   Open a file for editing or reading
  --del FILE, -d FILE    Deletes a file
  --tmp, -t              Specifies whether the operation will be done in the temporary or constant directory.
  --list, -l             List all files, whether temporary or not
  --print FILE, -p FILE
                         Print the file showing its details, without opening an editor.
  --help, -h             display this help and exit
  --version              display version and exit

Commands:
  new                    Create a new text file
  config                 Configure some program variables
```

## Installation

Install my-project with go

```bash
go install -v github.com/Tom5521/GoNotes/cmd/gonotes@latest
```

Or you can copy the [binaries](https://github.com/Tom5521/GoNotes/releases/latest) for your OS into the path

## License

[MIT](https://choosealicense.com/licenses/mit/)
