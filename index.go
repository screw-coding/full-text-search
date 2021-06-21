package main

import "strings"

// index is an invented index.It maps tokens to document IDs
type index map[string][]int

func (idx index) addIndexedData(k []byte, v []byte) error {
	strDocIds := strings.Split(string(v), ",")

	var intDocIds []int
	for id := range strDocIds {
		intDocIds = append(intDocIds, id)
	}
	idx[string(k)] = intDocIds
	return nil
}

// 索引中加入原始文档数据,此函数会进行原始数据的处理
func (idx index) add(docs []document) {
	//遍历每个文档
	for _, doc := range docs {
		//遍历文档中的词(已处理的)
		for _, token := range analyze(doc.Text) {
			ids := idx[token]
			//todo: ?
			if ids != nil && ids[len(ids)-1] == doc.ID {
				// Don't add same ID twice 同样的id就不用添加两次了
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

// 把倒排索引的文档列表处理成升序的
func intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

func (idx index) search(text string) []int {
	var r []int
	for _, token := range analyze(text) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = intersection(r, ids)
			}
		} else {
			return nil
		}
	}
	return r
}
