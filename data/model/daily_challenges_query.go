package model

import "github.com/vldcreation/leet-scrape/v2/domain/entity"

type DailyChallengesQuery struct {
	DailyCodingChallengeV2 struct {
		Challenges []struct {
			Question entity.Question `json:"question"`
		} `json:"challenges"`
	} `json:"dailyCodingChallengeV2"`
}
