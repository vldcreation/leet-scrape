package repo

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/vldcreation/leet-scrape/v2/data/datasource"
	"github.com/vldcreation/leet-scrape/v2/domain/entity"
	"github.com/vldcreation/leet-scrape/v2/internal/errors"
	"github.com/vldcreation/leet-scrape/v2/internal/util"
)

type Problem struct {
	graphQLClient datasource.GraphQLApi
	scrapper      datasource.WebScrapper
}

func NewProblem(client datasource.GraphQLApi, scrapper datasource.WebScrapper) *Problem {
	return &Problem{graphQLClient: client, scrapper: scrapper}
}

func (s *Problem) GetByName(name string) (*entity.Question, error) {
	nameSlug := util.ConvertToSlug(name)
	response, err := s.graphQLClient.FetchBySlug(nameSlug)
	if err != nil {
		if strings.Contains(err.Error(), "query does not exist") {
			return nil, errors.QuestionNotFound
		} else {
			return nil, err
		}
	}
	return &response.Question, nil
}

func (s *Problem) GetByUrl(url string) (*entity.Question, error) {
	isLeetcodeUrl, _ := regexp.MatchString(`leetcode\.com`, url)
	if isLeetcodeUrl {
		problemSetRegex := regexp.MustCompile(`.*problems/([\w-]*)/?$`)
		isFromProblemSet := problemSetRegex.MatchString(url)
		if isFromProblemSet {
			captureGroups := problemSetRegex.SubexpNames()
			if len(captureGroups) == 2 {
				slug := problemSetRegex.ReplaceAllString(url, "$1")
				q, err := s.GetByName(slug)
				if err != nil {
					return nil, err
				} else {
					return q, nil
				}
			} else {
				return nil, errors.InvalidURL
			}
		} else {
			return nil, errors.LoginRequired
		}
	} else {
		return nil, errors.InvalidURL
	}
}

func (s *Problem) GetByNumber(num int) (*entity.Question, error) {
	numString := strconv.Itoa(num)
	response, err := s.graphQLClient.FetchByNumber(numString)
	if err != nil {
		return nil, err
	}
	if response.QuestionList.TotalNum == 0 {
		return nil, errors.QuestionIdOutOfRange
	} else {
		ques := response.QuestionList.Data[0]
		if ques.Id != numString {
			return nil, errors.QuestionIdOutOfRange
		} else {
			return &ques, nil
		}
	}
}

func (s *Problem) GetDailyChallenge() (*entity.Question, error) {
	year, month, _ := time.Now().Date()
	response, err := s.graphQLClient.FetchDailyChallengesOfMonth(year, int(month))
	if err != nil {
		return nil, err
	}
	totalChallenges := len(response.DailyCodingChallengeV2.Challenges)
	if totalChallenges == 0 {
		return nil, errors.NoDailyChallenge
	} else {
		challengeForToday := response.DailyCodingChallengeV2.Challenges[totalChallenges-1].Question
		return &challengeForToday, nil
	}
}
