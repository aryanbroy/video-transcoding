package misc

import "github.com/google/uuid"

func GenerateVideoId() string {
	return uuid.New().String()
}
