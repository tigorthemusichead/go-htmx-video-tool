package merge_handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"time"
	"video-tool/internal/lib/files"
	"video-tool/internal/lib/tools"
)

const OutputDir = "/tmp/files"

type MergePhotosAndAudiosResponse struct {
	Success        bool
	VideoSrc       string
	VideoCreatedAt string
}

func MergePhotosAndAudioHandler(w http.ResponseWriter, r *http.Request) {
	mergePhotosAndAudioHandler(w, r, false)
}

func MergePhotosAndAudioHandlerHTMX(w http.ResponseWriter, r *http.Request) {
	mergePhotosAndAudioHandler(w, r, true)
}

func mergePhotosAndAudioHandler(w http.ResponseWriter, r *http.Request, isHTMX bool) {
	formKeys := []string{"photoFile", "audioFile"}
	var createdFiles []string
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	for _, key := range formKeys {
		fileName, err := files.CreateFromFormData(w, r, key)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		createdFiles = append(createdFiles, fileName)
	}
	fmt.Println("Created files", createdFiles)
	err = files.CheckDir(OutputDir)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	outputFileName := base64.StdEncoding.EncodeToString([]byte(time.Now().String())) + ".mp4"
	outputFilePath := filepath.Join(OutputDir, outputFileName)
	err = tools.MergePhotosAndAudios(createdFiles, outputFilePath)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	//http.Redirect(w, r, "/files/"+outputFileName, http.StatusSeeOther)
	data, err := json.Marshal(MergePhotosAndAudiosResponse{true, fmt.Sprintf("/files/%s", outputFileName), time.Now().UTC().Format(time.RFC3339)})
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	for _, filePath := range createdFiles {
		err = files.DeleteFile(filePath)
	}
	if err != nil {
		fmt.Errorf("%v", err)
	}
	go files.ScheduleFileDeletion(outputFilePath, time.Minute*5)
	w.WriteHeader(201)
	if isHTMX {
		w.Write([]byte(fmt.Sprintf("<video src=\"%s%s\" controls></video>", "/files/", outputFileName)))
	} else {
		w.Write(data)
	}
	return
}
