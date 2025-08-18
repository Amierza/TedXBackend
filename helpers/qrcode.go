package helpers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/skip2/go-qrcode"
)

func GenerateQRCodeFile(content string, filename string) (string, error) {
	localPath := filepath.Join("assets", "qrcodes", filename)

	err := qrcode.WriteFile(content, qrcode.Medium, 256, localPath)
	if err != nil {
		return "", err
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8000"
	}

	publicURL := fmt.Sprintf("%s/assets/qrcodes/%s", baseURL, filename)
	log.Println(publicURL)
	return publicURL, nil
}
