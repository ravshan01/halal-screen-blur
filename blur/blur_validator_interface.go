package blur

import "halal-screen-blur/proto"

type IBlurValidator interface {
	ValidateImagesCount(imagesForBlur []*proto.ImageForBlur) BlurValidationResult[*proto.BlurImagesResponse]
	ValidateImagesExist(imagesForBlur []*proto.ImageForBlur) BlurValidationResult[*proto.BlurImagesResponse]
	ValidateImagesLimits(imagesForBlur []*proto.ImageForBlur) BlurValidationResult[*proto.BlurImagesResponse]

	ValidateImage(imageForBlur *proto.ImageForBlur) BlurValidationResult[*proto.BlurredImage]
	ValidateImageSize(imageForBlur *proto.ImageForBlur) BlurValidationResult[*proto.BlurredImage]
	ValidateImageContent(imageForBlur *proto.ImageForBlur) BlurValidationResult[*proto.BlurredImage]
}

type BlurValidationResult[ResType any] struct {
	Success    bool
	ResWithErr ResType
}
