package utils

import (
	"os"
    "net/http"
)

func ServeLetsEncryptIfAvailable(w http.ResponseWriter, r *http.Request) (bool) {
	if r.URL.Path == "" || r.URL.Path == "/" {
		return false
	}
	filePath := "./letsencrypt" + r.URL.Path
	_, err := os.Stat(filePath)
	if err == nil {
        fileServer := http.FileServer(http.Dir("./letsencrypt"))
        fileServer.ServeHTTP(w, r)
		return true
	}
	return false
}