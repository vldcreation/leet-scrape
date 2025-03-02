package repo

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vldcreation/leet-scrape/v2/data/model"
	"github.com/vldcreation/leet-scrape/v2/domain/entity"
	internalErr "github.com/vldcreation/leet-scrape/v2/internal/errors"
	"github.com/vldcreation/leet-scrape/v2/internal/mock/datasource"
	"github.com/vldcreation/leet-scrape/v2/internal/testdata"
)

func TestGetByName(group *testing.T) {
	group.Run("should return valid Question using it's name", func(tt *testing.T) {
		testQues, _ := testdata.ImportFromFile("two_sum.json")
		testQuery := &model.QuestionQuery{
			Question: testQues,
		}
		data := []string{
			"Two sum",
			"two sum",
			"TWO SUM",
			"tWo SuM",
			"two-sum",
			"TWO-sum",
		}
		for _, testName := range data {
			tt.Run(testName, func(t *testing.T) {
				mockApi := &datasource.GraphQLApi{}
				mockScrapper := &datasource.WebScrapper{}
				mockApi.On("FetchBySlug", testQues.TitleSlug).Return(testQuery, nil)
				repository := NewProblem(mockApi, mockScrapper)

				actualQues, err := repository.GetByName(testName)

				assert.Nil(t, err)
				assert.Equal(t, &testQues, actualQues)
			})
		}
	})

	group.Run("should return no Question (nil) in case of an error from api", func(tt *testing.T) {
		testName := "Twwo sum"
		testSlug := "twwo-sum"

		tt.Run("when no question and random error is returned by the api", func(t *testing.T) {
			mockApi := &datasource.GraphQLApi{}
			mockScrapper := &datasource.WebScrapper{}
			mockApi.On("FetchBySlug", testSlug).Return(nil, errors.New("random error"))
			repository := NewProblem(mockApi, mockScrapper)

			ques, err := repository.GetByName(testName)

			assert.NotNil(t, err)
			assert.Nil(t, ques)
		})

		tt.Run("when a dummy question and random error is returned by the api", func(t *testing.T) {
			mockApi := &datasource.GraphQLApi{}
			mockScrapper := &datasource.WebScrapper{}
			mockApi.On("FetchBySlug", testSlug).Return(&model.QuestionQuery{}, errors.New("random error"))
			repository := NewProblem(mockApi, mockScrapper)

			ques, err := repository.GetByName(testName)

			assert.NotNil(t, err)
			assert.Nil(t, ques)
		})

		tt.Run("when no question and \"query does not exist\" error is returned by the api", func(t *testing.T) {
			mockApi := &datasource.GraphQLApi{}
			mockScrapper := &datasource.WebScrapper{}
			mockApi.On("FetchBySlug", testSlug).Return(nil, errors.New("query does not exist"))
			repository := NewProblem(mockApi, mockScrapper)

			ques, err := repository.GetByName(testName)

			assert.NotNil(t, err)
			assert.Nil(t, ques)
		})
	})

	group.Run("should return QuestionNotFound error when the query does not exist in the api", func(t *testing.T) {
		testName := "Twwo sum"
		testSlug := "twwo-sum"
		mockApi := &datasource.GraphQLApi{}
		mockScrapper := &datasource.WebScrapper{}
		mockApi.On("FetchBySlug", testSlug).Return(nil, errors.New("query does not exist"))
		repository := NewProblem(mockApi, mockScrapper)

		_, err := repository.GetByName(testName)

		assert.ErrorIs(t, internalErr.QuestionNotFound, err)
	})
}

func TestGetByUrl(group *testing.T) {
	group.Run("should return a valid Question using a valid url", func(tt *testing.T) {
		testQues, _ := testdata.ImportFromFile("two_sum.json")
		testQuery := &model.QuestionQuery{
			Question: testQues,
		}
		data := []string{
			"https://www.leetcode.com/problems/two-sum/",
			"https://www.leetcode.com/problems/two-sum",
			"https://leetcode.com/problems/two-sum/",
			"www.leetcode.com/problems/two-sum",
			"leetcode.com/problems/two-sum",
		}
		for _, testUrl := range data {
			tt.Run(testUrl, func(t *testing.T) {
				mockApi := &datasource.GraphQLApi{}
				mockScrapper := &datasource.WebScrapper{}
				mockApi.On("FetchBySlug", testQues.TitleSlug).Return(testQuery, nil)
				repository := NewProblem(mockApi, mockScrapper)

				actualQues, err := repository.GetByUrl(testUrl)

				assert.Nil(t, err)
				assert.Equal(t, &testQues, actualQues)
			})
		}
	})

	group.Run("error verification", func(tt *testing.T) {
		tt.Run("should return InvalidUrl error when domain name is not leetcode.com", func(t *testing.T) {
			testUrl := "https://codeforces.com/problemset/problem/1600/J"
			mockApi := &datasource.GraphQLApi{}
			mockScrapper := &datasource.WebScrapper{}
			repository := NewProblem(mockApi, mockScrapper)

			ques, err := repository.GetByUrl(testUrl)

			assert.Nil(t, ques)
			assert.ErrorIs(t, err, internalErr.InvalidURL)
		})

		tt.Run("should return LoginRequired error when problem is from a leetcode curated playlist", func(t *testing.T) {
			testUrl := "https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/631/"
			mockApi := &datasource.GraphQLApi{}
			mockScrapper := &datasource.WebScrapper{}
			repository := NewProblem(mockApi, mockScrapper)

			ques, err := repository.GetByUrl(testUrl)

			assert.Nil(t, ques)
			assert.ErrorIs(t, err, internalErr.LoginRequired)
		})

		tt.Run("should return InvalidUrl error when 2nd subdomain of url is not a valid slug string", func(t *testing.T) {
			testUrl := "https://leetcode.com/problems/pow(x,n)/"
			mockApi := &datasource.GraphQLApi{}
			mockScrapper := &datasource.WebScrapper{}
			repository := NewProblem(mockApi, mockScrapper)

			ques, err := repository.GetByUrl(testUrl)

			assert.Nil(t, ques)
			assert.ErrorIs(t, err, internalErr.LoginRequired)
		})
	})

	group.Run("should return no Question (nil) in case of an error from api", func(t *testing.T) {
		testUrl := "https://leetcode.com/problems/two-summ/"
		testSlug := "two-summ"
		mockApi := &datasource.GraphQLApi{}
		mockScrapper := &datasource.WebScrapper{}
		mockApi.On("FetchBySlug", testSlug).Return(nil, errors.New("query does not exist"))
		repository := NewProblem(mockApi, mockScrapper)

		ques, err := repository.GetByUrl(testUrl)

		assert.Nil(t, ques)
		assert.NotNil(t, err)
	})
}

func TestGetByNumber(group *testing.T) {
	testQues, _ := testdata.ImportFromFile("two_sum.json")

	group.Run("should return valid Question using its id", func(t *testing.T) {
		testQuery := &model.QuestionListQuery{
			QuestionList: struct {
				TotalNum int               `json:"totalNum"`
				Data     []entity.Question `json:"data"`
			}{
				TotalNum: 1,
				Data:     []entity.Question{testQues},
			},
		}
		testNum := 1
		mockApi := &datasource.GraphQLApi{}
		mockScrapper := &datasource.WebScrapper{}
		mockApi.On("FetchByNumber", strconv.Itoa(testNum)).Return(testQuery, nil)
		repository := NewProblem(mockApi, mockScrapper)

		actualQues, err := repository.GetByNumber(testNum)

		assert.Nil(t, err)
		assert.Equal(t, &testQues, actualQues)
	})

	group.Run("error verification", func(tt *testing.T) {
		tt.Run("should return QuestionIdOutOfRange error when no question list is returned by api", func(t *testing.T) {
			testNum := 1
			testEmptyQueryList := &model.QuestionListQuery{}
			mockApi := &datasource.GraphQLApi{}
			mockScrapper := &datasource.WebScrapper{}
			mockApi.On("FetchByNumber", strconv.Itoa(testNum)).Return(testEmptyQueryList, nil)
			repository := NewProblem(mockApi, mockScrapper)

			ques, err := repository.GetByNumber(testNum)

			assert.Nil(t, ques)
			assert.ErrorIs(t, err, internalErr.QuestionIdOutOfRange)
		})

		tt.Run("should return QuestionIdOutOfRange error when entered id is out of range", func(t *testing.T) {
			data := []int{
				-1,
				0,
				2030,
				25352325,
			}
			testQuery := &model.QuestionListQuery{
				QuestionList: struct {
					TotalNum int               `json:"totalNum"`
					Data     []entity.Question `json:"data"`
				}{
					TotalNum: 1,
					Data: []entity.Question{
						{
							Id: "5",
						},
						testQues,
					},
				},
			}
			for _, testNum := range data {
				mockApi := &datasource.GraphQLApi{}
				mockScrapper := &datasource.WebScrapper{}
				mockApi.On("FetchByNumber", strconv.Itoa(testNum)).Return(testQuery, nil)
				repository := NewProblem(mockApi, mockScrapper)

				ques, err := repository.GetByNumber(testNum)

				assert.Nil(t, ques)
				assert.ErrorIs(t, err, internalErr.QuestionIdOutOfRange)
			}
		})
	})
}

func TestGetProblemOfTheDay(group *testing.T) {
	anyInt := mock.AnythingOfType("int")

	group.Run("should return valid valid daily challenge Question", func(t *testing.T) {
		q0, _ := testdata.ImportFromFile("3sum.json")
		q1, _ := testdata.ImportFromFile("reverse_integer.json")
		q2, _ := testdata.ImportFromFile("two_sum.json")
		testQuery := &model.DailyChallengesQuery{
			DailyCodingChallengeV2: struct {
				Challenges []struct {
					Question entity.Question `json:"question"`
				} `json:"challenges"`
			}{
				Challenges: []struct {
					Question entity.Question `json:"question"`
				}{
					{q0},
					{q1},
					{q2},
				},
			},
		}
		mockApi := &datasource.GraphQLApi{}
		mockScrapper := &datasource.WebScrapper{}
		mockApi.On("FetchDailyChallengesOfMonth", anyInt, anyInt).Return(testQuery, nil)
		repository := NewProblem(mockApi, mockScrapper)

		actualQues, err := repository.GetDailyChallenge()

		assert.Nil(t, err)
		assert.Equal(t, &q2, actualQues)
	})

	group.Run("should return NoDailyChallenge error there is no challenge found (possibly due to different timezones", func(t *testing.T) {
		testQuery := &model.DailyChallengesQuery{
			DailyCodingChallengeV2: struct {
				Challenges []struct {
					Question entity.Question `json:"question"`
				} `json:"challenges"`
			}{
				Challenges: []struct {
					Question entity.Question `json:"question"`
				}{},
			},
		}
		mockApi := &datasource.GraphQLApi{}
		mockScrapper := &datasource.WebScrapper{}
		mockApi.On("FetchDailyChallengesOfMonth", anyInt, anyInt).Return(testQuery, nil)
		repository := NewProblem(mockApi, mockScrapper)

		actualQues, err := repository.GetDailyChallenge()

		assert.Nil(t, actualQues)
		assert.ErrorIs(t, err, internalErr.NoDailyChallenge)
	})
}
