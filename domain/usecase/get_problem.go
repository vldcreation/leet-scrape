package usecase

import (
	"github.com/vldcreation/leet-scrape/v2/domain/entity"
	"github.com/vldcreation/leet-scrape/v2/domain/repo"
	"github.com/vldcreation/leet-scrape/v2/internal/errors"
)

type searchingMethod int

const (
	byName searchingMethod = iota
	byNumber
	byUrl
	problemOfDay
)

type GetProblem struct {
	repo   repo.Problem
	search searchingMethod
	num    int
	url    string
	name   string
}

func NewGetProblemByName(scrapper repo.Problem, name string) *GetProblem {
	return &GetProblem{
		repo:   scrapper,
		search: byName,
		name:   name,
	}
}

func NewGetProblemByNumber(scrapper repo.Problem, num int) *GetProblem {
	return &GetProblem{
		repo:   scrapper,
		search: byNumber,
		num:    num,
	}
}

func NewGetProblemByUrl(scrapper repo.Problem, url string) *GetProblem {
	return &GetProblem{
		repo:   scrapper,
		search: byUrl,
		url:    url,
	}
}

func NewGetDailyChallenge(scrapper repo.Problem) *GetProblem {
	return &GetProblem{
		repo:   scrapper,
		search: problemOfDay,
	}
}

func (uc *GetProblem) Execute() (*entity.Question, error) {
	var question *entity.Question
	var err error
	switch uc.search {
	case byName:
		question, err = uc.repo.GetByName(uc.name)
	case byUrl:
		question, err = uc.repo.GetByUrl(uc.url)
	case byNumber:
		question, err = uc.repo.GetByNumber(uc.num)
	case problemOfDay:
		question, err = uc.repo.GetDailyChallenge()
	default:
		return nil, errors.InvalidSearchMethod
	}
	return question, err
}
