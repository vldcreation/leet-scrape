package main

import (
	"github.com/ISKalsi/leet-scrape/v2/data/service"
	"github.com/ISKalsi/leet-scrape/v2/domain/usecase"
	"github.com/urfave/cli/v2"
)

var solution = &cli.Command{
	Name:    SOLUTION,
	Aliases: []string{"sol"},
	Usage:   "Download the starter-solution-snippet file in your desired language format",
	Action: func(c *cli.Context) error {
		args := extractFlagArgs(c)
		ques, err := getQuestion(args)
		if err != nil {
			return exitCli(err)
		}

		if err = validateTemplate(args.template); err != nil {
			return exitCli(err)
		}

		fileName, err := getFileName(ques, args)
		if err != nil {
			return exitCli(err)
		}
		fw := service.NewFileWriter()
		generateFile := usecase.NewGenerateSolutionFile(ques, fw, fileName, args.path, args.boilerplate, args.lang, args.template)
		err = generateFile.Execute()
		if err != nil {
			return exitCli(err)
		}
		return nil
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    BOILERPLATE,
			Aliases: []string{"b"},
			Usage:   "Add boilerplate code to the Solution file at the top. Eg: " + CliName + " -n \"Two Sum\" sol -b \"#include<iostream>\\n using namespace std\\n\"",
		},
		&cli.StringFlag{
			Name:    TEMPLATE,
			Aliases: []string{"t"},
			Usage:   "Set template code snippet. Eg: easy | medium | hard",
		},
		&cli.StringFlag{
			Name:    LOCATION,
			Aliases: []string{"o"},
			Usage:   "Directory `<path>` for the output file",
		},
		&cli.StringFlag{
			Name:     LANGUAGE,
			Aliases:  []string{"l"},
			Usage:    "Generally available options: C++, C, C#, Kotlin, Java, Python, Python3, Swift, \n\t\tGo, PHP, Racket, Rust, Ruby, JavaScript, TypeScript, Scala, ErLang, Elixir",
			Required: true,
		},
	},
}
