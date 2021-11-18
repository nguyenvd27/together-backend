package usecases

import (
	"context"
	"fmt"
	"mime/multipart"
	"together-backend/pkg"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

const MAX_UPLOAD_SIZE = 10485760 // 10Mb

type ERROR_CODE int

func EventImageUpload(files []*multipart.FileHeader) ([]string, error) {
	for _, fileHeader := range files {
		if fileHeader.Size > MAX_UPLOAD_SIZE {
			return nil, fmt.Errorf("the uploaded file is too big. Please choose an file that's less than 1MB in size")
		}
	}

	var cld, err = cloudinary.New()
	if err != nil {
		return nil, fmt.Errorf("failed to intialize Cloudinary")
	}

	var ctx = context.Background()
	var imagesSlice []string

	if len(files) > 5 {
		return nil, fmt.Errorf("the number of files uploaded is too much. Upload up to 5 files")
	}
	for i, _ := range files {
		fileName := files[i].Filename
		file, err := files[i].Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file")
		}
		defer file.Close()

		uploadResult, err := cld.Upload.Upload(
			ctx,
			file,
			uploader.UploadParams{
				PublicID: "events/" + fileName + "_" + pkg.RandomID(6),
			})
		if err != nil {
			return nil, fmt.Errorf("failed to upload file")
		}
		imagesSlice = append(imagesSlice, uploadResult.SecureURL)
	}

	return imagesSlice, nil
}
