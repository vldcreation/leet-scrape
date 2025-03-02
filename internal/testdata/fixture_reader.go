package testdata

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"

	"github.com/vldcreation/leet-scrape/v2/data/model"
	"github.com/vldcreation/leet-scrape/v2/domain/entity"
)

func ImportFromFile(fixtureName string) (entity.Question, error) {
	_, pathToFixtures, _, _ := runtime.Caller(0)
	_ = os.Chdir(filepath.Dir(pathToFixtures))
	file, err := os.ReadFile(fixtureName)
	if err != nil {
		return entity.Question{}, err
	}
	var q map[string]model.QuestionQuery
	err = json.Unmarshal(file, &q)
	if err != nil {
		return entity.Question{}, err
	}
	return q["data"].Question, nil
}
