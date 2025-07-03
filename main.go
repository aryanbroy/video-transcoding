package main

import (
	"log"
	"net/http"

	"github.com/aryanbroy/video-transcoding/internal/http/handlers/videos"
)

func main() {
	// cmd := exec.Command("ls")
	// output, err := cmd.Output()
	// if err != nil {
	// 	log.Fatalln("Error executing the command", err)
	// }
	// fmt.Println("Command ran successfully")
	// fmt.Println(string(output))

	router := http.NewServeMux()
	addr := ":3000"

	router.HandleFunc("POST /videos", videos.UploadToMinIO())

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Println("Server started at port", addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln("Error starting the server:\n", err)
	}

}
