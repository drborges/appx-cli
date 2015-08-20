package main

import (
	"github.com/codegangsta/cli"
	"github.com/drborges/go-ast/commands"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "appx-cli"
	app.Usage = "$ appx generate [model|api]"
	app.Action = func(c *cli.Context) {
		println(app.Usage)
	}

	app.Commands = []cli.Command{
		{
			Name:  "generate",
			Usage: app.Usage,
			Subcommands: []cli.Command{
				commands.NewGenerateModel(app),
				commands.NewGenerateAPI(app),
			},
			Action: func(c *cli.Context) {
				println(app.Usage)
			},
		},
	}

	app.Run(os.Args)
}
