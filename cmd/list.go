package cmd

import (
	"fmt"

	"github.com/Tom5521/gonotes/internal/db"
	"github.com/Tom5521/gonotes/internal/options"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

func initList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists all files detected in normal and temporal storage.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			printNormal := func() {
				color.Red.Println("**Normal Files**")
				for _, f := range db.Files {
					if f.Temporal {
						continue
					}
					fmt.Println(f)
				}
			}

			printTmp := func() {
				color.Red.Println("**Temporal Files**")
				for _, f := range db.Files {
					if !f.Temporal {
						continue
					}
					fmt.Println(f)
				}
			}

			switch {
			case options.Temporal:
				printTmp()
			case options.Normal:
				printNormal()
			default:
				printNormal()
				printTmp()
			}
			return
		},
	}

	return cmd
}
