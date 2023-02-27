package dicom_image

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		PostDicomImage(w, r)
		return
	}
	http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
}
