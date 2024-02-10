package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/yosssi/gcss"
)

var (
	listen *string = flag.String("listen", ":8080", "listen address")
	dir    *string = flag.String("dir", "web", "directory to serve")
	/*resp  http.ResponseWriter
	req   *http.Request*/
)

func launchServer() {
	log.Printf("Serving the %s/ directory on http://localhost%s", *dir, *listen)
	err := http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir)))
	log.Fatalln(err)
}

func GCSS() {
	_, err := gcss.CompileFile("web/styles/style.gcss") // Pre-process the .gcss stylesheet into .css
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	GCSS()
	launchServer()
	fmt.Println("Hello, Jon!")
}
