package web

import (
	"fmt"
	"os"
)

const toggleKey string = "uuid"

func GetListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
