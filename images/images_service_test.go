package images

import (
	"errors"
	"fmt"
	"image"
	"os"
	"testing"
)

var imagesService = ImagesService{}

var img image.Image
var imgBytes []byte

func TestImagesService_CheckIsValidImage(t *testing.T) {
	t.Run("valid image", func(t *testing.T) {
		valid := imagesService.CheckIsValidImage(imgBytes)
		if !valid {
			t.Fatalf("expected image to be valid")
		}
	})

	t.Run("invalid image", func(t *testing.T) {
		valid := imagesService.CheckIsValidImage([]byte{})
		if valid {
			t.Fatalf("expected image to be invalid")
		}
	})
}

func TestImagesService_Blur(t *testing.T) {
	t.Run("should return InvalidImageError if image is invalid", func(t *testing.T) {
		_, err := imagesService.Blur([]byte{}, 50)
		if !errors.Is(err, InvalidImageError) {
			t.Fatalf("expected error to be InvalidImageError")
		}
	})

	t.Run("should return InvalidBlurPercentageError if percentage < 0 or > 100 ", func(t *testing.T) {
		_, err := imagesService.Blur(imgBytes, -1)
		if !errors.Is(err, InvalidBlurPercentageError) {
			t.Fatalf("expected error to be InvalidBlurPercentageError")
		}

		_, err = imagesService.Blur(imgBytes, 101)
		if !errors.Is(err, InvalidBlurPercentageError) {
			t.Fatalf("expected error to be InvalidBlurPercentageError")
		}
	})

	t.Run("should return blurred image", func(t *testing.T) {
		blurredImgBytes, err := imagesService.Blur(imgBytes, 0)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if len(*blurredImgBytes) == 0 {
			t.Fatalf("expected blurred image to have bytes")
		}

		// Uncomment the following line to save then compare the blurred image
		//os.WriteFile("./mock/street_blurred.jpg", *blurredImgBytes, 0644)
	})
}

func TestMain(m *testing.M) {
	bytes, readErr := os.ReadFile("./mock/street.jpg")
	if readErr != nil {
		fmt.Println("failed to read image file")
		os.Exit(1)
	}

	origImg, err := imagesService.bytesToImage(bytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	imgBytes = bytes
	img = *origImg

	os.Exit(m.Run())
}
