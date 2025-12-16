package reAct

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
)

func BuildReAct() {
	ctx := context.Background()

	//初始化模型
	arkModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  "doubao-1-5-pro-32k-250115",
	})
	if err != nil {
		panic(err)
	}
	//准备工具
	addTool := GetAddTool()
	subTool := GetSubTool()
	analyzeTool := GetAnalyzeTool()

	//构建ReAct agent
	reAgent, err := react.NewAgent(ctx, &react.AgentConfig{
		ToolCallingModel: arkModel,
		ToolsConfig: compose.ToolsNodeConfig{
			Tools:               []tool.BaseTool{addTool, subTool, analyzeTool},
			ExecuteSequentially: false, //true表示顺序执行，false表示允许并行执行
		},
		MaxStep:       10,
		GraphName:     "reActGraph",
		ModelNodeName: "arkModel",
		ToolsNodeName: "toolsNode",
	})
	if err != nil {
		panic(err)
	}

	//运行agent
	result, err := reAgent.Generate(ctx, []*schema.Message{{
		Role: schema.User,
		Content: "请你帮我分析一下以下数学问题的难度，并计算出结果：\n" +
			"问题：一个人有15个苹果，他给了朋友7个苹果，后来又买了10个苹果，现在他有多少个苹果？",
	}},
		agent.WithComposeOptions(compose.WithCallbacks(callbacks.NewHandlerBuilder(). //回调打印具体的过程
												OnStartFn(func(ctx context.Context, info *callbacks.RunInfo, input callbacks.CallbackInput) context.Context {
				println(fmt.Sprintf("runInfo: %s, 输入内容: %s", info, input))
				return ctx
			}).
			OnEndFn(func(ctx context.Context, info *callbacks.RunInfo, output callbacks.CallbackOutput) context.Context {
				println(fmt.Sprintf("runInfo: %s, 输出内容: %s", info, output))
				return ctx
			}).Build())))
	if err != nil {
		panic(err)
	}

	println()
	println("最终结果：", result.Content)
}
