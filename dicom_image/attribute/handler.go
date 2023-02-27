package attribute

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetDicomAttribute(w, r)
		return
	}
	http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
}
