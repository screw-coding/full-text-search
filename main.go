package main

import (
	"full-text-search/store"
	"log"
	"time"
)

func main() {
	start := time.Now()
	log.Println("starting...")
	//从文件中载入索引数据
	myStore, err2 := store.OpenStore("./data/file")
	if err2 != nil {
		log.Fatal(err2)
	}

	start = time.Now()
	idx := make(index)
	//索引遍历数据,写入内存
	err := myStore.ForEach(func(k, v []byte) error {
		return idx.addIndexedData(k, v)
	})
	if err != nil {
		log.Fatal("load index data from disk failed")
	}
	log.Println("start to log index data")
	for k, v := range idx {
		log.Printf("key:%s,value:%d", k, v)
	}

	////载入文档
	documents, err := loadDocument("test.xml")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d doc in %v", len(documents), time.Since(start))
	//创建索引文件,内存中
	start = time.Now()
	//idx.add(documents)
	//
	////索引数据写入磁盘
	//for k, v := range idx {
	//	s := strings.Replace(strings.Trim(fmt.Sprint(v), "[]"), " ", ",", -1)
	//	log.Printf("key:%s,value:%s\n", k, s)
	//	_ = myStore.Set([]byte(k), []byte(s))
	//}
	//
	//log.Printf("Indexed %d documents in %v", len(documents), time.Since(start))
	//搜索
	start = time.Now()
	matchedIDs := idx.search("dog")
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))
	log.Printf("MatchedIDS:%d", matchedIDs)
	for _, id := range matchedIDs {
		doc := documents[id]
		log.Printf("id=%d,content=%s\n", id, doc.Text)
	}
}
