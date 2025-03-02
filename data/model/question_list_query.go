package model

import "github.com/vldcreation/leet-scrape/v2/domain/entity"

type QuestionListQuery struct {
	QuestionList struct {
		TotalNum int               `json:"totalNum"`
		Data     []entity.Question `json:"data"`
	} `json:"questionList"`
}
