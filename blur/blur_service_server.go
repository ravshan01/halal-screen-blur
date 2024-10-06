package blur

import (
	"context"
	"fmt"
	"halal-screen-blur/config"
	"halal-screen-blur/images"
	"halal-screen-blur/proto"
)

type BlurServiceServer struct {
	proto.UnimplementedBlurServiceServer

	imagesService images.IIMagesService
}

func NewBlurServiceServer() *BlurServiceServer {
	return &BlurServiceServer{
		imagesService: &images.ImagesService{},
	}
}

func (s *BlurServiceServer) BlurImages(_ context.Context, req *proto.BlurImagesRequest) (*proto.BlurImagesResponse, error) {
	imagesForBlur := req.GetImages()
	fmt.Println("imagesForBlur:", imagesForBlur)

	res := s.checkHasImagesAndNotTooMany(imagesForBlur)
	if res != nil {
		return res, nil
	}

	return &proto.BlurImagesResponse{}, nil
}

/*
Return `BlurImagesResponse` with error if no images provided or too many images
*/
func (s *BlurServiceServer) checkHasImagesAndNotTooMany(imagesForBlur []*proto.ImageForBlur) *proto.BlurImagesResponse {
	if len(imagesForBlur) == 0 {
		return &proto.BlurImagesResponse{
			Error: &proto.BlurError{
				Code:    proto.BlurErrorCode_BadRequest,
				Message: "no imagesForBlur provided",
			},
		}
	}

	cfg := config.GetConfig()
	if len(imagesForBlur) > cfg.MaxImagesPerRequest {
		return &proto.BlurImagesResponse{
			Error: &proto.BlurError{
				Code:    proto.BlurErrorCode_MaxImagesExceeded,
				Message: "too many imagesForBlur",
			},
		}
	}

	return nil
}
