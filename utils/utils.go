package utils

import (
	"path/filepath"
	"strings"
)

func BuildTranscriptionPath(videoName string) string {
	baseName := strings.TrimSuffix(videoName, filepath.Ext(videoName))
	return filepath.Join("transcriptions", baseName+".txt")
}
