package file

import (
	"bytes"
	"client/config"
	"client/utils/token"
	"encoding/json"
	"fmt"
	"github.com/dustin/go-humanize"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const chunkSize = 1024 * 1024 * 10

type Info struct {
	Filename string `json:"filename"`
	Size     string `json:"size"`
}

// Upload sends file to server
func Upload(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Printf("error opening file: %v", err)
		return
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		log.Printf("error get file stats: %v", err)
		return
	}

	totalSize := humanize.Bytes(uint64(fileInfo.Size()))
	log.Printf("Starting upload file: %s, size: %d \n", filePath, totalSize)

	var bytesSent uint64 = 0
	for {
		chunk := make([]byte, chunkSize)
		n, err := f.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			log.Printf("error reading file: %v", err)
			return
		}

		err = sendChunk(chunk[:n], filePath, bytesSent)
		if err != nil {
			log.Printf("error sending chunk: %v", err)
			return
		}

		bytesSent += uint64(n)
		log.Printf("Sent chunk: %d bytes, total sent: %s/%s \n", n, humanize.Bytes(bytesSent), totalSize)
	}

	log.Println("File upload completed successfully")
	return
}

func Download(fileName string) {
	url := strings.Join([]string{config.Envs.ServerHost, "/api/v1/files/", fileName}, "")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}

	req.Header.Set("Authorization", token.LoadToken())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: Server returned status: %s", resp.Status)
		return
	}

	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("Error creating file: %v", err)
		return
	}
	defer file.Close()

	n, err := io.Copy(file, resp.Body)
	if err != nil {
		log.Printf("Error writing file: %v", err)
		return
	}

	log.Printf("Downloaded file: %s, size: %s \n", fileName, humanize.Bytes(uint64(n)))
}

func Index() {
	url := strings.Join([]string{config.Envs.ServerHost, "/api/v1/files"}, "")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("error creating request: %v", err.Error())
	}
	req.Header.Set("Authorization", token.LoadToken())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error sending request: %v", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("server response: %v", resp.Status)
	}

	var files []Info
	if err := json.NewDecoder(resp.Body).Decode(&files); err != nil {
		log.Fatalf("error parsing json: %v", err.Error())
	}

	for index, file := range files {
		fmt.Printf("%d) Filename: %s, Size: %s \n", index+1, file.Filename, file.Size)
	}
}

func sendChunk(chunk []byte, filename string, offset uint64) error {
	url := strings.Join([]string{config.Envs.ServerHost, "/api/v1/files"}, "")

	req, err := http.NewRequest("POST", url, bytes.NewReader(chunk))
	if err != nil {
		return err
	}

	req.Header.Set("X-Filename", filepath.Base(filename))
	req.Header.Set("X-Offset", strconv.FormatUint(offset, 10))
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Authorization", token.LoadToken())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned status: %s, response: %s", resp.Status, respBody)
	}

	return nil
}
