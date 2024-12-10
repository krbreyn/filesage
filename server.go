package filesage

import (
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/upload", fileUploadHandler)
	http.HandleFunc("/download", fileDownloadHandler)

	const port = ":8080"
	log.Printf("server running on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
