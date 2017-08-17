package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"net/http"
)

func main() {


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var path = strings.Replace(r.URL.Path, "/", "", 1)

		if path == "preload" {
			fmt.Println("preload")
			w.Header().Set("Link", "<styles.css>; rel=preload; as=style, <script.js>; rel=preload; as=script")

			fmt.Fprint(w, getIndex())
		}else if path == "postload" {
			fmt.Println("postload")

			fmt.Fprint(w, getIndex())
		}else {
			fmt.Println("default")
			fmt.Fprint(w, getResource(path))
		}

	})

	log.Fatal(http.ListenAndServe(":3000", nil))

}

func getIndex() string{
	return getResource("index.html")
}

func getResource(path string) string {
	dat, err := ioutil.ReadFile(path)
	check(err)
	return string(dat)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}