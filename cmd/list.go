package cmd

import (
	"fmt"

	"github.com/Tom5521/gonotes/internal/files"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

func initList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists all files detected in normal and temporal storage.",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			printNormal := func() {
				color.Red.Println("**Normal Files**")
				for _, f := range files.Files {
					if f.Temporal {
						continue
					}
					fmt.Println(f)
				}
			}

			printTmp := func() {
				color.Red.Println("**Temporal Files**")
				for _, f := range files.Files {
					if !f.Temporal {
						continue
					}
					fmt.Println(f)
				}
			}

			if !temporal && !normal {
				printNormal()
				printTmp()
			}
			if temporal && !normal {
				printTmp()
			}
			if normal && !temporal {
				printNormal()
			}
			return
		},
	}

	return cmd
}
