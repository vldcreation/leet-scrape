package mock

import (
	"github.com/stretchr/testify/mock"
	"github.com/vldcreation/leet-scrape/v2/domain/entity"
)

type ProblemScrapper struct {
	mock.Mock
}

func (s *ProblemScrapper) GetByName(name string) (*entity.Question, error) {
	args := s.Called(name)
	ques, _ := args.Get(0).(*entity.Question)
	err := args.Error(1)
	return ques, err
}

func (s *ProblemScrapper) GetByUrl(url string) (*entity.Question, error) {
	args := s.Called(url)
	ques, _ := args.Get(0).(*entity.Question)
	err := args.Error(1)
	return ques, err
}

func (s *ProblemScrapper) GetByNumber(num int) (*entity.Question, error) {
	args := s.Called(num)
	ques, _ := args.Get(0).(*entity.Question)
	err := args.Error(1)
	return ques, err
}
