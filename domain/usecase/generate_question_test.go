package usecase

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vldcreation/leet-scrape/v2/domain/entity"
	"github.com/vldcreation/leet-scrape/v2/internal/mock/service"
	"github.com/vldcreation/leet-scrape/v2/internal/testdata"
)

func TestGenerateQuestionFileUseCase(group *testing.T) {
	testPath := filepath.Join("test", "path")
	testQues, _ := testdata.ImportFromFile("convert_sorted_array_to_binary_search_tree.json")
	anyString := mock.AnythingOfType("string")
	group.Run("should write file with correct content and name at the correct path", func(t *testing.T) {
		mockFileWriter := &service.FileWriter{}
		mockImageDownloader := &service.ImageDownloader{}
		mockFileWriter.On("WriteDataToFile", testQues.TitleSlug+".html", testPath, anyString).Return(nil)
		mockImageDownloader.On("DownloadImageFromUrl", anyString, anyString, anyString).Return(nil)
		uc := NewGenerateQuestionFile(&testQues, testPath, mockFileWriter, mockImageDownloader)
		err := uc.Execute()
		assert.Nil(t, err)
		mockFileWriter.AssertNumberOfCalls(t, "WriteDataToFile", 1)
	})

	group.Run("should download images from the current url at the correct path", func(t *testing.T) {
		mockFileWriter := &service.FileWriter{}
		mockImageDownloader := &service.ImageDownloader{}
		mockFileWriter.On("WriteDataToFile", anyString, anyString, anyString).Return(nil)
		imgData := []struct {
			name string
			url  string
		}{
			{"btree.jpg", "https://assets.leetcode.com/uploads/2021/02/18/btree.jpg"},
			{"btree1.jpg", "https://assets.leetcode.com/uploads/2021/02/18/btree1.jpg"},
			{"btree2.jpg", "https://assets.leetcode.com/uploads/2021/02/18/btree2.jpg"},
		}
		mediaPath := filepath.Join(testPath, "media")
		for _, data := range imgData {
			name := testQues.TitleSlug + "_" + data.name
			mockImageDownloader.On("DownloadImageFromUrl", name, mediaPath, data.url).Return(nil)
		}
		uc := NewGenerateQuestionFile(&testQues, testPath, mockFileWriter, mockImageDownloader)
		err := uc.Execute()
		assert.Nil(t, err)
		mockFileWriter.AssertNumberOfCalls(t, "WriteDataToFile", 1)
		mockImageDownloader.AssertNumberOfCalls(t, "DownloadImageFromUrl", 3)
	})

	group.Run("should replace image url with local path of downloaded image before writing to html file", func(t *testing.T) {
		mockFileWriter := &service.FileWriter{}
		mockImageDownloader := &service.ImageDownloader{}
		tQues := &entity.Question{
			TitleSlug: "xyz",
			Content: "<img src=\"https://assets.leetcode.com/uploads/2021/02/18/btree1.jpg\"/>" +
				"<img alt=\"\" src=\"https://assets.leetcode.com/uploads/2021/02/18/btree1.jpg\" style=\"width: 302px; height: 222px;\" />",
		}
		mediaPath := filepath.Join(testPath, "media")
		imgName := tQues.TitleSlug + "_" + "btree1.jpg"
		imgPath := filepath.Join(mediaPath, imgName)
		expectedContent := "<img src=\"" + imgPath + "\"/><img alt=\"\" src=\"" + imgPath + "\" style=\"width: 302px; height: 222px;\" />"
		mockFileWriter.On("WriteDataToFile", tQues.TitleSlug+".html", testPath, expectedContent).Return(nil)
		mockImageDownloader.On("DownloadImageFromUrl", anyString, anyString, anyString).Return(nil)
		uc := NewGenerateQuestionFile(tQues, testPath, mockFileWriter, mockImageDownloader)
		err := uc.Execute()
		assert.Nil(t, err)
		mockFileWriter.AssertNumberOfCalls(t, "WriteDataToFile", 1)
		mockImageDownloader.AssertNumberOfCalls(t, "DownloadImageFromUrl", 2)
		mockImageDownloader.AssertExpectations(t)
	})
}
