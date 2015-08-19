package main

import (
	"github.com/drborges/go-ast/appx"
	"github.com/codegangsta/cli"
	"os"
	"log"
)

func main() {
	app := cli.NewApp()
	app.Name = "appx-cli"
	app.Usage = "TODO for usage..."
	app.Action = func(c *cli.Context) {
		println(app.Usage)
	}

	app.Commands = []cli.Command{
		{
			Name:      "model",
			Usage:     "Generate code for a new model",
			Action: func(c *cli.Context) {
				println("added task: ", c.Args().First())
			},
		},
		{
			Name:      "api",
			Usage:     "Generate the chain api code for a given model",
			Action: func(c *cli.Context) {
				if pkg := c.Args().First(); pkg != "" {
					appx.Process(pkg)
				} else {
					log.Fatal("Expected a package path to look for appx model definitions.")
				}
			},
		},
	}

	app.Run(os.Args)
}
