package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Node Name: %s\nPod Name: %s\nPod Namespace: %s\n", os.Getenv("NODE_NAME"), os.Getenv("POD_NAME"), os.Getenv("POD_NAMESPACE"))
	})

	http.ListenAndServe(":80", nil)
}
