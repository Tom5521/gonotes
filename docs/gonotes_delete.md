## gonotes delete

Deletes a file by its name.

### Synopsis

Delete a file by its name, an error will occur if there are two files with the same name in the normal or temporary storage, also if there are two files with the same name but different file type. In those cases you must specify with a flag.

```
gonotes delete [flags]
```

### Options

```
  -h, --help   help for delete
```

### Options inherited from parent commands

```
      --by-id           Specifies whether to search by id.
      --by-name         Specifies whether to search by name.
      --editor string   Specifies the editor to use. (default "nvim")
      --id int          Specifies the id to be searched for with the argument. (default -1)
      --name string     Specifies the name to be searched for with the argument.
      --normal          Perform the operation on a file that is specifically located in a normal directory.
      --tmp             Perform the operation on a file that is specifically located in the temporary directory.
      --type string     Specifies the file type. (default ".txt")
```

### SEE ALSO

* [gonotes](gonotes.md)	 - A note manager for the terminal

###### Auto generated by spf13/cobra on 19-Jun-2024