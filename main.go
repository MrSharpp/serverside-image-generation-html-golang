package main

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/xml"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type ExampleXml struct {
	XMLName xml.Name `xml:"examle"`
	One     string   `xml:"one,attr"`
	Two     string   `xml:"two,attr"`
}

type Param struct {
	Name string
	Age  int
}

type Image struct {
	Mimetype string
	Buffer   string
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Welcome, visit sub oages now."))
	})

	r.HandleFunc("/html/{message}", func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)

		// read image file as buffer
		f, err := os.Open("./Letter_A.jpg")
		defer f.Close()
		if err != nil {
			log.Fatal(err)
			w.Write([]byte("500 internal server error : reading image error"))
			return
		}
		bufferData := bytes.NewBuffer(nil)
		io.Copy(bufferData, f)
		bufferBytes := b64.StdEncoding.EncodeToString([]byte(bufferData.Bytes()))
		if err != nil {
			log.Fatal(err)
			w.Write([]byte("500 internal server error"))
			return
		}
		I := Image{Buffer: bufferBytes, Mimetype: "image/jpeg"}
		template, err := template.ParseFiles("./templates/" + params["message"] + ".html")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("file not found"))
			return
		}
		err = template.Execute(w, I)
		if err != nil {
			log.Fatal(err)
		}
	})

	http.ListenAndServe("127.0.0.1:3000", r)
}
