package moe

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/flow/agent/multiagent/host"
)

func newHost(ctx context.Context) (*host.Host, error) {
	chatModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  "doubao-1-5-pro-32k-250115",
	})
	if err != nil {
		fmt.Printf("failed to create chat model: %v", err)
		return nil, err
	}

	return &host.Host{
		ToolCallingModel: chatModel,
		SystemPrompt:     "你是医院的多专家协调员。你有多个专家可以帮助你回答患者的问题。当患者提问时，你需要根据问题的性质选择一些合适的专家来提供帮助,提问时请向专家总结对应专家业务方向的描述。",
	}, nil
}
