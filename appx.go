package main

import (
	"github.com/codegangsta/cli"
	"github.com/drborges/go-ast/commands"
	"os"
	"log"
)

func main() {
	app := cli.NewApp()
	app.Name = "appx"
	app.Usage = "$ appx generate [model|api]"
	app.Action = func(c *cli.Context) {
		println(app.Usage)
	}

	app.Commands = []cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"gen"},
			Usage:   app.Usage,
			Subcommands: []cli.Command{
				commands.NewGenerateModel(app),
				commands.NewGenerateAPI(app),
			},
			Action: func(c *cli.Context) {
				log.Println(app.Usage)
			},
		},
	}

	app.Run(os.Args)
}
