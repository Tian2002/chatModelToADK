package chain

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

func CallChain() {
	ctx := context.Background()
	//准备模型
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  "doubao-1-5-pro-32k-250115",
	})
	if err != nil {
		panic(err)
	}
	//准备链式调用的 Lambda 节点，这里使用了一个简单的 Prompt 模板节点
	lambda := compose.InvokableLambda(func(ctx context.Context, input map[string]any) (output []*schema.Message, err error) {
		template := prompt.FromMessages(schema.FString,
			schema.SystemMessage("你是一个{role}"),
			schema.MessagesPlaceholder("history_key", true), // 插入对话历史，optional=true表示没有对话历史时会被忽略
			&schema.Message{
				Role:    schema.User,
				Content: "请帮帮我，{task}",
			},
		)
		output, err = template.Format(ctx, input)
		if err != nil {
			panic(err)
		}
		return output, nil
	})
	//编译链式调用
	chain, err := compose.NewChain[map[string]any, *schema.Message]().
		AppendLambda(lambda).
		AppendChatModel(model).Compile(ctx)
	if err != nil {
		panic(err)
	}
	//调用链式调用
	output, err := chain.Invoke(ctx, map[string]any{
		"role": "智能机器人",
		"history_key": []*schema.Message{
			schema.UserMessage("你好"),
			schema.AssistantMessage("你好！很高兴见到你。有什么我可以帮你的吗？", nil),
		},
		"task": "raft算法我总是理解不了",
	})
	if err != nil {
		panic(err)
	}
	println(output.Content)
}
