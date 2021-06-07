package main

import (
	"encoding/xml"
	"log"
	"os"
)

// 负责处理文档

type document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

func loadDocument(path string) ([]document, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("close file failed")
		}
	}(file)

	decoder := xml.NewDecoder(file)
	dump := struct {
		Documents []document `xml:"doc"`
	}{}
	err = decoder.Decode(&dump)
	if err != nil {
		return nil, err
	}
	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil

}
