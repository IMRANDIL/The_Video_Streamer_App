package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

// StreamVideoHandler streams a video from a local directory
func StreamVideoHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the filename from the URL
	vars := mux.Vars(r)
	filename := vars["filename"]

	dir, err := os.Getwd()
	if err != nil {
		// Handle the error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dir = filepath.Join(dir, "videos")
	// Open the file for streaming
	file, err := os.Open(filepath.Join(dir, filename))

	if err != nil {
		// Handle the error if file not found
		if os.IsNotExist(err) {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	defer file.Close() // don't forget to close the file
	// Set the content type header based on file extension
	contentType := "video/mp4"
	if filepath.Ext(filename) == ".webm" {
		contentType = "video/webm"
	} else if filepath.Ext(filename) == ".ogv" {
		contentType = "video/ogg"
	}
	w.Header().Set("Content-Type", contentType)

	// Stream the video
	if _, err := io.Copy(w, file); err != nil {
		// Handle the error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
