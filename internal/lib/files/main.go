package files

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const op = "/internal/lib/files"
const UploadDir = "/tmp/uploaded"

func CheckDir(dir string) error {
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dir, os.ModePerm)
		}
	}
	return err
}

func CreateFromFormData(w http.ResponseWriter, r *http.Request, formKey string) (string, error) {
	file, header, err := r.FormFile(formKey)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		w.Write([]byte(": audioFile not found in multipart body"))
		return "", err
	}
	err = CheckDir(UploadDir)
	if err != nil {
		return "", err
	}
	fileName := filepath.Join(UploadDir, header.Filename)
	fmt.Println(fileName)
	out, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

func DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	fmt.Printf("File %s deleted successfully.\n", filePath)
	return nil
}
