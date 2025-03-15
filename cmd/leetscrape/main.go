package main

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/urfave/cli/v2"
)

const (
	URL    = "url"
	NAME   = "name"
	NUMBER = "number"
	TODAY  = "daily-challenge"

	LOCATION = "output-dir"

	QUESTION = "question"
	SOLUTION = "solution"

	BOILERPLATE = "boilerplate"
	LANGUAGE    = "lang"

	TEMPLATE = "template"
	VERSION  = "version"

	TESTCASE = "testcase"
	FILENAME = "filename"
)

const CliName = "leetscrape"

var Version = "development"
var Commit = "none"
var Date = "unknown"

func printVersion() {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("Unable to determine version information.")
		return
	}

	if buildInfo.Main.Version != "" {
		Version = buildInfo.Main.Version
		Commit = buildInfo.Main.Sum
	}
	fmt.Printf("%s version %s\ncommit %s \nbuilt at %s\n", CliName, Version, Commit, Date)
}

func main() {
	app := &cli.App{
		Name:      "Leetcode Scrapper",
		Usage:     "Download and create the default empty solution file (with the question statement as docstring at the top) from leetcode.com",
		UsageText: "leetscrape [global options] command [command options]\n    Examples -\n\t1. " + CliName + " --name \"Two Sum\" solution --lang C++\n\t2. " + CliName + " -N 455 question",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    URL,
				Aliases: []string{"u"},
				Usage:   "Search problem by its `<" + URL + ">`.Eg: " + CliName + " -u https://leetcode.com/problems/two-sum sol",
			},
			&cli.StringFlag{
				Name:    NAME,
				Aliases: []string{"n"},
				Usage:   "Search problem by its `<" + NAME + ">`. Eg: " + CliName + " -n \"two sum\" sol\n\tNote: `<" + NAME + ">` should have double quotes if it contains any whitespace",
			},
			&cli.IntFlag{
				Name:    NUMBER,
				Value:   -1,
				Aliases: []string{"N"},
				Usage:   "Search problem by its `<" + NUMBER + ">`.Eg: " + CliName + " -N 1 sol",
			},
			&cli.BoolFlag{
				Name:    TODAY,
				Aliases: []string{"d"},
				Usage:   "Get problem of the day. Eg: " + CliName + " -d sol -l C++",
			},
			&cli.StringFlag{
				Name:    LOCATION,
				Aliases: []string{"o"},
				Usage:   "Directory `<path>` for the output file",
			},
			&cli.BoolFlag{
				Name:    VERSION,
				Aliases: []string{"v"},
				Usage:   "Print version",
				Action: func(ctx *cli.Context, b bool) error {
					printVersion()
					os.Exit(0)
					return nil
				},
			},
		},
		Commands: []*cli.Command{question, solution, testcase},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
