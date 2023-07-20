package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type UploaderService struct {
	CloudName    string
	ApiKey       string
	ApiSecret    string
	UploadFolder string
}

func NewUploaderService(cloudName string, apiKey string, apiSecret string, uploadFolder string) *UploaderService {
	return &UploaderService{
		CloudName:    cloudName,
		ApiKey:       apiKey,
		ApiSecret:    apiSecret,
		UploadFolder: uploadFolder,
	}
}

func (s *UploaderService) UploadImage(ID int, input *multipart.FileHeader) (imageURL string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	cld, _ := cloudinary.NewFromParams(s.CloudName, s.ApiKey, s.ApiSecret)

	file, _ := input.Open()
	defer file.Close()

	imageName := s.generateFileName(input.Filename, ID)

	// upload file
	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: s.UploadFolder, PublicID: imageName})
	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}

func (s *UploaderService) generateFileName(fileName string, ID int) string {
	// get file extension
	ext := filepath.Ext(fileName)
	// remove trailing
	fileTrimSpace := strings.TrimSpace(fileName)
	// remove extension
	fileNameNoExt := strings.TrimSuffix(fileTrimSpace, ext)
	// replace space to "-"
	fileNameNoSpace := strings.ReplaceAll(fileNameNoExt, " ", "-")
	// create image name (fileName-ID)
	fileName = fmt.Sprintf("%s-%d", fileNameNoSpace, ID)

	return strings.ToLower(fileName)
}
