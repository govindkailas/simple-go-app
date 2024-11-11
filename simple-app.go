package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"nodeName":     os.Getenv("NODE_NAME"),
			"podName":      os.Getenv("POD_NAME"),
			"podNamespace": os.Getenv("POD_NAMESPACE"),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":80", nil)
}
