package graph

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"os"
)

func CallGraph() {
	ctx := context.Background()
	// 初始化模型
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  "doubao-1-5-pro-32k-250115",
	})
	if err != nil {
		panic(err)
	}
	//编写节点
	lambda0 := compose.InvokableLambda(func(ctx context.Context, input string) (output []*schema.Message, err error) {
		output, err = prompt.FromMessages(schema.FString,
			schema.SystemMessage("{role}"),
			schema.MessagesPlaceholder("history_key", true),
			schema.UserMessage("{content}"),
		).Format(ctx, map[string]any{
			"role":    "你是一个智能助手,请你分析我今天的心情，如果是负面的，请回复2，否则回复1。",
			"content": input,
		})
		if err != nil {
			panic(err)
		}
		return output, nil
	})
	lambda1 := compose.InvokableLambda(func(ctx context.Context, input *schema.Message) (output string, err error) {
		//读取短期记忆内容
		i := ""
		err = compose.ProcessState(ctx, func(ctx context.Context, s map[string]any) error {
			in, _ := s["input"]
			i, _ = in.(string)
			return nil
		})
		if err != nil {
			panic(err)
		}
		return fmt.Sprintf("%s 去学习", i), nil
	})
	lambda2 := compose.InvokableLambda(func(ctx context.Context, input *schema.Message) (output string, err error) {
		//读取短期记忆内容
		i := ""
		err = compose.ProcessState(ctx, func(ctx context.Context, s map[string]any) error {
			in, _ := s["input"]
			i, _ = in.(string)
			return nil
		})
		if err != nil {
			panic(err)
		}
		return fmt.Sprintf("%s 去休息", i), nil
	})
	//短期记忆节点
	graph := compose.NewGraph[string, string](
		compose.WithGenLocalState[map[string]any](func(ctx context.Context) (state map[string]any) {
			return make(map[string]any)
		}),
	)
	//构建图-添加节点
	err = graph.AddLambdaNode("Lambda0", lambda0,
		compose.WithStatePreHandler(func(ctx context.Context, in string, state map[string]any) (string, error) { //短期记忆写入输入
			state["input"] = in
			return in, nil
		}))
	if err != nil {
		panic(err)
	}
	err = graph.AddChatModelNode("ChatModel", model)
	if err != nil {
		panic(err)
	}
	err = graph.AddLambdaNode("Lambda1", lambda1)
	if err != nil {
		panic(err)
	}
	err = graph.AddLambdaNode("Lambda2", lambda2)
	if err != nil {
		panic(err)
	}
	//构建图-添加边
	err = graph.AddEdge(compose.START, "Lambda0")
	if err != nil {
		panic(err)
	}
	err = graph.AddEdge("Lambda0", "ChatModel")
	if err != nil {
		panic(err)
	}
	err = graph.AddBranch("ChatModel", compose.NewGraphBranch(func(ctx context.Context, in *schema.Message) (endNode string, err error) {
		switch in.Content {
		case "1":
			return "Lambda1", nil
		default:
			return "Lambda2", nil
		}
	}, map[string]bool{"Lambda1": true, "Lambda2": true}))
	if err != nil {
		panic(err)
	}
	err = graph.AddEdge("Lambda1", compose.END)
	if err != nil {
		panic(err)
	}
	err = graph.AddEdge("Lambda2", compose.END)
	if err != nil {
		panic(err)
	}
	//编译图
	runner, err := graph.Compile(ctx)
	if err != nil {
		panic(err)
	}
	//执行图
	result, err := runner.Invoke(ctx, "我今天很累",
		compose.WithCallbacks(callbacks.NewHandlerBuilder(). //切面注入回调
			OnStartFn(func(ctx context.Context, info *callbacks.RunInfo, input callbacks.CallbackInput) context.Context {
				println(info.Name, "执行开始，输入：", fmt.Sprint(input))
				return ctx
			}).
			OnEndFn(func(ctx context.Context, info *callbacks.RunInfo, output callbacks.CallbackOutput) context.Context {
				println(info.Name, "执行结束，输出：", fmt.Sprint(output))
				return ctx
			}).
			Build()))
	if err != nil {
		panic(err)
	}
	println("最终结果:", result)

}
