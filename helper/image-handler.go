package helper

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"fmt"
)

func HandleUploadImage(request *http.Request, path string, field string) string {

	file, _, err := request.FormFile(field)

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
	}
	defer file.Close()
	// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	// fmt.Printf("File Size: %+v\n", handler.Size)
	// fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := os.CreateTemp(path, "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)
	return strings.Split(filepath.ToSlash(tempFile.Name()), "/")[1]
}
