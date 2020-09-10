package main

import (
	"fmt"
	"log"
	"net/http"
	"server"
)

func main(){
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gothercises/urlshort",
		"/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := server.MapHandler(pathsToUrls, mux)

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	yamlHandler, err := server.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	err = http.ListenAndServe(":8080", yamlHandler)
	if err != nil {
		log.Fatal(err)
	}
}

func defaultMux() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello, world!")
}