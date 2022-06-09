package main

import (
	"log"
	"net/http"

	pc "Simplego/Pack"
)

func main() {
	//ch := make(chan float64)
	//http.HandleFunc("/getinsert", pc.GetInsert)
	http.HandleFunc("/getPost", pc.GetPost)
	http.HandleFunc("/insert", pc.Insert)

	http.HandleFunc("/Post", pc.Post)
	// http.HandleFunc("/edit/", editHandler)
	// http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
