package main

import (
	"fmt"
	"google.golang.org/grpc"
	"halal-screen-blur/app"
	"halal-screen-blur/config"
	"halal-screen-blur/proto"
	"log"
	"net"
)

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
	proto.RegisterBlurServiceServer(grpcServer, app.NewBlurServiceServer())

	log.Printf("server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
