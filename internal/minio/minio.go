package minio

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func UploadToContainer(filePath string, bucketName string, objectName string) (bool, error) {
	ctx := context.Background()
	endpoint := "localhost:9000"
	accessKeyId := "minioadmin"
	secrectAccessKey := "minioadmin"
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secrectAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Println("Error initializing minio client")
		return false, err
	}

	location := "us-east-1"

	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		log.Panicln("Error checking if bucket exists or not!")
		return false, err
	}

	if !exists {
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			log.Println("Error creating bucket", err)
			return false, err
		}
	}

	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: "video/mp4"})
	if err != nil {
		log.Println("Error storing video in container")
		return false, err
	}

	log.Printf("Successfully upload %v to the container. Size: %v\n", objectName, info.Size)
	return true, nil
}
