package rag

func BuildRAG() {
	docs := TransDoc()
	IndexerRAG(docs)
	results := RetrieverRAG("环境需求")
	println("检索结果如下：")
	for _, doc := range results {
		println(doc.ID)
		println("================================================")
		println(doc.Content)
	}
}
