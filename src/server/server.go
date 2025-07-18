package server

import (
	"fmt"
	"net/http"
)

func ServeContent(port int, contentPath string) error {
	fs := http.FileServer(http.Dir(contentPath))
	http.Handle("/", fs)
	fmt.Printf("Serving content on port: %d\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}