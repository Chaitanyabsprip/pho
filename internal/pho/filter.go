package pho

import (
	"path"
	"strings"
)

func hiddenFileFilter(filepath string) bool {
	basename := path.Base(filepath)
	return basename[0] == '.'
}

func hiddenPathFilter(filepath string) bool {
	return strings.Contains(filepath, "/.")
}
