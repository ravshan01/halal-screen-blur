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

func (s *BlurServiceServer) BlurImages(_ context.Context, req *proto.BlurImagesRequest) (*proto.BlurImageResponse, error) {
	imagesForBlur := req.GetImages()
	fmt.Println("imagesForBlur:", imagesForBlur)

	res := s.checkHasImagesAndNotTooMany(imagesForBlur)
	if res != nil {
		return res, nil
	}

	return &proto.BlurImageResponse{}, nil
}

/*
Return `BlurImagesResponse` with error if no images provided or too many images
*/
func (s *BlurServiceServer) checkHasImagesAndNotTooMany(imagesForBlur []*proto.ImageForBlur) *proto.BlurImageResponse {
	if len(imagesForBlur) == 0 {
		return &proto.BlurImageResponse{
			Error: &proto.BlurError{
				Code:    proto.BlurErrorCode_BadRequest,
				Message: "no imagesForBlur provided",
			},
		}
	}

	cfg := config.GetConfig()
	if len(imagesForBlur) > cfg.MaxImagesPerRequest {
		return &proto.BlurImageResponse{
			Error: &proto.BlurError{
				Code:    proto.BlurErrorCode_BadRequest,
				Message: "too many imagesForBlur",
			},
		}
	}

	return nil
}
