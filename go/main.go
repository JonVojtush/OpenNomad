package main

import (
	"net/http"

	"github.com/yosssi/gcss"
)

var (
	w http.ResponseWriter
	r *http.Request
)

// Pre-process the .gcss stylesheet into .css
func GCSS() {
	cssPath, err := gcss.CompileFile("web/styles/style.gcss")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.ServeFile(w, r, cssPath)
}

func main() {
	GCSS()
}
