package handlers

import (
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"orchid.admin.service/utils"
	aws "orchid.admin.service/utils/file"
)

func (hd *Handlers) UploadFile(ctx *fiber.Ctx) error {
	rqFile, err := ctx.FormFile("file")
	if err != nil {
		slog.Error("unable to read file from request", slog.Any("err", err))
		return utils.ErrMsg(err)
	}

	// Create AWS S3 Client
	s3, err := aws.NewAwsBucket(hd.c)
	if err != nil {
		slog.Error("unable to create aws instance", slog.Any("err", err))
		return utils.ErrMsg(err)
	}

	// Open the uploaded file
	localFile, err := rqFile.Open()
	if err != nil {
		slog.Error("Unable to open request file", slog.Any("err", err))
		return utils.ErrMsg(err)
	}
	defer localFile.Close()

	folderName := ctx.FormValue("foldername")
	validFolderNames := []string{
		"user-pic",
		"product-en",
		"product-mn",
		"banner-pic",
	}
	found := false
	for _, validName := range validFolderNames {
		if folderName == validName {
			found = true
			break
		}
	}

	if !found {
		return utils.CustomErrMsg("Folder name is incorrect")
	}

	guuid := uuid.NewString()
	newFileName := folderName + "/" + guuid + "$" + rqFile.Filename

	var contentDisposition, fileType string
	switch folderName {
	case "images":
		contentDisposition = "inline"
		fileType = "image/png"
	case "music":
		contentDisposition = "inline"
		fileType = "audio/mpeg"
	case "videos":
		contentDisposition = "inline"
		fileType = "video/mp4"
	default:
		contentDisposition = "inline"
		fileType = "application/octet-stream"
	}

	// Upload file to S3
	if err := s3.UploadFile(newFileName, localFile); err != nil {
		slog.Error("Unable to upload file to S3", slog.Any("err", err))
		_ = s3.DeleteObject(newFileName)
		return utils.ErrMsg(err)
	}

	presignedURL, err := s3.GetPresignedObject(&contentDisposition, &fileType, newFileName, 3600) // Assuming a 1-hour expiration
	if err != nil {
		slog.Error("Unable to generate presigned URL", slog.Any("err", err))
		return utils.ErrMsg(err)
	}

	return ctx.JSON(fiber.Map{"PresignedURL": presignedURL, "Path": newFileName, "ContentDisposition": contentDisposition})
}

func (hd *Handlers) GetFileURL(ctx *fiber.Ctx) error {
	filePath := ctx.Query("filepath")

	if filePath == "" {
		return utils.CustomErrMsg("Filepath query parameter is required")
	}

	s3, err := aws.NewAwsBucket(hd.c)
	if err != nil {
		return utils.ErrMsg(err)
	}

	// Determine MIME type based on file extension
	var contentDisposition, fileType string
	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".jpg", ".jpeg":
		contentDisposition = "inline"
		fileType = "image/jpeg"
	case ".png":
		contentDisposition = "inline"
		fileType = "image/png"
	case ".gif":
		contentDisposition = "inline"
		fileType = "image/gif"
	case ".mp3":
		contentDisposition = "inline"
		fileType = "audio/mpeg"
	case ".wav":
		contentDisposition = "inline"
		fileType = "audio/wav"
	case ".mp4":
		contentDisposition = "inline"
		fileType = "video/mp4"
	case ".avi":
		contentDisposition = "inline"
		fileType = "video/x-msvideo"
	default:
		contentDisposition = "inline"
		fileType = "application/octet-stream" // default binary type
	}

	// Generate presigned URL
	presignedURL, err := s3.GetPresignedObject(&contentDisposition, &fileType, filePath, 3600) // Assuming a 1-hour expiration
	if err != nil {
		return utils.ErrMsg(err)
	}

	return ctx.JSON(fiber.Map{"PresignedURL": presignedURL})
}
