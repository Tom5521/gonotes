package cmd

import "github.com/Tom5521/gonotes/internal/options"

func initFlags() {
	flags := root.PersistentFlags()
	flags.StringVar(&options.Filetype, "type", settings.String(DefaultTypeKey), "Specifies the file type.")
	flags.StringVar(&options.Editor, "editor", settings.String(DefaultEditorKey), "Specifies the editor to use.")

	flags.BoolVar(&options.ByID, "by-id", false, "Specifies whether to search by id.")
	flags.BoolVar(&options.ByName, "by-name", false, "Specifies whether to search by name.")

	flags.IntVar(&options.CustomID, "id", -1, "Specifies the id to be searched for with the argument.")
	flags.StringVar(&options.CustomName, "name", "",
		"Specifies the name to be searched for with the argument.",
	)
	root.MarkFlagsMutuallyExclusive("by-id", "by-name")

	flags.BoolVar(&options.Temporal, "tmp", settings.Bool(DefaultTmpKey),
		"Perform the operation on a file that is specifically located in the temporary directory.")
	flags.BoolVar(&options.Normal, "normal", settings.Bool(DefaultNormalKey),
		"Perform the operation on a file that is specifically located in a normal directory.",
	)
	root.MarkFlagsMutuallyExclusive("tmp", "normal")
}
