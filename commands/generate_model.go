package commands

import (
	"github.com/codegangsta/cli"
	"log"
)

func NewGenerateModel(app *cli.App) cli.Command {
	return cli.Command{
		Name:  "model",
		Usage: "Generate code for a new model",
		Flags: []cli.Flag{cli.BoolFlag{
			Name: "api",
			Usage: "Use this flag in order to generate the chainable API for the corresponding model",
		}},
		Action: func(c *cli.Context) {
			log.Println("added task: ", c.Args().First())
		},
	}
}
