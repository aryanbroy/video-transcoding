package worker

import (
	"fmt"
	"time"
)

func ProcessVideo(videoName string) {
	fmt.Println("Video is being processed...")
	time.Sleep(5 * time.Second)
	fmt.Printf("Video %v has been processed", videoName)
}
