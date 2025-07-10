package ffmpeg

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func ScaleVideo(fileName string, outputPath string) error {

	log.Println("Scaling video to certain resolution...")

	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("Error creating output directory: %w\n", err)
	}

	err := ffmpeg_go.Input(fileName).Output(outputPath, ffmpeg_go.KwArgs{"s": "854x480", "c:v": "libx264", "preset": "fast"}).OverWriteOutput().ErrorToStdOut().Run()

	if err != nil {
		return fmt.Errorf("Error running ffmpeg comamnd: %w\n", err)
	}
	log.Println("Scaling completed successfully!")

	return nil
}
