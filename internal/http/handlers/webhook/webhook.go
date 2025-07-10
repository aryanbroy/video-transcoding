package webhook

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aryanbroy/video-transcoding/internal/utils/response"
	"github.com/aryanbroy/video-transcoding/internal/worker"
	"github.com/minio/minio-go/v7"
)

type MinIOEvent struct {
	EventName string `json:"EventName"`
	Key       string `json:"Key"`
	Records   []struct {
		EventVersion string `json:"eventVersion"`
		S3           struct {
			Bucket struct {
				Name string `json:"name"`
			} `json:"bucket"`
			Object struct {
				Key  string `json:"key"`
				Size int64  `json:"size"`
				ETag string `json:"eTag"`
			} `json:"object"`
		} `json:"s3"`
	} `json:"Records"`
}

func WebhookHandler(ctx context.Context, minioClient *minio.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Webhook triggered")
		var event MinIOEvent

		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			log.Println("Unable to decode body to variable")
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err, http.StatusBadRequest))
			return
		}

		if len(event.Records) == 0 {
			log.Println("No records present")
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err, http.StatusBadRequest))
			return
		}

		videoName := event.Records[0].S3.Object.Key
		log.Printf("Starting processing for: %s", videoName)

		go worker.ProcessVideo(ctx, minioClient, "video-transcoder", videoName)

		response.WriteJson(w, 200, "Webhook triggered")
		fmt.Println("Http response finishd")
	}
}
