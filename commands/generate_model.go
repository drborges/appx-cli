package commands

import (
	"github.com/codegangsta/cli"
	"github.com/drborges/go-ast/appxcli"
	"log"
	"strings"
	"os"
)

//
// $ appx gen model User \
//       --field Name:string \
//       --field Email:string \
//       --field Token:string \
//       --query-by Name \
//       --cache-by Token \
//       --api
//
func NewGenerateModel(app *cli.App) cli.Command {
	return cli.Command{
		Name:  "model",
		Usage: "Generate code for a new model",
		Flags: []cli.Flag{
			cli.StringSliceFlag{
				Name:  "field, f",
				Usage: "Defines a field of a model in the format FieldName:FieldType",
			},
			cli.StringFlag{
				Name:  "query-by, q",
				Usage: "If defined the model will implement appx.Queryable providing a datastore query by the given field",
			},
			cli.StringFlag{
				Name:  "cache-by, c",
				Usage: "If defined the model will implement appx.Cacheable by the given field value. If the flag value is wrapped within double quotes, cache ID will be the given string",
			},
			cli.StringFlag{
				Name:  "kind, k",
				Usage: "Override the entity kind value. By default it is the model's name",
			},
			cli.BoolFlag{
				Name:  "has-parent, p",
				Usage: "Whether or not the model's key has a parent key",
			},
			cli.StringFlag{
				Name:  "id, i",
				Usage: "The model field to use as the entity key's id component. If not specified the key is then marks as incomplete (auto generated by datastore)",
			},
			cli.StringFlag{
				Name:  "out, o",
				Usage: "Where to spit the generated code",
			},
			cli.BoolFlag{
				Name:  "api, a",
				Usage: "Use this flag in order to generate the chainable API for the corresponding model",
			},
		},
		Action: func(c *cli.Context) {
			id := c.String("id")
			name, pkg := "", "main"
			pkgAndModel := strings.Split(c.Args().First(), ".")
			if len(pkgAndModel) == 1 {
				name = pkgAndModel[0]
			} else if len(pkgAndModel) == 2 {
				pkg = pkgAndModel[0]
				name = pkgAndModel[1]
			} else {
				log.Fatalln("Model parameter must be in the form [package.]ModelName. Ex.: models.User, models.Post")
			}

			hasParent := c.Bool("has-parent")
			fieldsToTypes := parseFields(c.StringSlice("field"))
			intID := fieldsToTypes[id] != "string" && id != ""
			stringID := fieldsToTypes[id] == "string" && id != ""
			cacheBy := c.String("cache-by")
			cacheable := cacheBy != ""
			queryBy := c.String("query-by")
			queryable := queryBy != ""
			outPath := c.String("out")
			if outPath == "" {
				dir, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}
				outPath = dir
			}

			kind := name
			if c.String("kind") != "" {
				kind = c.String("kind")
			}
			err := appxcli.GenerateModel(outPath, &appxcli.Model{
				ID:         id,
				IntID:      intID,
				StringID:   stringID,
				Name:       name,
				Pkg:        pkg,
				Kind:       kind,
				Fields:     fieldsToTypes,
				Cacheable:  cacheable,
				CacheBy:    cacheBy,
				Queryable:  queryable,
				QueryBy:    queryBy,
				Incomplete: id == "",
				HasParent:  hasParent,
			})

			if err != nil {
				log.Fatal(err)
			}

			// TODO generate API code if --api flag is specified
		},
	}
}

func parseFields(fields []string) map[string]string {
	fieldsToTypes := make(map[string]string)
	for _, f := range fields {
		parts := strings.Split(f, ":")
		if len(parts) != 2 {
			log.Fatalln("Field definition must be in the format: FieldName:FieldType. Ex.: Name:string Age:int")
		}
		fieldsToTypes[parts[0]] = parts[1]
	}
	return fieldsToTypes
}
