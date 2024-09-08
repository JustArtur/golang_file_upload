package controllers

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"server/helpers"
	"strconv"
	"strings"
)

type FileInfo struct {
	Filename string `json:"filename"`
	Size     string `json:"size"`
}

// Upload uploads file from client
func Upload(w http.ResponseWriter, r *http.Request) {
	filename := r.Header.Get("X-Filename")
	offsetStr := r.Header.Get("X-Offset")

	if filename == "" || offsetStr == "" {
		http.Error(w, "Missing filename or offset", http.StatusBadRequest)
		return
	}

	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid offset", http.StatusBadRequest)
		return
	}

	filePath := strings.Join([]string{"storage", strconv.Itoa(helpers.GetUserIDFromContext(r)), filename}, "/")
	log.Println(filePath)
	err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		log.Printf("error creating directories: %v", err)
		return
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error opening file: %v", err), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer file.Close()

	_, err = file.Seek(offset, 0)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error seeking file: %v", err), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	n, err := io.Copy(file, r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error writing to file: %v", err), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	log.Printf("Received chunk: %s bytes, written at offset: %d\n", humanize.Bytes(uint64(n)), offset)
	w.WriteHeader(http.StatusOK)
}

// Download send files to client
func Download(w http.ResponseWriter, r *http.Request) {
	filename := mux.Vars(r)["fileName"]
	if filename == "" {
		http.Error(w, "Missing filename", http.StatusBadRequest)
		return
	}

	filePath := filepath.Join("storage", strconv.Itoa(helpers.GetUserIDFromContext(r)), filename)

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error opening file: %v", err), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	w.Header().Set("Content-Type", "application/octet-stream")

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error sending file: %v", err), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	userDir := "storage/" + strconv.Itoa(helpers.GetUserIDFromContext(r))
	files, err := os.ReadDir(userDir)
	if err != nil {
		http.Error(w, "error reading directory", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var fileInfos []FileInfo
	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(userDir, file.Name())
			fileInfo, err := os.Stat(filePath)
			if err != nil {
				continue
			}
			fileInfos = append(fileInfos, FileInfo{
				Filename: file.Name(),
				Size:     humanize.Bytes(uint64(fileInfo.Size())),
			})
		}
	}

	helpers.SendResponse(w, http.StatusOK, fileInfos)
}
