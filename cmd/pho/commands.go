// Package main provides main  
package main

import (
	"os"

	"github.com/teris-io/cli"
)

// ParseArgs function  
func ParseArgs(args []string) (*Config, error) {
	c := new(Config)
	app := cli.New("line length analysis").
		WithOption(cli.NewOption("recurse", "Recurse into directories").WithChar('r').WithType(cli.TypeBool)).
		WithOption(cli.NewOption("display", "Display method").WithChar('d').WithType(cli.TypeString)).
		WithArg(cli.NewArg("filepaths", "Filepaths to analyse").AsOptional()).
		WithAction(func(args []string, options map[string]string) int {
			_, c.Recurse = options["recurse"]
			c.DisplayMethod = options["display"]
			if c.DisplayMethod == "" {
				c.DisplayMethod = "bar"
			}
			c.Filepaths = args
			return 0
		})
	app.Run(append([]string{""}, args...), os.Stdout)
	return c, nil
}

var displayMethods = [4]string{"line", "box", "stdout", "line-cli"}
