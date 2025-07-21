package services

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func ListVideos() ([]string, error) {
	path := "./videos"
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var videos []string
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".mp4" || filepath.Ext(f.Name()) == ".MOV" {
			videos = append(videos, f.Name())
		}
	}
	return videos, nil
}

func TranscribeVideo(name string) (string, error) {
	path := filepath.Join("videos", name)
	if _, err := os.Stat(path); err != nil {

		return "", fmt.Errorf("file not found")
	}

	cmd := exec.Command("python", "transcribe.py", path)
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error executing script:")
		fmt.Println("Error:", err)
		fmt.Println("Output:", string(output))
		return "", fmt.Errorf("error executing transcription")
	}

	fmt.Println("Transcription got successfully")
	fmt.Println(string(output))
	return string(output), nil
}
