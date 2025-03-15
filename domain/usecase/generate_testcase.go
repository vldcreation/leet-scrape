package usecase

import (
	"fmt"

	"github.com/vldcreation/leet-scrape/v2/domain/entity"
	"github.com/vldcreation/leet-scrape/v2/domain/service"
)

type GenerateTestCaseFile struct {
	writer   service.FileWriter
	question *entity.Question
	testcase []entity.TestCase
	fileName string
	path     string
}

func NewGenerateTestCaseFile(writer service.FileWriter, question *entity.Question, testcase []entity.TestCase, fileName, path string) *GenerateTestCaseFile {
	if fileName == "" {
		fileName = question.TitleSlug
	}

	return &GenerateTestCaseFile{
		writer:   writer,
		question: question,
		testcase: testcase,
		fileName: fileName,
		path:     path,
	}
}

func (uc *GenerateTestCaseFile) Execute() error {
	for i, tc := range uc.testcase {
		inputFileName := fmt.Sprintf("%s-%d.in", uc.fileName, i+1)
		outputFileName := fmt.Sprintf("%s-%d.out", uc.fileName, i+1)
		explanationFileName := fmt.Sprintf("%s-%d.explanation", uc.fileName, i+1)

		if err := uc.writer.WriteDataToFile(inputFileName, uc.path, tc.Input); err != nil {
			return err
		}

		if err := uc.writer.WriteDataToFile(outputFileName, uc.path, tc.Output); err != nil {
			return err
		}

		if tc.Explanation != "" {
			if err := uc.writer.WriteDataToFile(explanationFileName, uc.path, tc.Explanation); err != nil {
				return err
			}
		}
	}

	return nil
}
