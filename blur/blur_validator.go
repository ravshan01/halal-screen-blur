package blur

import (
	"fmt"
	"halal-screen-blur/config"
	"halal-screen-blur/images"
	"halal-screen-blur/proto"
)

type BlurValidator struct {
	imagesService images.IIMagesService
}

func NewBlurValidator() *BlurValidator {
	return &BlurValidator{
		imagesService: &images.ImagesService{},
	}
}

func (v *BlurValidator) ValidateImagesCount(imagesForBlur []*proto.ImageForBlur) BlurValidationResult[*proto.BlurImagesResponse] {
	existsResult := v.ValidateImagesExist(imagesForBlur)
	if !existsResult.Success {
		return existsResult
	}

	limitsResult := v.ValidateImagesLimits(imagesForBlur)
	if !limitsResult.Success {
		return limitsResult
	}

	return BlurValidationResult[*proto.BlurImagesResponse]{Success: true}
}

func (v *BlurValidator) ValidateImagesExist(imagesForBlur []*proto.ImageForBlur) BlurValidationResult[*proto.BlurImagesResponse] {
	if len(imagesForBlur) == 0 {
		return BlurValidationResult[*proto.BlurImagesResponse]{
			Success: false,
			ResWithErr: &proto.BlurImagesResponse{
				Error: &proto.BlurError{
					Code:    proto.BlurErrorCode_BadRequest,
					Message: "No images",
				},
			},
		}
	}

	return BlurValidationResult[*proto.BlurImagesResponse]{Success: true}
}

func (v *BlurValidator) ValidateImagesLimits(imagesForBlur []*proto.ImageForBlur) BlurValidationResult[*proto.BlurImagesResponse] {
	cfg := config.GetConfig()

	if len(imagesForBlur) > cfg.MaxImagesPerRequest {
		return BlurValidationResult[*proto.BlurImagesResponse]{
			Success: false,
			ResWithErr: &proto.BlurImagesResponse{
				Error: &proto.BlurError{
					Code:    proto.BlurErrorCode_MaxImagesExceeded,
					Message: fmt.Sprintf("Too many images, max %d", cfg.MaxImagesPerRequest),
				},
			},
		}
	}

	return BlurValidationResult[*proto.BlurImagesResponse]{Success: true}
}

func (v *BlurValidator) ValidateImage(imageForBlur *proto.ImageForBlur) BlurValidationResult[*proto.BlurredImage] {
	sizeResult := v.ValidateImageSize(imageForBlur)
	if !sizeResult.Success {
		return sizeResult
	}

	contentResult := v.ValidateImageContent(imageForBlur)
	if !contentResult.Success {
		return contentResult
	}

	return BlurValidationResult[*proto.BlurredImage]{Success: true}
}

func (v *BlurValidator) ValidateImageSize(imageForBlur *proto.ImageForBlur) BlurValidationResult[*proto.BlurredImage] {
	cfg := config.GetConfig()

	if len(imageForBlur.GetContent()) > cfg.MaxImageSize {
		return BlurValidationResult[*proto.BlurredImage]{
			Success: false,
			ResWithErr: &proto.BlurredImage{
				Error: &proto.BlurredImage_Error{
					Code:    proto.BlurredImage_MaxSizeExceeded,
					Message: fmt.Sprintf("Image size exceeded. max %d bytes", cfg.MaxImageSize),
				},
			},
		}
	}

	return BlurValidationResult[*proto.BlurredImage]{Success: true}
}

func (v *BlurValidator) ValidateImageContent(imageForBlur *proto.ImageForBlur) BlurValidationResult[*proto.BlurredImage] {
	if !v.imagesService.CheckIsValidImage(imageForBlur.GetContent()) {
		return BlurValidationResult[*proto.BlurredImage]{
			Success: false,
			ResWithErr: &proto.BlurredImage{
				Error: &proto.BlurredImage_Error{
					Code:    proto.BlurredImage_InvalidImage,
					Message: "Invalid image content",
				},
			},
		}
	}

	return BlurValidationResult[*proto.BlurredImage]{Success: true}
}
