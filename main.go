package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aryanbroy/video-transcoding/internal/http/handlers/videos"
	"github.com/aryanbroy/video-transcoding/internal/http/handlers/webhook"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	// cmd := exec.Command("ls")
	// output, err := cmd.Output()
	// if err != nil {
	// 	log.Fatalln("Error executing the command", err)
	// }
	// fmt.Println("Command ran successfully")
	// fmt.Println(string(output))
	ctx := context.Background()

	router := http.NewServeMux()
	addr := ":3000"

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
	}

	router.HandleFunc("POST /videos", videos.UploadToMinIO(ctx, minioClient))
	router.HandleFunc("POST /webhook", webhook.WebhookHandler(ctx, minioClient))
	router.HandleFunc("POST /webhook/", webhook.WebhookHandler(ctx, minioClient))

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Println("Server started at port", addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln("Error starting the server:\n", err)
	}
}
