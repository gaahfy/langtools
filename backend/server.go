package main

import (
    "fmt"
    "log"
	"os"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Langtools updated %s!", r.URL.Path[1:])
}

func main() {
	port := os.Getenv("LANGTOOLS_BACKEND_PORT")

    http.HandleFunc("/", handler)
	fmt.Printf("Server is running on port %s\n", port)
    log.Fatal(http.ListenAndServe(port, nil))
}