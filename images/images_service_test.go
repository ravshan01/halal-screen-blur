package images

import (
	"errors"
	"fmt"
	"image"
	"os"
	"testing"
)

var imagesService = ImagesService{}

var imgBytes []byte
var img image.Image

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

	// Check './mock/processed/mock-image__blurred.jpg' to see the blurred image
	t.Run("should return blurred image", func(t *testing.T) {
		blurredImgBytes, err := imagesService.Blur(imgBytes, 0)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if len(*blurredImgBytes) == 0 {
			t.Fatalf("expected blurred image to have bytes")
		}

		blurredImgPointer, toImgErr := imagesService.BytesToImage(*blurredImgBytes)
		if toImgErr != nil {
			t.Fatalf("failed to convert bytes to image")
		}

		blurredImg := *blurredImgPointer
		if blurredImg.Bounds().Dx() != img.Bounds().Dx() || blurredImg.Bounds().Dy() != img.Bounds().Dy() {
			t.Fatalf("expected blurred image to have same dimensions as original image")
		}

		err = os.WriteFile("./mock/processed/mock-image__blurred.png", *blurredImgBytes, 0644)
		if err != nil {
			fmt.Println("failed to write blurred image")
		}
	})
}

func TestImagesService_Crop(t *testing.T) {
	t.Run("should return InvalidImageError if image is invalid", func(t *testing.T) {
		_, err := imagesService.Crop([]byte{}, [4]int{0, 0, 0, 0})
		if !errors.Is(err, InvalidImageError) {
			t.Fatalf("expected error to be InvalidImageError")
		}
	})

	// Check './mock/processed/mock-image__cropped.jpg' to see the cropped image
	t.Run("should return cropped image", func(t *testing.T) {
		x1, y1, width, height := 0, 0, 400, 200

		croppedImgBytes, err := imagesService.Crop(imgBytes, [4]int{x1, y1, width, height})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if len(*croppedImgBytes) == 0 {
			t.Fatalf("expected cropped image to have bytes")
		}

		croppedImgPointer, toImgErr := imagesService.BytesToImage(*croppedImgBytes)
		if toImgErr != nil {
			t.Fatalf("failed to convert bytes to image")
		}

		croppedImg := *croppedImgPointer
		if croppedImg.Bounds().Dx() != width || croppedImg.Bounds().Dy() != height {
			t.Fatalf("expected cropped image to have width %d and height %d, got width %d and height %d", width, height, croppedImg.Bounds().Dx(), croppedImg.Bounds().Dy())
		}

		err = os.WriteFile("./mock/processed/mock-image__cropped.png", *croppedImgBytes, 0644)
		if err != nil {
			fmt.Println("failed to write cropped image")
		}
	})
}

func TestMain(m *testing.M) {
	bytes, err := os.ReadFile("./mock/mock-image.jpg")
	if err != nil {
		fmt.Println("failed to read image file")
		os.Exit(1)
	}
	imgBytes = bytes

	origImg, toImgErr := imagesService.BytesToImage(bytes)
	if toImgErr != nil {
		fmt.Println("failed to convert bytes to image")
		os.Exit(1)
	}

	img = *origImg

	os.Exit(m.Run())
}
