package dicom_png

import (
	"image/png"
	"net/http"
	"os"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
)

// Gets the png image from the `storage` directory and returns it to the
// client.
func GetDicomPng(w http.ResponseWriter) {
	err := createPngImageFile()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bytes, err := os.ReadFile("./storage/image.png")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bytes)
}

// Creates a new png image file by converting the `image.dicom` file found in
// the `storage` directory if it exists.
func createPngImageFile() error {
	dataset, err := dicom.ParseFile("./storage/image.dicom", nil)
	if err != nil {
		return err
	}

	pixelDataElement, err := dataset.FindElementByTag(tag.PixelData)
	if err != nil {
		return err
	}

	pixelDataInfo := dicom.MustGetPixelDataInfo(pixelDataElement.Value)

	for _, frame := range pixelDataInfo.Frames {
		img, err := frame.GetImage()
		if err != nil {
			return err
		}

		file, err := os.Create("./storage/image.png")
		if err != nil {
			return err
		}
		defer file.Close()

		err = png.Encode(file, img)
		if err != nil {
			return err
		}
	}

	return nil
}
