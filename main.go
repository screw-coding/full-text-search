package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()
	log.Println("starting...")
	//载入文档
	documents, err := loadDocument("test.xml")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d doc in %v", len(documents), time.Since(start))
	//创建索引文件,内存中
	start = time.Now()
	idx := make(index)
	idx.add(documents)
	log.Printf("Indexed %d documents in %v", len(documents), time.Since(start))
	//搜索
	start = time.Now()
	matchedIDs := idx.search("cat")
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))
	for _, id := range matchedIDs {
		doc := documents[id]
		log.Printf("id=%d,content=%s\n", id, doc.Text)
	}

}
