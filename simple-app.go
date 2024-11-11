package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json := map[string]string{
			"nodeName":     os.Getenv("NODE_NAME"),
			"podName":      os.Getenv("POD_NAME"),
			"podNamespace": os.Getenv("POD_NAMESPACE"),
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"nodeName":"%s","podName":"%s","podNamespace":"%s"}`, json["nodeName"], json["podName"], json["podNamespace"])
	})

	http.ListenAndServe(":80", nil)
}
