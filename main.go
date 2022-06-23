package main

import (
	"firstweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Handler)
	mux.HandleFunc("/post", handler.PostGet)
	mux.HandleFunc("/siswa", handler.SiswaHandler)
	mux.HandleFunc("/form", handler.Form)
	mux.HandleFunc("/submit", handler.Submit)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting web on port 8822")

	err := http.ListenAndServe(":8822", mux)

	log.Fatal(err)
}