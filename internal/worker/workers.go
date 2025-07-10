package worker

import (
	"context"
	"fmt"
	"log"

	"github.com/aryanbroy/video-transcoding/internal/utils/ffmpeg"
	"github.com/aryanbroy/video-transcoding/internal/utils/files"
	"github.com/minio/minio-go/v7"
)

func ProcessVideo(ctx context.Context, minioClient *minio.Client, bucketName string, objectName string) {
	fmt.Println("Video is being processed...")

	receivedPath := fmt.Sprintf("./received/%v.mp4", objectName)

	err := minioClient.FGetObject(ctx, bucketName, objectName, receivedPath, minio.GetObjectOptions{})
	if err != nil {
		log.Panicln("Error fetching video from container: \n", err)
	}

	ok, err := files.FileExists(receivedPath)
	if err != nil {
		log.Panicln("Error checking if file exist in the local storage: \n", err)
	}

	if !ok {
		log.Panicf("%v: no such path exists\n", receivedPath)
	}

	outputPath := fmt.Sprintf("./received/resize_%v.mp4", objectName)

	err = ffmpeg.ScaleVideo(receivedPath, outputPath)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("Deleting the downloaded file...")
	if err = files.DeleteFile(receivedPath); err != nil {
		log.Panicln(err)
	}
	log.Println("Successfully deleted the downloaded file")

	fmt.Printf("Video %v has been processed", objectName)
}
