package main

import (
	"slices"
	"strings"

	"github.com/vldcreation/leet-scrape/v2/consts"
)

func validateTemplate(tmpl string) error {
	if tmpl == "" {
		return nil
	}

	if slices.Contains(consts.TEMPLATES, tmpl) {
		return nil
	}

	if !strings.HasSuffix(tmpl, ".tmpl") {
		tmpl = tmpl + ".tmpl"
	}

	return nil
}
