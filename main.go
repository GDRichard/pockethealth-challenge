package main

import (
	"fmt"
	"log"
	"net/http"

	"pockethealth/dicom_image"
	"pockethealth/dicom_image/attribute"
	"pockethealth/dicom_image/dicom_png"
)

func main() {
	http.HandleFunc("/dicom_image", dicom_image.Handler)
	http.HandleFunc("/dicom_image/attribute", attribute.Handler)
	http.HandleFunc("/dicom_image/png", dicom_png.Handler)

	fmt.Println("server running on port :3333")
	log.Fatal(http.ListenAndServe(":3333", nil))
}
