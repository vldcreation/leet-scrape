package usecase

import (
	"path/filepath"
	"regexp"

	"github.com/vldcreation/leet-scrape/v2/domain/entity"
	"github.com/vldcreation/leet-scrape/v2/domain/service"
	"github.com/vldcreation/leet-scrape/v2/internal/util"
)

type GenerateQuestionFile struct {
	question  *entity.Question
	path      string
	writer    service.FileWriter
	imgGetter service.ImageDownloader
}

type imgInfo struct {
	name string
	url  string
}

func NewGenerateQuestionFile(q *entity.Question, path string, w service.FileWriter, ig service.ImageDownloader) *GenerateQuestionFile {
	return &GenerateQuestionFile{
		path:      path,
		question:  q,
		writer:    w,
		imgGetter: ig,
	}
}

func (uc *GenerateQuestionFile) Execute() error {
	uc.question.Content = util.FixNewlineAndTabs(uc.question.Content)
	imgData := uc.extractNameAndUrlFromImgTags()
	localImgPaths, err := uc.downloadImages(imgData)
	if err != nil {
		return err
	}
	content := uc.contentWithLocalImgPaths(localImgPaths)
	fileName := uc.question.TitleSlug + ".html"
	return uc.writer.WriteDataToFile(fileName, uc.path, content)
}

func (uc *GenerateQuestionFile) extractNameAndUrlFromImgTags() []imgInfo {
	imgTagRegex := regexp.MustCompile(`<img.*? src="(.*?)".*?>`)
	imgTags := imgTagRegex.FindAllStringSubmatch(uc.question.Content, -1)
	imgNameRegex := regexp.MustCompile(`([^/]+)(/$|$)`)
	var imgData []imgInfo
	for _, match := range imgTags {
		url := match[1]
		name := imgNameRegex.FindStringSubmatch(url)[1]
		imgData = append(imgData, imgInfo{name, url})
	}
	return imgData
}

func (uc *GenerateQuestionFile) downloadImages(imgData []imgInfo) ([]string, error) {
	var localImgPaths []string
	mediaPath := filepath.Join(uc.path, "media")
	for _, data := range imgData {
		imgName := uc.question.TitleSlug + "_" + data.name
		err := uc.imgGetter.DownloadImageFromUrl(imgName, mediaPath, data.url)
		imgPath := filepath.Join(mediaPath, imgName)
		localImgPaths = append(localImgPaths, imgPath)
		if err != nil {
			return nil, err
		}
	}
	return localImgPaths, nil
}

func (uc *GenerateQuestionFile) contentWithLocalImgPaths(imgPaths []string) string {
	i := 0
	imgTagRegex := regexp.MustCompile(`(<img.*? src=").*?(".*?>)`)
	content := imgTagRegex.ReplaceAllStringFunc(uc.question.Content, func(imgTag string) string {
		modifiedTag := imgTagRegex.ReplaceAllString(imgTag, "${1}"+imgPaths[i]+"$2")
		i += 1
		return modifiedTag
	})
	return content
}
