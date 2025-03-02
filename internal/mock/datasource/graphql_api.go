package datasource

import (
	"github.com/stretchr/testify/mock"
	"github.com/vldcreation/leet-scrape/v2/data/model"
)

type GraphQLApi struct {
	mock.Mock
}

func (g *GraphQLApi) FetchBySlug(titleSlug string) (*model.QuestionQuery, error) {
	args := g.Called(titleSlug)
	r0, _ := args.Get(0).(*model.QuestionQuery)
	r1 := args.Error(1)
	return r0, r1
}

func (g *GraphQLApi) FetchByNumber(id string) (*model.QuestionListQuery, error) {
	args := g.Called(id)
	r0, _ := args.Get(0).(*model.QuestionListQuery)
	r1 := args.Error(1)
	return r0, r1
}

func (g *GraphQLApi) FetchDailyChallengesOfMonth(year int, month int) (*model.DailyChallengesQuery, error) {
	args := g.Called(year, month)
	r0, _ := args.Get(0).(*model.DailyChallengesQuery)
	r1 := args.Error(1)
	return r0, r1
}
