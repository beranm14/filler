package main

// From https://gobyexample.com/http-servers

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func hi(w http.ResponseWriter, req *http.Request) {
	b, _ := ioutil.ReadFile("./filler.jpg")
	base64.StdEncoding.EncodeToString(b)
	fmt.Fprintf(w, "<!DOCTYPE html><html lang=\"en\">\n")
	fmt.Fprintf(w, "<body>\n")
	fmt.Fprintf(w, "<img src=\"data:image/jpeg;base64,%v\">\n", base64.StdEncoding.EncodeToString(b))
	fmt.Fprintf(w, "</body>\n")
	fmt.Fprintf(w, "</html>\n")
	fmt.Println("hi called")
}

func main() {
	http.HandleFunc("/", hi)

	fmt.Println("Listening on port", os.Getenv("PORT"))
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
	if err != nil {
		panic(err)
	}
}
