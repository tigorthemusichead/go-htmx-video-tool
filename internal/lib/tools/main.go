package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"video-tool/internal/lib/ffmpeg"
)

const op = "/internal/lib/tools"

func MergePhotosAndAudios(inputPaths []string, outputPath string) error {
	var inputs []ffmpeg.Input
	for _, filePath := range inputPaths {
		if _, err := os.Stat(filePath); err != nil {
			return fmt.Errorf("%s MergePhotosAndAudios: file with path %s doesn't exist", op, filePath)
		} else {
			input := ffmpeg.Input{
				FilePath: filePath,
				Params:   []string{},
			}
			inputs = append(inputs, input)
		}
	}
	output := ffmpeg.Output{
		FilePath: filepath.Join(outputPath),
		Params:   []string{},
	}
	f, _ := ffmpeg.New(inputs, output, "", "")
	return f.Start()
}
