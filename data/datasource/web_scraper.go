package datasource

import (
	"github.com/gocolly/colly/v2"
	"github.com/vldcreation/leet-scrape/v2/internal/errors"
)

type WebScrapper interface {
	ScrapeNameOfDailyChallenge() (string, error)
}

type WebScrapperImpl struct {
	collector *colly.Collector
}

func NewWebScrapperImpl(c *colly.Collector) *WebScrapperImpl {
	return &WebScrapperImpl{collector: c}
}

func (ws *WebScrapperImpl) ScrapeNameOfDailyChallenge() (string, error) {
	return "", errors.NotImplemented
}
