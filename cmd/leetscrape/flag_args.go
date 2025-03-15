package main

import "github.com/urfave/cli/v2"

type flagArgs struct {
	url         string
	name        string
	num         int
	today       bool
	boilerplate string
	fielaname   string
	path        string
	lang        string
	template    string
}

func extractFlagArgs(c *cli.Context) *flagArgs {
	url := c.String(URL)
	num := c.Int(NUMBER)
	name := c.String(NAME)
	today := c.Bool(TODAY)
	boilerplate := c.String(BOILERPLATE)
	filename := c.String(FILENAME)
	path := c.String(LOCATION)
	lang := c.String(LANGUAGE)
	template := c.String(TEMPLATE)

	return &flagArgs{
		url:         url,
		name:        name,
		num:         num,
		today:       today,
		boilerplate: boilerplate,
		path:        path,
		lang:        lang,
		template:    template,
		fielaname:   filename,
	}
}
