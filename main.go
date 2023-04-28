package main

import (
	"fmt"
	"log"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {

	method := r.Method
	path := r.URL.Path
	rawQuery := r.URL.RawQuery

	fmt.Println(method + " " + path + " " + rawQuery)
	w.Write([]byte("Hello world!\n"))
	w.Write([]byte("Method: " + method + "\n"))
	w.Write([]byte("Path: " + path + "\n"))
	w.Write([]byte("RawQuery: " + rawQuery + "\n"))

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", requestHandler)

	// TODO Add a static file handler using http.FileServer
	staticHandler := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", staticHandler))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
