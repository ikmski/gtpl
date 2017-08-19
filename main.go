package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func mainAction(c *cli.Context) error {

	if c.NArg() != 1 {
		return fmt.Errorf("A template is required as an argument")
	}
	tmpl := c.Args()[0]
	t := template.New("t")
	template.Must(t.Parse(tmpl))

	var paramMap map[string]string = make(map[string]string)
	for _, p := range c.StringSlice("params") {

		ss := strings.Split(p, "=")
		if len(ss) > 0 {

			key := ss[0]
			var val string = ""
			if len(ss) > 1 {
				val = ss[1]
			}

			paramMap[key] = val
		}
	}

	t.Execute(os.Stdout, paramMap)

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "gtpl"
	app.Usage = "simple template"

	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "params, p",
			Usage: "parameters for the template. ex) key=value",
		},
	}

	app.Action = mainAction

	app.Run(os.Args)
}
