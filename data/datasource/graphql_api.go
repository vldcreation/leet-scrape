package datasource

import (
	"context"

	"github.com/machinebox/graphql"
	"github.com/vldcreation/leet-scrape/v2/api"
	"github.com/vldcreation/leet-scrape/v2/data/model"
)

type GraphQLApi interface {
	FetchBySlug(titleSlug string) (*model.QuestionQuery, error)
	FetchByNumber(id string) (*model.QuestionListQuery, error)
	FetchDailyChallengesOfMonth(year int, month int) (*model.DailyChallengesQuery, error)
}

type GraphQLApiImpl struct {
	client *graphql.Client
}

func NewGraphQLApiImpl(c *graphql.Client) *GraphQLApiImpl {
	return &GraphQLApiImpl{client: c}
}

func (g *GraphQLApiImpl) FetchBySlug(titleSlug string) (*model.QuestionQuery, error) {
	query := api.GetQuery(api.Question)
	req := graphql.NewRequest(query)
	req.Var("titleSlug", titleSlug)
	req.Header.Set("Content-Type", "application/json")

	var response model.QuestionQuery
	err := g.client.Run(context.Background(), req, &response)
	if err != nil {
		return nil, err
	} else {
		return &response, nil
	}
}

func (g *GraphQLApiImpl) FetchByNumber(id string) (*model.QuestionListQuery, error) {
	query := api.GetQuery(api.QuestionList)
	req := graphql.NewRequest(query)
	req.Var("categorySlug", "")
	req.Var("limit", 1)
	req.Var("skip", 0)
	req.Var("filters", map[string]string{
		"searchKeywords": id,
	})
	req.Header.Set("Content-Type", "application/json")

	var response model.QuestionListQuery
	err := g.client.Run(context.Background(), req, &response)
	if err != nil {
		return nil, err
	} else {
		return &response, nil
	}
}

func (g *GraphQLApiImpl) FetchDailyChallengesOfMonth(year int, month int) (*model.DailyChallengesQuery, error) {
	query := api.GetQuery(api.DailyChallenges)
	req := graphql.NewRequest(query)
	req.Var("year", year)
	req.Var("month", month)
	req.Header.Set("Content-Type", "application/json")

	var response model.DailyChallengesQuery
	err := g.client.Run(context.Background(), req, &response)
	if err != nil {
		return nil, err
	} else {
		return &response, nil
	}
}
