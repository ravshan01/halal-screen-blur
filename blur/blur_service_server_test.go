package blur

import (
	"context"
	"fmt"
	"halal-screen-blur/config"
	"halal-screen-blur/proto"
	"os"
	"testing"
)

var blurServiceServer *BlurServiceServer

func TestBlurServiceServerIsDefined(t *testing.T) {
	setupBlurServiceServerTest()
	defer teardownBlurServiceServerTest()

	if blurServiceServer == nil {
		t.Fatalf("expected blurServiceServer to be defined")
	}
}

func TestBlurServiceServer_BlurImages(t *testing.T) {
	setupBlurServiceServerTest()
	defer teardownBlurServiceServerTest()
	cfg := config.GetConfig()

	t.Run("should return BadRequest error if no images provided", func(t *testing.T) {
		res, err := blurServiceServer.BlurImages(context.Context(context.Background()), &proto.BlurImagesRequest{})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if res.GetError().GetCode() != proto.BlurErrorCode_BadRequest {
			t.Fatalf("expected BadRequest error, got %v", res.GetError().GetCode())
		}
	})

	t.Run("should return MaxImagesExceeded error if too many images provided", func(t *testing.T) {
		images := make([]*proto.ImageForBlur, cfg.MaxImagesPerRequest+1)
		res, err := blurServiceServer.BlurImages(context.Context(context.Background()), &proto.BlurImagesRequest{
			Images: images,
		})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if res.GetError().GetCode() != proto.BlurErrorCode_MaxImagesExceeded {
			t.Fatalf("expected MaxImagesExceeded error, got %v", res.GetError().GetCode())
		}
	})
}

func setupBlurServiceServerTest() {
	_, err := config.LoadConfig()
	if err != nil {
		fmt.Println("failed to load config, got:", err)
		os.Exit(1)
	}

	blurServiceServer = NewBlurServiceServer()
}

func teardownBlurServiceServerTest() {
}
