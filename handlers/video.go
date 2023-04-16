package handlers

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// StreamVideoHandler streams a video from a local directory
func StreamVideoHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the filename from the URL
	vars := mux.Vars(r)
	filename := vars["filename"]

	// Open the file for streaming
	file, err := http.Dir("/path/to/local/videos").Open(filename)
	if err != nil {
		// Handle the error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the content type header
	w.Header().Set("Content-Type", "video/mp4")

	// Stream the video
	if _, err := io.Copy(w, file); err != nil {
		// Handle the error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
