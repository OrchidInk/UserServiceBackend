package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"orchid.admin.service/utils"
	aws "orchid.admin.service/utils/file"
)

func (hd *Handlers) UploadFile(ctx *fiber.Ctx) error {
	rqFile, err := ctx.FormFile("file")
	if err != nil {
		slog.Error("Unable to read file from request", slog.Any("err", err))
		return utils.ErrMsg(err)
	}

	s3, err := aws.NewAwsBucket(hd.c)
	if err != nil {
		slog.Error("Unable to create AWS instance", slog.Any("err", err))
		return utils.ErrMsg(err)
	}

	fileObj, err := rqFile.Open()
	if err != nil {
		return utils.ErrMsg(err)
	}
	defer fileObj.Close()

	folderName := ctx.FormValue("foldername")
	guuid := uuid.NewString()
	newFileName := folderName + "/" + guuid + "$" + rqFile.Filename

	if err := s3.UploadFile(newFileName, fileObj); err != nil {
		slog.Error("Unable to upload file to S3", slog.Any("err", err))
		return utils.ErrMsg(err)
	}

	presignedURL, err := s3.GetPresignedObject(nil, nil, newFileName, 3600)
	if err != nil {
		return utils.ErrMsg(err)
	}

	return ctx.JSON(fiber.Map{"PresignedURL": presignedURL, "Path": newFileName})
}
