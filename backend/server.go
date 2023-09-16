package main

import (
    "encoding/json"
    "fmt"
    "log"
	"os"
    "path/filepath"
    "net/http"
    "github.com/gaahfy/langtools/backend/utils"
)

func handler(w http.ResponseWriter, r *http.Request) {
    cleanedPath := filepath.Clean(r.URL.Path)
	if cleanedPath == "." {
		cleanedPath = "/"
	}

    if(utils.ServeLetsEncryptIfAvailable(w, r)) {
        return
    }
    w.Header().Set("Content-Type", "application/json")
    escapedJSON, err := json.Marshal("Path: "+cleanedPath[1:])
    if err != nil {
        fmt.Fprintf(w, "{\"message\":\"Error\"}", escapedJSON)
    }
    fmt.Fprintf(w, "{\"message\":%s}", escapedJSON)
}

func main() {
	port := os.Getenv("LANGTOOLS_BACKEND_PORT")

    http.HandleFunc("/", handler)
	fmt.Printf("Server is running on port %s\n", port)
    log.Fatal(http.ListenAndServe(port, nil))
}