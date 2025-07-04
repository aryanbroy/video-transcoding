package videos

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aryanbroy/video-transcoding/internal/minio"
	"github.com/aryanbroy/video-transcoding/internal/utils/files"
	"github.com/aryanbroy/video-transcoding/internal/utils/response"
)

func UploadToMinIO() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseMultipartForm(32 << 20)

		file, header, err := r.FormFile("video")
		if err != nil {
			log.Fatalln("Error fetching video from body! \n", err)
		}
		defer file.Close()

		workingDir, err := os.Getwd()
		if err != nil {
			log.Fatalln("Error getting working dir\n", err)
		}

		uploadPath := filepath.Join(workingDir, "uploads", "videos", header.Filename)
		exists, err := files.FileExists(uploadPath)
		if exists {
			response.WriteJson(w, http.StatusAlreadyReported, response.GeneralError(fmt.Errorf("File already exists in that path!"), http.StatusAlreadyReported))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err, http.StatusBadRequest))
			return
		}

		dst, err := os.Create(uploadPath)
		if err != nil {
			log.Fatalln("Error creating local file!\n", err)
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err, http.StatusBadRequest))
			return
		}

		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			log.Fatalln("Error copying file to destination\n", err)
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err, http.StatusBadRequest))
			return
		}

		ok, err := minio.UploadToContainer(uploadPath, "video-transcoder", "video")
		if err != nil || !ok {
			log.Fatalln("Error uploading to container")
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err, http.StatusBadRequest))
			return
		}

		err = files.DeleteFile(uploadPath)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err, http.StatusBadRequest))
			return
		}

		response.WriteJson(w, http.StatusCreated, response.CustomResponse("File uploaded to container successfully", http.StatusCreated))
	}
}
