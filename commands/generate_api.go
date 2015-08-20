package commands

import (
	"github.com/codegangsta/cli"
	"github.com/drborges/go-ast/appxcli"
	"log"
	"os"
)

func NewGenerateAPI(app *cli.App) cli.Command {
	return cli.Command{
		Name:  "api",
		Usage: "Generates Appx API for appx models detected in a given path",
		Flags: []cli.Flag{cli.StringFlag{
			Name:  "path, p",
			Usage: "Use this flag to change the path where appx models will be searched for",
		}},
		Action: func(c *cli.Context) {
			path := c.String("path")
			if path == "" {
				dir, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}
				path = dir
			}
			appxcli.Process(path)
		},
	}
}
