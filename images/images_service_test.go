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

	// Check './mock/processed/mock-image__blurred.jpg' to see the blurred image
	t.Run("should return blurred image", func(t *testing.T) {
		blurredImgPointer, err := imagesService.Blur(img, 0)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		blurredImg := *blurredImgPointer
		if blurredImg.Bounds().Dx() != img.Bounds().Dx() || blurredImg.Bounds().Dy() != img.Bounds().Dy() {
			t.Fatalf("expected blurred image to have same dimensions as original image")
		}

		blurredImgBytes, toBytesErr := imagesService.ImageToBytes(blurredImg)
		if toBytesErr != nil {
			t.Fatalf("failed to convert image to bytes")
		}

		err = os.WriteFile("./mock/processed/mock-image__blurred.png", *blurredImgBytes, 0644)
		if err != nil {
			fmt.Println("failed to write blurred image")
		}
	})
}

func TestImagesService_Crop(t *testing.T) {
	// Check './mock/processed/mock-image__cropped.jpg' to see the cropped image
	t.Run("should return cropped image", func(t *testing.T) {
		x1, y1, width, height := 0, 0, 400, 200
		rect := image.Rect(x1, y1, x1+width, y1+height)

		croppedImgPointer, err := imagesService.Crop(img, rect)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		croppedImg := *croppedImgPointer
		if croppedImg.Bounds().Dx() != width || croppedImg.Bounds().Dy() != height {
			t.Fatalf("expected cropped image to have width %d and height %d, got width %d and height %d", width, height, croppedImg.Bounds().Dx(), croppedImg.Bounds().Dy())
		}

		croppedImgBytes, toBytesErr := imagesService.ImageToBytes(croppedImg)
		if toBytesErr != nil {
			t.Fatalf("failed to convert image to bytes")
		}

		err = os.WriteFile("./mock/processed/mock-image__cropped.png", *croppedImgBytes, 0644)
		if err != nil {
			fmt.Println("failed to write cropped image")
		}
	})
}

func TestImagesService_Paste(t *testing.T) {
	// Check './mock/processed/mock-image__pasted.jpg' to see the pasted image
	t.Run("should return pasted image", func(t *testing.T) {
		cropRect := image.Rect(0, 0, 400, 400)
		croppedImgPointer, cropErr := imagesService.Crop(img, cropRect)
		croppedImg := *croppedImgPointer
		if cropErr != nil {
			t.Fatalf("expected no error, got %v", cropErr)
		}

		pasteRect := image.Point{X: 200, Y: 200}
		pastedImgPointer, err := imagesService.Paste(img, croppedImg, pasteRect)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		pastedImg := *pastedImgPointer
		if pastedImg.Bounds().Dx() != img.Bounds().Dx() || pastedImg.Bounds().Dy() != img.Bounds().Dy() {
			t.Fatalf("expected pasted image to have same dimensions as original image")
		}

		pastedImgBytes, toBytesErr := imagesService.ImageToBytes(pastedImg)
		if toBytesErr != nil {
			t.Fatalf("failed to convert image to bytes")
		}

		err = os.WriteFile("./mock/processed/mock-image__pasted.png", *pastedImgBytes, 0644)
		if err != nil {
			fmt.Println("failed to write pasted image")
		}
	})
}

func TestImagesService_BytesToImage(t *testing.T) {
	t.Run("should return image", func(t *testing.T) {
		imgPointer, err := imagesService.BytesToImage(imgBytes)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if imgPointer == nil {
			t.Fatalf("expected image to be returned")
		}
	})

	t.Run("should return error if image is invalid", func(t *testing.T) {
		_, err := imagesService.BytesToImage([]byte{})
		if err == nil {
			t.Fatalf("expected error to be returned")
		}
	})
}

func TestImagesService_ImageToBytes(t *testing.T) {
	t.Run("should return bytes", func(t *testing.T) {
		bytesPointer, err := imagesService.ImageToBytes(img)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if bytesPointer == nil {
			t.Fatalf("expected bytes to be returned")
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
