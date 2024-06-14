package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

const (
	NormalPathKey    = "normal-path"
	TemporalPathKey  = "temporal-path"
	DefaultTmpKey    = "default-tmp"
	DefaultNormalKey = "default-normal"
	DefaultTypeKey   = "default-type"
	DefaultEditorKey = "default-editor"

	AppName = "gonotes"
)

var (
	Keys = [][]string{
		BoolKeys,
		StringKeys,
	}
	BoolKeys = []string{
		DefaultTmpKey,
		DefaultNormalKey,
	}
	StringKeys = []string{
		NormalPathKey,
		TemporalPathKey,
		DefaultEditorKey,
		DefaultTypeKey,
	}
)

func initConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manages the possible program configurations.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			errNotFound := errors.New("key not found")
			for _, arg := range args {
				parts := strings.SplitN(arg, "=", 2)
				if len(parts) < 1 {
					return errors.New("syntax error")
				}
				key, value := parts[0], parts[1]

				var foundKey bool
				for _, keytype := range Keys {
					for _, dkey := range keytype {
						if key == dkey {
							foundKey = true
							break
						}
					}
				}
				if !foundKey {
					return errNotFound
				}
				var isString, isBool bool
				for _, k := range StringKeys {
					if k == key {
						isString = true
						break
					}
				}
				for _, k := range BoolKeys {
					if k == key {
						isBool = true
						break
					}
				}

				if !isString && !isBool {
					return errNotFound
				}

				if isString {
					settings.SetString(key, value)
				}
				if isBool {
					b, err := strconv.ParseBool(value)
					if err != nil {
						return err
					}
					settings.SetBool(key, b)
				}
			}
			return nil
		},
	}

	return cmd
}

func initPrintSettings() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "print-settings",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			print := func(key string) {
				v := settings.Read(key)
				fmt.Print(color.Green.Render(key), ": ")
				b, ok := v.(bool)
				if ok {
					if b {
						color.Green.Println(b)
					} else {
						color.Red.Println(b)
					}
				} else {
					fmt.Println(v)
				}
			}
			color.Bold.Println("Boolean Keys")
			for _, k := range BoolKeys {
				print(k)
			}
			color.Bold.Println("String Keys")
			for _, k := range StringKeys {
				print(k)
			}
		},
	}

	return cmd
}
