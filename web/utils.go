package web

import (
	"fmt"
	"os"
)

const key string = "status"
const toggleKey string = "uuid"

const yes string = "oui"
const no string = "non"

func GetListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
