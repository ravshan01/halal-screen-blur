package images

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

var imagesService = ImagesService{}

var img []byte

func TestImagesService_CheckIsValidImage(t *testing.T) {
	t.Run("valid image", func(t *testing.T) {
		valid := imagesService.CheckIsValidImage(img)
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
		_, err := imagesService.Blur(img, -1)
		if !errors.Is(err, InvalidBlurPercentageError) {
			t.Fatalf("expected error to be InvalidBlurPercentageError")
		}

		_, err = imagesService.Blur(img, 101)
		if !errors.Is(err, InvalidBlurPercentageError) {
			t.Fatalf("expected error to be InvalidBlurPercentageError")
		}
	})

	t.Run("should return blurred image", func(t *testing.T) {
		blurredImgBytes, err := imagesService.Blur(img, 0)
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

func TestImagesService_Crop(t *testing.T) {
	t.Run("should return InvalidImageError if image is invalid", func(t *testing.T) {
		_, err := imagesService.Crop([]byte{}, [4]int{0, 0, 0, 0})
		if !errors.Is(err, InvalidImageError) {
			t.Fatalf("expected error to be InvalidImageError")
		}
	})

	t.Run("should return cropped image", func(t *testing.T) {
		x1, y1, width, height := 0, 0, 400, 200

		croppedImgBytes, err := imagesService.Crop(img, [4]int{x1, y1, width, height})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if len(*croppedImgBytes) == 0 {
			t.Fatalf("expected cropped image to have bytes")
		}

		croppedImgPointer, toImgErr := imagesService.bytesToImage(*croppedImgBytes)
		if toImgErr != nil {
			t.Fatalf("failed to convert bytes to image")
		}

		croppedImg := *croppedImgPointer
		if croppedImg.Bounds().Dx() != width || croppedImg.Bounds().Dy() != height {
			t.Fatalf("expected cropped image to have width %d and height %d, got width %d and height %d", width, height, croppedImg.Bounds().Dx(), croppedImg.Bounds().Dy())
		}

		// Uncomment the following line to save then compare the cropped image
		//os.WriteFile("./mock/street_cropped.jpg", *croppedImgBytes, 0644)
	})
}

func TestMain(m *testing.M) {
	bytes, err := os.ReadFile("./mock/street.jpg")
	if err != nil {
		fmt.Println("failed to read image file")
		os.Exit(1)
	}
	img = bytes

	os.Exit(m.Run())
}
