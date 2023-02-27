package attribute

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
)

// Gets an attribute from a dicom image file based on a given `tag` received
// via a query paramater. Returns the result to the client as a JSON response.
func GetDicomAttribute(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("tag")
	tags := strings.Split(param, ",")
	if len(tags) != 2 {
		http.Error(w, "invalid query parameter", http.StatusBadRequest)
		return
	}

	tag, err := createNewTag(tags)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dataset, err := dicom.ParseFile("./storage/image.dicom", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := dataset.FindElementByTagNested(tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// Creates a new `tag.Tag` struct needed to search the dicom dataset by tag.
func createNewTag(tags []string) (tag.Tag, error) {
	group, err := strconv.ParseUint(tags[0], 16, 16)
	if err != nil {
		return tag.Tag{}, err
	}

	element, err := strconv.ParseUint(tags[1], 16, 16)
	if err != nil {
		return tag.Tag{}, err
	}

	tag := tag.Tag{
		Group:   uint16(group),
		Element: uint16(element),
	}

	return tag, nil
}
