package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vldcreation/leet-scrape/v2/internal/errors"
	"github.com/vldcreation/leet-scrape/v2/internal/mock/repo"
	"github.com/vldcreation/leet-scrape/v2/internal/testdata"
)

func TestGetProblemUseCase(group *testing.T) {
	testName := "Two Sum"
	testUrl := "https://leetcode.com/problems/two-sum/"
	testNum := 45

	testQuestion := &testdata.QuestionWithImg

	group.Run("should return valid GetProblem use case", func(tt *testing.T) {
		tt.Run("by name", func(t *testing.T) {
			mockRepo := repo.Problem{}
			getProblem := NewGetProblemByName(&mockRepo, testName)
			assert.Equal(t, "", getProblem.url)
			assert.Equal(t, 0, getProblem.num)
			assert.Equal(t, testName, getProblem.name)
			assert.Equal(t, byName, getProblem.search)
			mockRepo.AssertNotCalled(t, "Execute")
		})

		tt.Run("by num", func(t *testing.T) {
			mockRepo := repo.Problem{}
			getProblem := NewGetProblemByUrl(&mockRepo, testUrl)
			assert.Equal(t, testUrl, getProblem.url)
			assert.Equal(t, 0, getProblem.num)
			assert.Equal(t, "", getProblem.name)
			assert.Equal(t, byUrl, getProblem.search)
			mockRepo.AssertNotCalled(t, "Execute")
		})

		tt.Run("by url", func(t *testing.T) {
			mockRepo := repo.Problem{}
			getProblem := NewGetProblemByNumber(&mockRepo, testNum)
			assert.Equal(t, "", getProblem.url)
			assert.Equal(t, testNum, getProblem.num)
			assert.Equal(t, "", getProblem.name)
			assert.Equal(t, byNumber, getProblem.search)
			mockRepo.AssertNotCalled(t, "Execute")
		})

		tt.Run("problem of day", func(t *testing.T) {
			mockRepo := repo.Problem{}
			getProblem := NewGetDailyChallenge(&mockRepo)
			assert.Equal(t, "", getProblem.url)
			assert.Equal(t, 0, getProblem.num)
			assert.Equal(t, "", getProblem.name)
			assert.Equal(t, problemOfDay, getProblem.search)
			mockRepo.AssertNotCalled(t, "Execute")
		})
	})

	group.Run("should call the correct repository method to fetch problem data", func(tt *testing.T) {
		tt.Run("by name", func(t *testing.T) {
			mockRepo := repo.Problem{}
			mockRepo.On("GetByName", testName).Return(testQuestion, nil)
			getProblem := NewGetProblemByName(&mockRepo, testName)
			actualQuestion, err := getProblem.Execute()
			assert.Nil(t, err)
			assert.Equal(t, testQuestion, actualQuestion)
			mockRepo.AssertCalled(t, "GetByName", testName)
			mockRepo.AssertExpectations(t)
		})

		tt.Run("by url", func(t *testing.T) {
			mockRepo := repo.Problem{}
			mockRepo.On("GetByUrl", testUrl).Return(&testdata.QuestionWithImg, nil)
			getProblem := NewGetProblemByUrl(&mockRepo, testUrl)
			actualQuestion, err := getProblem.Execute()
			assert.Nil(t, err)
			assert.Equal(t, testQuestion, actualQuestion)
			mockRepo.AssertCalled(t, "GetByUrl", testUrl)
			mockRepo.AssertExpectations(t)
		})

		tt.Run("by num", func(t *testing.T) {
			mockRepo := repo.Problem{}
			mockRepo.On("GetByNumber", testNum).Return(&testdata.QuestionWithImg, nil)
			getProblem := NewGetProblemByNumber(&mockRepo, testNum)
			actualQuestion, err := getProblem.Execute()
			assert.Nil(t, err)
			assert.Equal(t, testQuestion, actualQuestion)
			mockRepo.AssertCalled(t, "GetByNumber", testNum)
			mockRepo.AssertExpectations(t)
		})

		tt.Run("daily challenge", func(t *testing.T) {
			mockRepo := repo.Problem{}
			mockRepo.On("GetDailyChallenge").Return(&testdata.QuestionWithImg, nil)
			getProblem := NewGetDailyChallenge(&mockRepo)
			actualQuestion, err := getProblem.Execute()
			assert.Nil(t, err)
			assert.Equal(t, testQuestion, actualQuestion)
			mockRepo.AssertCalled(t, "GetDailyChallenge")
			mockRepo.AssertExpectations(t)
		})
	})

	group.Run("should return InvalidSearchMethod error when incorrect method is provided", func(t *testing.T) {
		mockRepo := repo.Problem{}
		getProblem := NewGetProblemByNumber(&mockRepo, testNum)
		getProblem.search = -1
		actualQuestion, err := getProblem.Execute()
		assert.Nil(t, actualQuestion)
		assert.Error(t, err)
		assert.Equal(t, errors.InvalidSearchMethod, err)
		mockRepo.AssertNotCalled(t, "GetByNumber", testNum)
	})
}
