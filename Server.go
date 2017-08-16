package main

import (
	"fmt"

	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/preload", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Link", "<styles.css>; rel=preload; as=style")

		fmt.Fprint(w,"<html><div>test</div></html>")
	})

	http.HandleFunc("/postload", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Link", "<styles.css>; rel=preload; as=style")

		fmt.Fprint(w,"<html><div>test</div></html>")
	})

	log.Fatal(http.ListenAndServe(":9000", nil))

}