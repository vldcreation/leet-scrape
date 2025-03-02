package main

import (
	"github.com/gocolly/colly/v2"
	"github.com/machinebox/graphql"
	"github.com/urfave/cli/v2"
	"github.com/vldcreation/leet-scrape/v2/api"
	"github.com/vldcreation/leet-scrape/v2/data/datasource"
	"github.com/vldcreation/leet-scrape/v2/data/repo"
	"github.com/vldcreation/leet-scrape/v2/domain/entity"
	"github.com/vldcreation/leet-scrape/v2/domain/usecase"
	"github.com/vldcreation/leet-scrape/v2/internal/errors"
)

func handleError(err errors.Error) cli.ExitCoder {
	return cli.Exit(err.GetMessage(CliName), err.GetCode())
}

func exitCli(err error) cli.ExitCoder {
	switch e := err.(type) {
	case errors.Error:
		return handleError(e)
	case nil:
		return nil
	default:
		return handleError(errors.Unexpected)
	}
}

func getQuestion(args *flagArgs) (*entity.Question, error) {
	client := graphql.NewClient(api.GraphqlApiUrl)
	graphqlClient := datasource.NewGraphQLApiImpl(client)
	collector := colly.NewCollector(colly.AllowedDomains("leetcode.com"))
	webScrapper := datasource.NewWebScrapperImpl(collector)
	repository := repo.NewProblem(graphqlClient, webScrapper)

	var getProblem *usecase.GetProblem
	if args.url != "" {
		getProblem = usecase.NewGetProblemByUrl(repository, args.url)
	} else if args.name != "" {
		getProblem = usecase.NewGetProblemByName(repository, args.name)
	} else if args.num != -1 {
		getProblem = usecase.NewGetProblemByNumber(repository, args.num)
	} else if args.today {
		getProblem = usecase.NewGetDailyChallenge(repository)
	} else {
		return nil, errors.FlagMissing
	}

	ques, err := getProblem.Execute()
	if err != nil {
		return nil, err
	}
	return ques, nil
}

func getFileName(q *entity.Question, args *flagArgs) (string, error) {
	if args.url != "" || args.name != "" {
		return q.TitleSlug, nil
	} else if args.num != -1 || args.today {
		return q.Id, nil
	} else {
		return "", errors.FlagMissing
	}
}
