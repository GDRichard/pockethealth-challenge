package dicom_image

import (
	"io"
	"net/http"
	"os"
)

// Saves the dicom image data locally as a file inside the `storage`
// directory.
func PostDicomImage(w http.ResponseWriter, r *http.Request) {
	buf, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = os.WriteFile("./storage/image.dicom", buf, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "image successfully saved")
}
