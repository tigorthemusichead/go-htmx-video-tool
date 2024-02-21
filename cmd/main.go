package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"video-tool/internal/handlers/merge"
)

const op = "internal/main"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	r.HandleFunc("/api/merge", merge_handlers.MergePhotosAndAudioHandler)

	r.HandleFunc("/htmx/merge", merge_handlers.MergePhotosAndAudioHandlerHTMX)

	r.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("/tmp/files"))))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	//credentials := gHandlers.AllowCredentials()
	//methods := gHandlers.AllowedMethods([]string{"*"})
	//ttl := gHandlers.MaxAge(3600)
	//origins := gHandlers.AllowedOrigins([]string{"*"})
	//
	//c := gHandlers.CORS(credentials, methods, ttl, origins)(r)

	fmt.Println("Starting api on port 8081")

	//log.Fatal(http.ListenAndServeTLS(":8081", "server.crt", "server.key", c))
	log.Fatal(http.ListenAndServe(":8081", r))
}
