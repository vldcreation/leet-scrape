package main

import (
	"github.com/urfave/cli/v2"
	"github.com/vldcreation/leet-scrape/v2/data/service"
	"github.com/vldcreation/leet-scrape/v2/domain/usecase"
)

var question = &cli.Command{
	Name:    QUESTION,
	Aliases: []string{"ques"},
	Usage:   "Download the question statement (including images, if any) as an HTML page",
	Action: func(c *cli.Context) error {
		args := extractFlagArgs(c)
		ques, err := getQuestion(args)
		if err != nil {
			return exitCli(err)
		}
		fw := service.NewFileWriter()
		id := service.NewImageDownloader()
		generateFile := usecase.NewGenerateQuestionFile(ques, args.path, fw, id)
		err = generateFile.Execute()
		if err != nil {
			return exitCli(err)
		}
		return nil
	},
}
