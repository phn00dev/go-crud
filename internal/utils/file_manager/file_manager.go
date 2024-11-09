package filemanager

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// UploadFile funksýasy HTTP isleginden faýly alýar we serwerde saklap, ýoly yzyna gaýtarýar
func UploadFile(ctx *gin.Context, fileName, publicFolder string) (string, error) {
	file, err := ctx.FormFile(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve the file. Please try again.: %v", err)
	}

	// Faýl formatyny barlaýar
	if !isValidImageFormat(file) {
		return "", fmt.Errorf("invalid file format. Only 'png', 'jpg', and 'jpeg' files are allowed")
	}

	if _, err := os.Stat(publicFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(publicFolder, os.ModePerm); err != nil {
			return "", fmt.Errorf("katalogy döretmek başartmady: %v", err)
		}
	}

	newFileName := generateRandomFileName() + filepath.Ext(file.Filename)

	destPath := filepath.Join(publicFolder, newFileName)

	if err := ctx.SaveUploadedFile(file, destPath); err != nil {
		return "", fmt.Errorf("failed to save the file. Please try again. %v", err)
	}

	return destPath, nil
}

func DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("faýly pozmak başartmady: %v", err)
	}
	return nil
}

func generateRandomFileName() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

// Faýl formatyny validasiýa etmek
func isValidImageFormat(file *multipart.FileHeader) bool {
	allowedExtensions := []string{".png", ".jpg", ".jpeg"}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			return true
		}
	}
	return false
}
