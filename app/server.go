package app

import (
	"context"
	"fmt"
	"halal-screen-blur/config"
	"halal-screen-blur/proto"
)

type Server struct {
	proto.UnimplementedBlurServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) BlurImages(_ context.Context, req *proto.BlurImagesRequest) (*proto.BlurImageResponse, error) {
	images := req.GetImages()
	fmt.Println("images:", images)

	res := s.checkHasImagesAndNotTooMany(images)
	if res != nil {
		return res, nil
	}

	return &proto.BlurImageResponse{}, nil
}

/*
Return `BlurImagesResponse` with error if no images provided or too many images
*/
func (s *Server) checkHasImagesAndNotTooMany(images []*proto.ImageForBlur) *proto.BlurImageResponse {
	if len(images) == 0 {
		return &proto.BlurImageResponse{
			Error: &proto.BlurError{
				Code:    proto.BlurErrorCode_BadRequest,
				Message: "no images provided",
			},
		}
	}

	cfg := config.GetConfig()
	if len(images) > cfg.MaxImagesPerRequest {
		return &proto.BlurImageResponse{
			Error: &proto.BlurError{
				Code:    proto.BlurErrorCode_BadRequest,
				Message: "too many images",
			},
		}
	}

	return nil
}
