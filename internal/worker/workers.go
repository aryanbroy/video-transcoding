package worker

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
)

func ProcessVideo(ctx context.Context, minioClient *minio.Client, bucketName string, objectName string) {
	fmt.Println("Video is being processed...")

	err := minioClient.FGetObject(ctx, bucketName, objectName, "./received/download.mp4", minio.GetObjectOptions{})
	if err != nil {
		log.Panicln("Error fetching video from container: \n", err)
	}

	fmt.Printf("Video %v has been processed", objectName)
}
