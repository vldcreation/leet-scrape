package util

import (
	"path/filepath"
	"slices"
)

func GetPath(basePath string, joinPaths ...string) string {
	joinPaths = slices.Compact(joinPaths)
	joinPaths = append([]string{basePath}, joinPaths...)
	return filepath.Join(joinPaths...)
}
