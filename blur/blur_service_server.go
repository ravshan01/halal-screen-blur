package blur

import (
	"context"
	"halal-screen-blur/images"
	"halal-screen-blur/proto"
	"image"
)

type BlurServiceServer struct {
	proto.UnimplementedBlurServiceServer

	imagesService images.IIMagesService
	blurValidator IBlurValidator
}

func NewBlurServiceServer() *BlurServiceServer {
	return &BlurServiceServer{
		imagesService: images.NewImagesService(),
		blurValidator: NewBlurValidator(),
	}
}

func (s *BlurServiceServer) BlurImages(_ context.Context, req *proto.BlurImagesRequest) (*proto.BlurImagesResponse, error) {
	imagesForBlur := req.GetImages()

	valResult := s.blurValidator.ValidateImagesCount(imagesForBlur)
	if !valResult.Success {
		return valResult.ResWithErr, nil
	}

	blurredImages := make([]*proto.BlurredImage, 0, len(imagesForBlur))

	for _, elem := range imagesForBlur {
		blurredImg, blurErr := s.blurImage(elem)
		if blurErr != nil {
			blurredImages = append(blurredImages, &proto.BlurredImage{
				Error: &proto.BlurredImage_Error{
					Code:    proto.BlurredImage_InternalError,
					Message: blurErr.Error(),
				},
			})
		}

		blurredImages = append(blurredImages, blurredImg)
	}

	return &proto.BlurImagesResponse{
		BlurredImages: blurredImages,
	}, nil
}

func (s *BlurServiceServer) blurImage(imageForBlur *proto.ImageForBlur) (*proto.BlurredImage, error) {
	valResult := s.blurValidator.ValidateImage(imageForBlur)
	if !valResult.Success {
		return valResult.ResWithErr, nil
	}

	img, toImgErr := s.imagesService.BytesToImage(imageForBlur.GetContent())
	if toImgErr != nil {
		return nil, toImgErr
	}
	var blurredImg = *img

	for _, coords := range imageForBlur.GetCoords() {
		rect := image.Rect(int(coords.GetX()), int(coords.GetY()), int(coords.GetWidth()), int(coords.GetHeight()))

		// TODO: create and use IImagesService.BlurPart method
		croppedImg, cropErr := s.imagesService.Crop(blurredImg, rect)
		if cropErr != nil {
			return nil, cropErr
		}

		croppedBlurredImg, blurErr := s.imagesService.Blur(*croppedImg, int32(imageForBlur.GetPercentage()))
		if blurErr != nil {
			return nil, blurErr
		}

		pastedImg, pasteErr := s.imagesService.Paste(blurredImg, *croppedBlurredImg, image.Point{int(coords.GetX()), int(coords.GetY())})
		if pasteErr != nil {
			return nil, pasteErr
		}

		blurredImg = *pastedImg
	}

	blurredImgBytes, imgToBytesErr := s.imagesService.ImageToBytes(blurredImg)
	if imgToBytesErr != nil {
		return nil, imgToBytesErr
	}

	return &proto.BlurredImage{
		Content: *blurredImgBytes,
	}, nil
}
