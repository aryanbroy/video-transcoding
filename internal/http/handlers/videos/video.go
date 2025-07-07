package videos

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aryanbroy/video-transcoding/internal/utils/misc"
	"github.com/aryanbroy/video-transcoding/internal/utils/response"
	"github.com/minio/minio-go/v7"
)

func UploadToMinIO(ctx context.Context, minioClient *minio.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseMultipartForm(32 << 20); err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err, http.StatusBadRequest))
			return
		}

		file, header, err := r.FormFile("video")
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err, http.StatusBadRequest))
			return
		}
		defer file.Close()

		objectName := fmt.Sprintf("%s", misc.GenerateVideoId())
		_, err = minioClient.PutObject(
			ctx,
			"video-transcoder",
			objectName,
			file,
			header.Size,
			minio.PutObjectOptions{
				ContentType: header.Header.Get("Content-Type"),
			},
		)

		if err != nil {
			log.Printf("MinIO upload failed: %v", err)
			response.WriteJson(w, http.StatusInternalServerError,
				response.GeneralError(err, http.StatusInternalServerError))
			return
		}

		response.WriteJson(w, http.StatusCreated,
			response.CustomResponse("File uploaded to MinIO successfully", http.StatusCreated))
	}
}
