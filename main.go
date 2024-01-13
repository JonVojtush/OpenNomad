// https://github.com/golang/go/wiki/WebAssembly#getting-started

package main

import (
	"log"

	"github.com/yosssi/gcss"
)

func GCSS() {
	if _, err := gcss.CompileFile("styles/style.gcss"); err != nil {
		// http.Error(window, err.Error(), http.StatusInternalServerError)
		log.Print(err)
	} /*else {
		http.ServeFile(window, r, cssPath)
	}*/
}

func main() {
	GCSS()
}
