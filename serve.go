package main

import (
	"log"
	"net/http"
)

type fileSystemWithFallback struct {
	http.FileSystem
}

func (fs fileSystemWithFallback) Open(name string) (http.File, error) {
	file, err := fs.FileSystem.Open(name)
	if err != nil {
		// if the file is not found, fall back to index.html
		file, err := fs.FileSystem.Open("index.html")
		if err != nil {
			return nil, err
		}

		return file, nil
	}

	return file, nil
}

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.FileServer(fileSystemWithFallback{http.Dir("build")})))
}
