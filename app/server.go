package app

import (
	"context"
	"fmt"
	"halal-screen-blur/proto"
)

type Server struct {
	proto.UnimplementedBlurServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) BlurImages(ctx context.Context, req *proto.BlurImagesRequest) (*proto.BlurImageResponse, error) {
	images := req.GetImages()
	fmt.Println("images:", images)

	return &proto.BlurImageResponse{}, nil
}
