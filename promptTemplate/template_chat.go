package promptTemplate

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

func TemplateChat() {
	ctx := context.Background()
	template := prompt.FromMessages(schema.FString,
		schema.SystemMessage("你是一个{role}"),
		schema.MessagesPlaceholder("history_key", true), // 插入对话历史，optional=true表示没有对话历史时会被忽略
		&schema.Message{
			Role:    schema.User,
			Content: "请帮帮我，{task}",
		},
	)
	params := map[string]any{
		"role": "智能机器人",
		"history_key": []*schema.Message{
			schema.UserMessage("你好"),
			schema.AssistantMessage("你好！很高兴见到你。有什么我可以帮你的吗？", nil),
		},
		"task": "raft算法我总是理解不了",
	}
	messages, err := template.Format(ctx, params)
	if err != nil {
		panic(err)
	}

	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  os.Getenv("MODEL"),
	})
	if err != nil {
		panic(err)
	}

	answer, err := model.Generate(ctx, messages)
	if err != nil {
		panic(err)
	}

	println(answer.Content)
}
