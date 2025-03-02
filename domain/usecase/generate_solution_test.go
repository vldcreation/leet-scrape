package usecase

import (
	"testing"

	"github.com/ISKalsi/leet-scrape/v2/domain/entity"
	"github.com/ISKalsi/leet-scrape/v2/internal/errors"
	"github.com/ISKalsi/leet-scrape/v2/internal/mock/service"
	"github.com/ISKalsi/leet-scrape/v2/internal/testdata"
	"github.com/stretchr/testify/assert"
)

func TestGenerateSolutionFileUseCase(group *testing.T) {
	testQuestion := &testdata.QuestionWithImg
	testPath := "testFiles/cpp"
	testFileName := "two-sum"
	testBoilerplate := "// this is a test comment\n\n"
	testLang := "C++"
	testTemplate := "easy"

	group.Run("should return valid object from constructor", func(t *testing.T) {
		mockWriter := &service.FileWriter{}
		uc := NewGenerateSolutionFile(testQuestion, mockWriter, testFileName, testPath, testBoilerplate, testLang, testTemplate)
		assert.Equal(t, testQuestion, uc.question)
		assert.Equal(t, testPath, uc.path)
		assert.Equal(t, testBoilerplate, uc.boilerplate)
		assert.Equal(t, testLang, uc.language)
		assert.Equal(t, testFileName, uc.fileName)
	})

	group.Run("should prepend boilerplate code to the given code snippet (with newline chars fixed) before writing it to file", func(t *testing.T) {
		mw := &service.FileWriter{}
		q := &entity.Question{
			TitleSlug: "two-sum",
			Content:   "<sample content>",
			CodeSnippets: []entity.CodeSnippet{
				{
					"C++",
					"cpp",
					"<sample\\r\\n code\\n>",
				},
			},
		}
		tb := "// sample comment \\n"
		expectedData := "// sample comment \n<sample\r\n code\n>"
		mw.On("WriteDataToFile", testFileName+".cpp", testPath, expectedData).Return(nil)
		uc := NewGenerateSolutionFile(q, mw, testFileName, testPath, tb, testLang, testTemplate)
		err := uc.Execute()
		assert.Nil(t, err)
	})

	group.Run("should write the file with correct filename", func(t *testing.T) {
		mw := &service.FileWriter{}
		q := &entity.Question{
			CodeSnippets: []entity.CodeSnippet{
				{
					testLang,
					"cpp",
					"",
				},
			},
		}
		mw.On("WriteDataToFile", testFileName+".cpp", testPath, "").Return(nil)
		uc := NewGenerateSolutionFile(q, mw, testFileName, testPath, "", testLang, testTemplate)
		err := uc.Execute()
		assert.Nil(t, err)
	})

	group.Run("error verification", func(tt *testing.T) {
		tt.Run("should return NoCodeSnippetsFound error when there are no code snippets in the given question", func(t *testing.T) {
			mw := &service.FileWriter{}
			q := &entity.Question{
				TitleSlug:    "two-sum",
				Content:      "<question description>",
				CodeSnippets: []entity.CodeSnippet{},
			}
			uc := NewGenerateSolutionFile(q, mw, testFileName, testPath, testBoilerplate, testLang, testTemplate)
			err := uc.Execute()
			assert.ErrorIs(t, err, errors.NoCodeSnippetsFound)
		})

		tt.Run("should return ExtensionNotFound error when there is no file extension corresponding to the given language", func(t *testing.T) {
			mw := &service.FileWriter{}
			q := &entity.Question{
				TitleSlug: "two-sum",
				Content:   "<question description>",
				CodeSnippets: []entity.CodeSnippet{
					{
						"D",
						"d",
						"<sample code>",
					},
				},
			}
			uc := NewGenerateSolutionFile(q, mw, testFileName, testPath, testBoilerplate, q.CodeSnippets[0].Lang, testTemplate)
			err := uc.Execute()
			assert.ErrorIs(t, err, errors.ExtensionNotFound)
		})

		tt.Run("should return SnippetNotFoundInGivenLang error when there is no code snippets corresponding to the given language", func(t *testing.T) {
			mw := &service.FileWriter{}
			q := &entity.Question{
				TitleSlug: "two-sum",
				Content:   "<question description>",
				CodeSnippets: []entity.CodeSnippet{
					{
						"D",
						"d",
						"<sample code>",
					},
				},
			}
			uc := NewGenerateSolutionFile(q, mw, testFileName, testPath, testBoilerplate, testLang, testTemplate)
			err := uc.Execute()
			assert.ErrorIs(t, err, errors.SnippetNotFoundInGivenLang)
		})
	})
}
