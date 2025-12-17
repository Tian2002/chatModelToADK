package adk

import (
	"context"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
)

// AddTool 实现了一个简单的加法工具
type AddTool struct {
}

func GetAddTool() tool.InvokableTool {
	return &AddTool{}
}

type AddParam struct {
	A int `json:"a"`
	B int `json:"b"`
}

func (t *AddTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "add",
		Desc: "add two numbers",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"a": {
				Type:     schema.Integer,
				Desc:     "first number",
				Required: true,
			},
			"b": {
				Type:     schema.Integer,
				Desc:     "second number",
				Required: true,
			},
		}),
	}, nil
}

func (t *AddTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	p := &AddParam{}
	err := json.Unmarshal([]byte(argumentsInJSON), p)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", p.A-p.B), nil
}

// SubTool 实现了一个简单的减法工具
type SubTool struct{}

func GetSubTool() tool.InvokableTool {
	return &SubTool{}
}

type SubParam struct {
	A int `json:"a"`
	B int `json:"b"`
}

func (t *SubTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "sub",
		Desc: "sub two numbers",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"a": {
				Type:     schema.Integer,
				Desc:     "first number",
				Required: true,
			},
			"b": {
				Type:     schema.Integer,
				Desc:     "second number",
				Required: true,
			},
		}),
	}, nil
}

func (t *SubTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	p := &SubParam{}
	err := json.Unmarshal([]byte(argumentsInJSON), p)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", p.A-p.B), nil
}

// AnalyzeTool 实现了一个数学题内容难度分析工具
type AnalyzeTool struct {
}

func GetAnalyzeTool() tool.InvokableTool {
	return &AnalyzeTool{}
}

type AnalyzeParam struct {
	Content string `json:"content"`
}

func (a *AnalyzeTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "analyze",
		Desc: "analyze the difficulty of the content",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"content": {
				Type:     schema.String,
				Desc:     "content to be analyzed",
				Required: true,
			},
		}),
	}, nil
}
func (a *AnalyzeTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	// 解析输入参数
	p := &AnalyzeParam{}
	err := json.Unmarshal([]byte(argumentsInJSON), p)
	if err != nil {
		return "", err
	}
	//调用模型
	arkModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  "doubao-1-5-pro-32k-250115",
	})
	if err != nil {
		fmt.Printf("failed to create chat model: %v", err)
		return "", err
	}
	//调用模型
	AnalyzeInput := []*schema.Message{
		{
			Role:    schema.System,
			Content: "你是一个数学老师，你需要分析用户的问题，判断用户的问题的难度，难度分为简单，中等，困难，你需要根据用户的问题给出一个难度的评分，评分范围为1-10，1为简单，10为困难",
		},
		{
			Role:    schema.User,
			Content: p.Content,
		},
	}
	response, err := arkModel.Generate(ctx, AnalyzeInput)
	if err != nil {
		fmt.Printf("failed to generate: %v", err)
		return "", err
	}
	return response.Content, nil
}

// SentimentTool 文本情感分析工具
type SentimentTool struct{}

func GetSentimentTool() tool.InvokableTool {
	return &SentimentTool{}
}

type SentimentParam struct {
	Text string `json:"text"`
}

func (t *SentimentTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "sentiment_analysis",
		Desc: "analyze the sentiment of the text",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"text": {
				Type:     schema.String,
				Desc:     "text to be analyzed",
				Required: true,
			},
		}),
	}, nil
}

func (t *SentimentTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	p := &SentimentParam{}
	err := json.Unmarshal([]byte(argumentsInJSON), p)
	if err != nil {
		return "", err
	}
	//使用大模型进行情感分析
	arkModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  "doubao-1-5-pro-32k-250115",
	})
	if err != nil {
		fmt.Printf("failed to create chat model: %v", err)
		return "", err
	}
	//调用模型
	SentimentInput := []*schema.Message{
		{
			Role:    schema.System,
			Content: "你是一个情感分析专家，你需要分析用户的文本，判断用户的文本的情感，情感分为积极，中性，消极",
		},
		{
			Role:    schema.User,
			Content: p.Text,
		},
	}
	response, err := arkModel.Generate(ctx, SentimentInput)
	if err != nil {
		fmt.Printf("failed to generate: %v", err)
		return "", err
	}
	return response.Content, nil
}

// KeywordExtractionTool 文本关键词提取工具
type KeywordExtractionTool struct{}

func GetKeywordExtractionTool() tool.InvokableTool {
	return &KeywordExtractionTool{}
}

type KeywordExtractionParam struct {
	Text string `json:"text"`
}

func (t *KeywordExtractionTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "keyword_extraction",
		Desc: "extract keywords from the text",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"text": {
				Type:     schema.String,
				Desc:     "text to be analyzed",
				Required: true,
			},
		}),
	}, nil
}

func (t *KeywordExtractionTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	p := &KeywordExtractionParam{}
	err := json.Unmarshal([]byte(argumentsInJSON), p)
	if err != nil {
		return "", err
	}
	//使用大模型进行关键词提取
	arkModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  "doubao-1-5-pro-32k-250115",
	})
	if err != nil {
		fmt.Printf("failed to create chat model: %v", err)
		return "", err
	}
	//调用模型
	KeywordExtractionInput := []*schema.Message{
		{
			Role:    schema.System,
			Content: "你是一个关键词提取专家，你需要从用户的文本中提取出关键词，关键词之间用逗号分隔",
		},
		{
			Role:    schema.User,
			Content: p.Text,
		},
	}
	response, err := arkModel.Generate(ctx, KeywordExtractionInput)
	if err != nil {
		fmt.Printf("failed to generate: %v", err)
		return "", err
	}
	return response.Content, nil
}

// TranslationTool 文本翻译工具
type TranslationTool struct{}

func GetTranslationTool() tool.InvokableTool {
	return &TranslationTool{}
}

type TranslationParam struct {
	Text       string `json:"text"`
	TargetLang string `json:"target_lang"`
}

func (t *TranslationTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "translation",
		Desc: "translate text to target language",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"text": {
				Type:     schema.String,
				Desc:     "text to be translated",
				Required: true,
			},
			"target_lang": {
				Type:     schema.String,
				Desc:     "target language",
				Required: true,
			},
		}),
	}, nil
}

func (t *TranslationTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	p := &TranslationParam{}
	err := json.Unmarshal([]byte(argumentsInJSON), p)
	if err != nil {
		return "", err
	}
	//使用大模型进行翻译
	arkModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  "doubao-1-5-pro-32k-250115",
	})
	if err != nil {
		fmt.Printf("failed to create chat model: %v", err)
		return "", err
	}
	//调用模型
	TranslationInput := []*schema.Message{
		{
			Role:    schema.System,
			Content: fmt.Sprintf("你是一个翻译专家，你需要将用户的文本翻译成%s", p.TargetLang),
		},
		{
			Role:    schema.User,
			Content: p.Text,
		},
	}
	response, err := arkModel.Generate(ctx, TranslationInput)
	if err != nil {
		fmt.Printf("failed to generate: %v", err)
		return "", err
	}
	return response.Content, nil
}
