package usecases

import (
	"context"
	"fmt"
	"mime/multipart"
	"together-backend/internal/repositories"
	"together-backend/pkg"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

const MAX_UPLOAD_SIZE = 10485760       // 10Mb
const MAX_UPLOAD_AVATAR_SIZE = 5242880 // 5Mb

type UploadUseCase interface {
	EventImageUpload(files []*multipart.FileHeader) ([]string, error)
	UploadAvatar(file *multipart.FileHeader) (string, error)
}

type uploadUsecase struct {
	imageRepo repositories.ImageRepo
}

func NewUploadUsecase(imageRepo repositories.ImageRepo) UploadUseCase {
	return &uploadUsecase{
		imageRepo: imageRepo,
	}
}

func (uc *uploadUsecase) EventImageUpload(files []*multipart.FileHeader) ([]string, error) {
	for _, fileHeader := range files {
		if fileHeader.Size > MAX_UPLOAD_SIZE {
			return nil, fmt.Errorf("the uploaded file is too big. Please choose an file that's less than 10MB in size")
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
		// fmt.Println("fileName: ", files[i])
		// image, err := uc.imageRepo.GetImageByUrl(fileName)
		// if err != nil && err.Error() != "record not found" {
		// 	return nil, err
		// }
		// fmt.Println("file: ", image)
		// if image.Id != 0 {
		// 	imagesSlice = append(imagesSlice, image.ImageUrl)
		// 	continue
		// }

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

func (uc *uploadUsecase) UploadAvatar(file *multipart.FileHeader) (string, error) {
	if file.Size > MAX_UPLOAD_AVATAR_SIZE {
		return "", fmt.Errorf("the uploaded file is too big. Please choose an file that's less than 5MB in size")
	}

	var cld, err = cloudinary.New()
	if err != nil {
		return "", fmt.Errorf("failed to intialize Cloudinary")
	}

	var ctx = context.Background()
	fileName := file.Filename

	fileOpen, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file")
	}
	defer fileOpen.Close()

	uploadResult, err := cld.Upload.Upload(
		ctx,
		fileOpen,
		uploader.UploadParams{
			PublicID: "avatars/" + fileName + "_" + pkg.RandomID(6),
		})
	if err != nil {
		return "", fmt.Errorf("failed to upload file")
	}

	return uploadResult.SecureURL, nil
}
