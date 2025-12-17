package adk

import (
	"context"
	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/schema"
	"os"
)

func WorkflowAgent() {
	ctx := context.Background()
	//初始化模型
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  "doubao-1-5-pro-32k-250115",
	})
	if err != nil {
		panic(err)
	}
	//初始化各个agent
	//情感分析agent
	sentimentAgent, err := adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Name:        "情感分析Agent",
		Description: "该Agent用于分析文本的情感倾向。",
		Instruction: "你是一个情感分析专家，能够准确地判断文本的情感倾向是积极、消极还是中立。",
		Model:       model,
	})
	if err != nil {
		panic(err)
	}
	//关键词提取agent
	keywordAgent, err := adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Name:        "关键词提取Agent",
		Description: "该Agent用于从文本中提取关键词。",
		Instruction: "你是一个关键词提取专家，能够从文本中提取出最重要的关键词。",
		Model:       model,
	})
	if err != nil {
		panic(err)
	}
	//翻译agent
	translationAgentTemp, err := adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Name:        "翻译Agent",
		Description: "该Agent用于将文本翻译为英文。",
		Instruction: "你是一个翻译专家，能够将各种语言的文本准确地翻译为英文。请翻译关键词为英文。",
		Model:       model,
	})
	if err != nil {
		panic(err)
	}
	translationAgent := adk.AgentWithOptions(ctx, translationAgentTemp, //重写历史记录处理函数，只保留关键词提取agent的回复作为输入，避免前面情感分析的内容干扰翻译结果
		adk.WithHistoryRewriter(func(ctx context.Context, entries []*adk.HistoryEntry) ([]adk.Message, error) {
			msg := make([]adk.Message, 0)
			for _, entry := range entries {
				if entry.AgentName != "关键词提取Agent" {
					continue
				}
				msg = append(msg, &schema.Message{
					Role:    schema.User,
					Content: entry.Message.Content,
				})
			}
			return msg, nil
		}))
	//初始化顺序工作流，用于提取关键词并且翻译
	sequentialAgent, err := adk.NewSequentialAgent(ctx, &adk.SequentialAgentConfig{
		Name:        "关键词提取翻译为英文",
		Description: "该工作流用于提取文本中的关键词，并将其翻译为英文。",
		SubAgents: []adk.Agent{
			keywordAgent,
			translationAgent,
		},
	})
	if err != nil {
		panic(err)
	}
	//初始化并发工作流，用于情感分析和关键词提取翻译
	parallelAgent, err := adk.NewParallelAgent(ctx, &adk.ParallelAgentConfig{
		Name:        "情感分析和关键词提取翻译",
		Description: "该工作流用于同时进行情感分析和关键词提取翻译。",
		SubAgents: []adk.Agent{
			sentimentAgent,
			sequentialAgent,
		},
	})
	if err != nil {
		panic(err)
	}

	//使用工作流agent
	run := parallelAgent.Run(ctx, &adk.AgentInput{
		Messages: []adk.Message{
			{
				Role:    "user",
				Content: "今天的清晨，我推开窗户，看见阳光柔和地洒在花园里，树叶在微风中摇曳，空气中弥漫着泥土的清香。这种宁静让我感到无比的满足和平静。然而，想到昨天发生的一些事情，心中不免又泛起了些许不快。朋友的冷漠眼神和那些未曾解释的误会带来了沉重的失落感，仿佛一块石头压在心头，久久无法释怀。许多次我尝试去沟通，可似乎没有任何改变，他那无声的疏离让我感到无力而沮丧。我真的很怀念小时候我们一起玩耍时那种无忧无虑的快乐，那时没有争吵，没有误解，只有笑声和纯粹的信任。或许生活就是这样，总会有令人欢欣的时刻，但也难免有灰暗的日子。希望下一次，我能够重新找到那份简单的快乐，或者至少能让心绪得以平稳。",
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
