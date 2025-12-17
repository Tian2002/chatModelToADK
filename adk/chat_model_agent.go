package adk

import (
	"context"
	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"os"
)

func ChatModelAgent() {
	ctx := context.Background()
	// 初始化模型
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  "doubao-1-5-pro-32k-250115",
	})
	if err != nil {
		panic(err)
	}
	//初始化tools
	addTool := GetAddTool()
	subTool := GetSubTool()
	analyzeTool := GetAnalyzeTool()
	// 创建 ChatModelAgent
	agent, err := adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Name:        "幼儿园数学教师",
		Description: "幼儿园数学教师，能使用工具计算简单的加减法，并能分析数学问题",
		Instruction: "你是一个幼儿园数学教师，擅长使用加减法工具帮助小朋友解决简单的数学问题。当你遇到需要计算的数学问题时，请使用提供的工具进行计算，并将结果告诉小朋友。如果问题需要分析，请使用分析工具进行分析。",
		Model:       model,
		ToolsConfig: adk.ToolsConfig{
			ToolsNodeConfig: compose.ToolsNodeConfig{
				Tools: []tool.BaseTool{addTool, subTool, analyzeTool},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	// 使用 agent 进行对话或其他操作
	run := agent.Run(ctx, &adk.AgentInput{
		Messages: []adk.Message{
			{
				Role:    schema.User,
				Content: "小明有5个苹果，他又买了3个苹果，又给了其他小朋友2个苹果，现在他有多少个苹果？并且请判断这个题目的难度。",
			},
		},
		EnableStreaming: false,
	})
	for {
		event, ok := run.Next()
		if !ok {
			break
		}
		if event.Err != nil {
			panic(event.Err)
		}
		if event.Output != nil {
			println("输出角色：", event.Output.MessageOutput.Message.Role)
			println("输出内容：", event.Output.MessageOutput.Message.Content)
		}
	}
}
