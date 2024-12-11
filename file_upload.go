package filesage

import (
	"fmt"
	"log"
	"net/http"
)

func fileUploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid usage", http.StatusBadRequest)
		return
	}

	const maxUploadSize = 50 << 20 //50Mb
	if r.ContentLength > maxUploadSize {
		http.Error(w, "File size too large", http.StatusRequestEntityTooLarge)
		log.Printf("Attempt to upload file exceeding limit. Size: %v", r.ContentLength)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

	err := r.ParseMultipartForm(maxUploadSize)
	if err != nil {
		http.Error(w, "Form parse error", http.StatusRequestEntityTooLarge)
		log.Printf("Form parse error: %v", err)
		return
	}
	defer r.MultipartForm.RemoveAll()

	file, handler, err := r.FormFile("upload_file")
	if err != nil {
		http.Error(w, "Failure uploading file", http.StatusBadRequest)
		log.Printf("File upload error: %v", err)
		return
	}
	defer file.Close()

	err = SaveFile(file, handler.Filename)
	if err != nil {
		http.Error(w, "Failure saving file", http.StatusInternalServerError)
		log.Printf("Failure saving file: %v", err)
		return
	}

	log.Printf("Handled file upload %s of %v bytes", handler.Filename, handler.Size)
	fmt.Fprintf(w, "Uploaded file: %s\n", handler.Filename)
	fmt.Fprintf(w, "File size: %d\n", handler.Size)
	fmt.Fprintf(w, "MIME header: %v\n", handler.Header)
}
