package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"halal-screen-blur/config"
	"halal-screen-blur/proto"
	"log"
	"net"
)

type server struct {
	proto.UnimplementedBlurServiceServer
}

func (s *server) BlurImages(ctx context.Context, req *proto.BlurImagesRequest) (*proto.BlurImageResponse, error) {
	images := req.GetImages()
	fmt.Println("images:", images)

	return &proto.BlurImageResponse{}, nil
}

func main() {
	cfg, cfgErr := config.LoadConfig(".env")
	if cfgErr != nil {
		log.Fatalf("failed to load config: %v", cfgErr)
	}

	listener, listErr := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if listErr != nil {
		log.Fatalf("failed to listen: %v", listErr)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterBlurServiceServer(grpcServer, &server{})

	log.Printf("server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
