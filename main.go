package main

import (
	"log"
	"net/http"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("content-type", "text/html")
		dat, err := ioutil.ReadFile("404.html")
		if err != nil {
			w.Write(([]byte)("502 bad gateway"))
		} else {
			w.Write(dat)
		}
	})
	http.HandleFunc("/logo.png", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("content-type", "image/png")
		dat, err := ioutil.ReadFile("logo.png")
		if err != nil {
			w.Write(([]byte)("502 bad gateway"))
		} else {
			w.Write(dat)
		}
	})
	http.HandleFunc("/proximanova.woff", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("content-type", "application/font-woff")
		dat, err := ioutil.ReadFile("proximanova.woff")
		if err != nil {
			w.Write(([]byte)("502 bad gateway"))
		} else {
			w.Write(dat)
		}
	})

	log.Fatal(http.ListenAndServe(":9000", nil))
}

