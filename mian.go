package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load() // 加载环境变量
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	/*
		使用 chatModel 进行聊天生成,通常有流式和非流式两种调用方式
	*/

	//println("=============================chatModel.ChatGenerate=============================")
	//chatModel.ChatGenerate()
	//
	//println("=============================chatModel.ChatStream=============================")
	//chatModel.ChatStream()

	/*
		使用 promptTemplate 进行模板化聊天生成
	*/

	//println("=============================promptTemplate.TemplateChat=============================")
	//promptTemplate.TemplateChat()

	/*
		使用 tool 进行工具调用
	*/

	//println("=============================tool.CallTool=============================")
	//tool.CallTool()

	/*
		使用 rag 进行embedding、indexing、retrieval、transform 操作
	*/

	//println("=============================rag.EmbedText=============================")
	//rag.EmbedText()

	//println("=============================rag.IndexerRAG=============================")
	//rag.IndexerRAG([]*schema.Document{{
	//	ID:      "doc1",
	//	Content: "这是一个用于测试RAG功能的文档内容。",
	//	MetaData: map[string]any{
	//		"author": "Tian2002",
	//		"date":   "2025-12-01",
	//	},
	//},
	//	{
	//		ID:      "doc2",
	//		Content: "另一个测试文档，包含不同的信息以验证索引功能。",
	//		MetaData: map[string]any{
	//			"author": "AI助手",
	//			"date":   "2025-12-02",
	//		},
	//	},
	//	{
	//		ID:      "doc3",
	//		Content: "北京世纪好未来教育科技有限公司旗下的在线教育平台\n好未来是在线教育平台，前身为学而思，2013年8月更名为现名。该公司成立于2003年，2010年在美国纽交所上市（股票代码TAL），主营业务涵盖智慧教育、开放平台及K12课外教育，致力于通过科技推动教育创新",
	//		MetaData: map[string]any{
	//			"author": "教育专家",
	//			"date":   "2025-12-03",
	//		},
	//	},
	//})

	//println("=============================rag.RetrieverRAG=============================")
	//retrieverRAG := rag.RetrieverRAG("RAG")
	//for _, v := range retrieverRAG {
	//	println("检索到的内容：\n", v.Content)
	//}

	//println("=============================rag.TransDoc=============================")
	//rag.TransDoc()

	//println("=============================rag.BuildRAG=============================")
	//rag.BuildRAG()

	/*
		使用 chain 进行链式调用
	*/

	//println("=============================chain.CallChain=============================")
	//chain.CallChain()

	//println("=============================chain.SimpleAgent=============================")
	//chain.SimpleAgent()

	/*
		使用 graph 进行图结构调用
	*/

	//println("=============================graph.CallGraph=============================")
	//graph.CallGraph()

	/*
		使用reAct进行ReAct编排调用
	*/

	//println("=============================graph.OrcGraphWithState=============================")
	//reAct.BuildReAct()

	/*
		使用moe进行专家模型编排调用
	*/

	//println("=============================moe.BuildMOE=============================")
	//moe.BuildMOE()
}
