package utils

import (
	"os"
    "net/http"
	"path/filepath"
	"strings"
)

func ServeLetsEncryptIfAvailable(w http.ResponseWriter, r *http.Request) (bool) {
	if r.URL.Path == "" || r.URL.Path == "/" {
		return false
	}
	currentExecutable, err := os.Executable()
    if err != nil {
        return false
    }
	currentDir := filepath.Dir(currentExecutable)
	filePath := currentDir + "/letsencrypt" + r.URL.Path
	fileInfo, err := os.Stat(filePath)
	if err == nil {
		absolutePath := fileInfo.Name()
		if strings.HasPrefix(absolutePath, currentDir) {
			fileServer := http.FileServer(http.Dir(currentDir + "/letsencrypt"))
			fileServer.ServeHTTP(w, r)
			return true
		}
	}
	return false
}