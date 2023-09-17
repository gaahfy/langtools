package main

import (
    "encoding/json"
    "fmt"
    "log"
    "path/filepath"
    "net/http"
    "github.com/gaahfy/langtools/backend/utils"
    "github.com/gaahfy/langtools/backend/config"
)

func handler(w http.ResponseWriter, r *http.Request) {
	domainName := config.DomainName()
    if(r.Host != domainName) {
        httpsURL := "https://" + domainName + r.URL.String()
        http.Redirect(w, r, httpsURL, http.StatusMovedPermanently)
        return
    }

    if(utils.ServeLetsEncryptIfAvailable(w, r)) {
        return
    }

    cleanedPath := filepath.Clean(r.URL.Path)
	if cleanedPath == "." {
		cleanedPath = "/"
	}
    w.Header().Set("Content-Type", "application/json")
    escapedJSON, err := json.Marshal("Path: "+cleanedPath[1:])
    if err != nil {
        fmt.Fprintf(w, "{\"message\":\"Error\"}", escapedJSON)
    }
    fmt.Fprintf(w, "{\"message\":%s}", escapedJSON)
}

func redirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	domainName := config.DomainName()
    if(utils.ServeLetsEncryptIfAvailable(w, r)) {
        return
    }
	httpsURL := "https://" + domainName + r.URL.String()
	http.Redirect(w, r, httpsURL, http.StatusMovedPermanently)
}

func main() {
    config.InitDB()

    portHTTP := config.HttpPort()
	portHTTPS := config.HttpsPort()
	domainName := config.DomainName()

    if config.IsProduction() {
        go http.ListenAndServe(portHTTP, http.HandlerFunc(redirectToHTTPS))

		// Configuration du serveur HTTPS
		http.HandleFunc("/", handler)
		fmt.Printf("Server is running on port %s (HTTP)\n", portHTTP)
		fmt.Printf("Server is running on port %s (HTTPS)\n", portHTTPS)
		log.Fatal(http.ListenAndServeTLS(portHTTPS, "/etc/letsencrypt/live/"+domainName+"/fullchain.pem", "/etc/letsencrypt/live/"+domainName+"/privkey.pem", nil))
    } else {
		// Si la variable LANGTOOLS_IS_PRODUCTION n'est pas "yes", exécutez simplement le serveur HTTP
		http.HandleFunc("/", handler)
		fmt.Printf("Server is running on port %s (HTTP)\n", portHTTP)
		log.Fatal(http.ListenAndServe(portHTTP, nil))
	}
}