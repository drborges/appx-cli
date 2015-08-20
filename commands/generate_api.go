package commands

import (
	"github.com/codegangsta/cli"
	"github.com/drborges/go-ast/appxcli"
	"log"
)

func NewGenerateAPI(app *cli.App) cli.Command {
	return cli.Command{
		Name: "api",
		Usage: "TODO for chain API",
		Action: func(c *cli.Context) {
			if pkg := c.Args().First(); pkg != "" {
				appxcli.Process(pkg)
			} else {
				log.Fatal("Expected a package path to look for appx model definitions.")
			}
		},
	}
}
