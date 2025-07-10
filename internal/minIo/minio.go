package minIo

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

func UploadToContainer(ctx context.Context, minioClient *minio.Client, bucketName string, objectName string, file io.Reader, header *multipart.FileHeader) error {

	_, err := minioClient.PutObject(
		ctx,
		bucketName,
		objectName,
		file,
		header.Size,
		minio.PutObjectOptions{
			ContentType: header.Header.Get("Content-Type"),
		},
	)

	if err != nil {
		return err
	}
	return nil
}

func FetchVideos(ctx context.Context, minioClient *minio.Client, bucketName string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	objectCh := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
		Prefix:    "video",
		Recursive: true,
	})

	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return object.Err
		}
	}

	fmt.Println(objectCh)
	return nil
}
