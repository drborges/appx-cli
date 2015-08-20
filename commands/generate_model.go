package commands

import "github.com/codegangsta/cli"

func NewGenerateModel(app *cli.App) cli.Command {
	return cli.Command{
		Name:  "model",
		Usage: "Generate code for a new model",
		Action: func(c *cli.Context) {
			println("added task: ", c.Args().First())
		},
	}
}
