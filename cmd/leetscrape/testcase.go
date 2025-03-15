package main

import (
	"github.com/urfave/cli/v2"
	"github.com/vldcreation/leet-scrape/v2/data/service"
	"github.com/vldcreation/leet-scrape/v2/domain/usecase"
)

var testcase = &cli.Command{
	Name:    TESTCASE,
	Aliases: []string{"tc"},
	Usage:   "Download the test case sample as in/out file",
	Action: func(c *cli.Context) error {
		args := extractFlagArgs(c)
		ques, tcs, err := getTestCase(args)
		if err != nil {
			return exitCli(err)
		}
		fw := service.NewFileWriter()
		generateFile := usecase.NewGenerateTestCaseFile(fw, ques, tcs, args.fielaname, args.path)
		err = generateFile.Execute()
		if err != nil {
			return exitCli(err)
		}
		return nil
	},
}
