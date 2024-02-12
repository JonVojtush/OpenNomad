package main

import (
	"fmt"
	"log"
)

/*var (
	listen *string = flag.String("listen", ":8080", "listen address")
	dir    *string = flag.String("dir", "web", "directory to serve")
	w  http.ResponseWriter
	r   *http.Request
)

func launchServer() {
	log.Printf("Serving the %s/ directory on http://localhost%s", *dir, *listen)
	err := http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir)))
	log.Fatalln(err)
}*/

func main() {
	fmt.Println("B4 launchServer()")
	log.Println("B4 launchServer()")
	/*flag.Parse()
	launchServer()
	fmt.Println("B4 launchServer()")
	log.Println("B4 launchServer()")*/
}
