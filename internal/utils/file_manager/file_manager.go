package filemanager

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// UploadFile funksýasy HTTP isleginden faýly alýar we serwerde saklap, ýoly yzyna gaýtarýar
func UploadFile(ctx *gin.Context, fileName, publicFolder string) (string, error) {
	file, err := ctx.FormFile(fileName)
	if err != nil {
		return "", fmt.Errorf("faýly almak başartmady: %v", err)
	}

	// Katalogyň bardygyny barlaýar, ýok bolsa, döredýär
	if _, err := os.Stat(publicFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(publicFolder, os.ModePerm); err != nil {
			return "", fmt.Errorf("katalogy döretmek başartmady: %v", err)
		}
	}

	// Faýly saklamak üçin doly ýol döreýär
	destPath := filepath.Join(publicFolder, file.Filename)

	// Faýly göçürip saklaýar
	if err := ctx.SaveUploadedFile(file, destPath); err != nil {
		return "", fmt.Errorf("faýly saklamak başartmady: %v", err)
	}

	// Başarýan ýagdaýynda, faýlyň doly ýoluny yzyna gaýtarýar
	return destPath, nil
}

// DeleteFile funksýasy berlen ýoldaky faýly pozýar
func DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("faýly pozmak başartmady: %v", err)
	}
	return nil
}
