package rag

import (
	"context"
	"log"
	"sync"

	cli "github.com/milvus-io/milvus-sdk-go/v2/client"
)

/*
向量知识库客户端初始化
为了方便测试，这里通过docker-compose启动一个milvus服务
需要根据indexer_RAG.go里面的配置来创建一个collection和fields
这里通过docker-compose启动的milvus服务地址为localhost:19530
*/

var (
	milvusCli  cli.Client
	milvusOnce sync.Once
)

func getMilvusClient() cli.Client {
	milvusOnce.Do(func() {
		//初始化客户端
		ctx := context.Background()
		client, err := cli.NewClient(ctx, cli.Config{
			Address: "localhost:19530",
		})
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}
		milvusCli = client
	})
	return milvusCli
}
