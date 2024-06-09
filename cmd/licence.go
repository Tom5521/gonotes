package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const mitLicence string = `MIT License

Copyright (c) 2023 Tom5521

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.`

func initLicence() *cobra.Command {
	var (
		short bool
		long  bool
	)

	var cmd = &cobra.Command{
		Use:   "licence",
		Short: "Prints the current program license.",
		Run: func(cmd *cobra.Command, args []string) {
			switch {
			case long:
				fmt.Println(mitLicence)
			default:
				fmt.Println("MIT")
			}
		},
	}
	flags := cmd.Flags()
	flags.BoolVar(&short, "short", false, "Prints only the name of the current license of the program.")
	flags.BoolVar(&long, "long", false, "Prints the entire licence of the program.")

	return cmd
}
