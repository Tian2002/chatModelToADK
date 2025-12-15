package chain

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino-ext/components/tool/browseruse"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

type InputParams struct {
	Url string `json:"url"`
}

func GoToURL(_ context.Context, params *InputParams) (string, error) {
	browserUseTool, err := browseruse.NewBrowserUseTool(context.Background(), &browseruse.Config{})
	if err != nil {
		panic(err)
	}
	result, err := browserUseTool.Execute(&browseruse.Param{
		Action: browseruse.ActionGoToURL,
		URL:    &params.Url,
	})
	if err != nil {
		panic(err)
	}

	return result.Output, nil
}

func CreateGoToURLTool() tool.InvokableTool {
	goToURLTool := utils.NewTool(&schema.ToolInfo{
		Name: "go_to_url",
		Desc: "Use a browser to visit a specified URL and return the page content.",
		ParamsOneOf: schema.NewParamsOneOfByParams(
			map[string]*schema.ParameterInfo{
				"url": {
					Type:     schema.String,
					Desc:     "The URL of the webpage to visit.",
					Required: true,
				},
			},
		),
	}, GoToURL)
	return goToURLTool
}

func SimpleAgent() {
	ctx := context.Background()
	// 初始化模型
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  "doubao-1-5-pro-32k-250115",
	})
	if err != nil {
		panic(err)
	}
	//初始化工具
	but := CreateGoToURLTool()
	//编译链式调用
	toolInfo, err := but.Info(ctx)
	if err != nil {
		panic(err)
	}
	err = model.BindTools([]*schema.ToolInfo{toolInfo})
	if err != nil {
		panic(err)
	}
	toolNode, err := compose.NewToolNode(ctx, &compose.ToolsNodeConfig{
		Tools: []tool.BaseTool{but},
	})
	if err != nil {
		panic(err)
	}
	chain, err := compose.NewChain[[]*schema.Message, []*schema.Message]().
		AppendChatModel(model, compose.WithNodeName("chat_model")).
		AppendToolsNode(toolNode, compose.WithNodeName("tools")).
		Compile(ctx)
	if err != nil {
		panic(err)
	}
	//调用链式调用
	output, err := chain.Invoke(ctx, []*schema.Message{
		schema.UserMessage("访问一下https://www.baidu.com"),
	})
	if err != nil {
		panic(err)
	}
	println(output[0].Content)
}
