package main

import (
	_ "embed"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed index.tmpl.html
var tmpl string

func main() {
	walk("./reading")

	tmpl, err := template.New("index.html").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	ff, err := os.OpenFile("index.html", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer ff.Close()

	err = tmpl.Execute(ff, container)
	if err != nil {
		panic(err)
	}

}

func walk(dir string) {
	fs, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range fs {
		name := f.Name()
		fullpath := filepath.Join(dir, f.Name())
		// dir
		if f.IsDir() {
			walk(fullpath)
		}

		// file
		// fullfile := filepath.Join(dir, name)
		lowername := strings.ToLower(name)
		if strings.HasSuffix(lowername, ".pdf") {
			log.Println(fullpath)
			container = append(container, &Book{
				Name: name,
				Path: fullpath,
			})

		}
	}
}

type Book struct {
	Name  string
	Path  string
	Image string
}

var container = make([]*Book, 0)
