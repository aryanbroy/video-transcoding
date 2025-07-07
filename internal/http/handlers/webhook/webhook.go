package webhook

import (
	"log"
	"net/http"

	"github.com/aryanbroy/video-transcoding/internal/utils/response"
)

func WebhookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Webhook triggered")
		log.Printf("Webhook received: %s %s\nHeaders: %v\n", r.Method, r.URL.Path, r.Header)
		response.WriteJson(w, 200, "webhook triggered")
	}
}
