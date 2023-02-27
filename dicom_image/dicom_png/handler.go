package dicom_png

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetDicomPng(w)
		return
	}
	http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
}
