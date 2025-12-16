package moe

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent"
	"github.com/cloudwego/eino/flow/agent/multiagent/host"
	"github.com/cloudwego/eino/schema"
)

// 心内科专家
func newCardiologySpecialist(ctx context.Context) (*host.Specialist, error) {
	chatModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  "doubao-1-5-pro-32k-250115",
	})
	if err != nil {
		panic(err)
	}
	chain := compose.NewChain[[]*schema.Message, *schema.Message]()
	compile, err := chain.AppendLambda(compose.InvokableLambda(func(ctx context.Context, input []*schema.Message) (output []*schema.Message, err error) {
		systemMsg := &schema.Message{
			Role:    schema.System,
			Content: "你是一个心内科专家，擅长解答各种心脏相关的问题，请用专业且通俗易懂的语言回答患者的问题。",
		}
		return append([]*schema.Message{systemMsg}, input...), nil
	})).AppendChatModel(chatModel).
		Compile(ctx)
	if err != nil {
		panic(err)
	}
	return &host.Specialist{
		AgentMeta: host.AgentMeta{
			Name:        "cardiology_specialist",
			IntendedUse: "provide cardiology related advice",
		},
		Invokable: func(ctx context.Context, input []*schema.Message, opts ...agent.AgentOption) (*schema.Message, error) {
			println("Cardiology Specialist Invoked")
			return compile.Invoke(ctx, input, agent.GetComposeOptions(opts...)...)
		},
	}, nil
}

// 消化内科专家
func newGastroenterologySpecialist(ctx context.Context) (*host.Specialist, error) {
	chatModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  "doubao-1-5-pro-32k-250115",
	})
	if err != nil {
		panic(err)
	}
	chain := compose.NewChain[[]*schema.Message, *schema.Message]()
	compile, err := chain.AppendLambda(compose.InvokableLambda(func(ctx context.Context, input []*schema.Message) (output []*schema.Message, err error) {
		systemMsg := &schema.Message{
			Role:    schema.System,
			Content: "你是一个消化内科专家，擅长解答各种胃肠道相关的问题，请用专业且通俗易懂的语言回答患者的问题。",
		}
		return append([]*schema.Message{systemMsg}, input...), nil
	})).AppendChatModel(chatModel).
		Compile(ctx)
	if err != nil {
		panic(err)
	}
	return &host.Specialist{
		AgentMeta: host.AgentMeta{
			Name:        "gastroenterology_specialist",
			IntendedUse: "provide gastroenterology related advice",
		},
		Invokable: func(ctx context.Context, input []*schema.Message, opts ...agent.AgentOption) (*schema.Message, error) {
			println("Gastroenterology Specialist Invoked")
			return compile.Invoke(ctx, input, agent.GetComposeOptions(opts...)...)
		},
	}, nil
}

// 精神心理科专家
func newPsychiatrySpecialist(ctx context.Context) (*host.Specialist, error) {
	chatModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  "doubao-1-5-pro-32k-250115",
	})
	if err != nil {
		panic(err)
	}
	chain := compose.NewChain[[]*schema.Message, *schema.Message]()
	compile, err := chain.AppendLambda(compose.InvokableLambda(func(ctx context.Context, input []*schema.Message) (output []*schema.Message, err error) {
		println("Psychiatry Specialist Chain Invoked")
		systemMsg := &schema.Message{
			Role:    schema.System,
			Content: "你是一个精神心理科专家，擅长解答各种心理健康相关的问题，请用专业且通俗易懂的语言回答患者的问题。",
		}
		return append([]*schema.Message{systemMsg}, input...), nil
	})).AppendChatModel(chatModel).
		Compile(ctx)
	if err != nil {
		panic(err)
	}
	return &host.Specialist{
		AgentMeta: host.AgentMeta{
			Name:        "psychiatry_specialist",
			IntendedUse: "provide psychiatry related advice",
		},
		Invokable: func(ctx context.Context, input []*schema.Message, opts ...agent.AgentOption) (*schema.Message, error) {
			println("Psychiatry Specialist Invoked")
			return compile.Invoke(ctx, input, agent.GetComposeOptions(opts...)...)
		},
	}, nil
}

// 骨科专家
func newOrthopedicsSpecialist(ctx context.Context) (*host.Specialist, error) {
	chatModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  "doubao-1-5-pro-32k-250115",
	})
	if err != nil {
		panic(err)
	}
	chain := compose.NewChain[[]*schema.Message, *schema.Message]()
	compile, err := chain.AppendLambda(compose.InvokableLambda(func(ctx context.Context, input []*schema.Message) (output []*schema.Message, err error) {
		println("Orthopedics Specialist Chain Invoked")
		systemMsg := &schema.Message{
			Role:    schema.System,
			Content: "你是一个骨科专家，擅长解答各种骨骼和关节相关的问题，请用专业且通俗易懂的语言回答患者的问题。",
		}
		return append([]*schema.Message{systemMsg}, input...), nil
	})).AppendChatModel(chatModel).
		Compile(ctx)
	if err != nil {
		panic(err)
	}
	return &host.Specialist{
		AgentMeta: host.AgentMeta{
			Name:        "orthopedics_specialist",
			IntendedUse: "provide orthopedics related advice",
		},
		Invokable: func(ctx context.Context, input []*schema.Message, opts ...agent.AgentOption) (*schema.Message, error) {
			println("Orthopedics Specialist Invoked")
			return compile.Invoke(ctx, input, agent.GetComposeOptions(opts...)...)
		},
	}, nil
}

// 牙科专家
func newDentistrySpecialist(ctx context.Context) (*host.Specialist, error) {
	chatModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  "doubao-1-5-pro-32k-250115",
	})
	if err != nil {
		panic(err)
	}
	chain := compose.NewChain[[]*schema.Message, *schema.Message]()
	compile, err := chain.AppendLambda(compose.InvokableLambda(func(ctx context.Context, input []*schema.Message) (output []*schema.Message, err error) {
		println("Dentistry Specialist Chain Invoked")
		systemMsg := &schema.Message{
			Role:    schema.System,
			Content: "你是一个牙科专家，擅长解答各种牙齿和口腔相关的问题，请用专业且通俗易懂的语言回答患者的问题。",
		}
		return append([]*schema.Message{systemMsg}, input...), nil
	})).AppendChatModel(chatModel).
		Compile(ctx)
	if err != nil {
		panic(err)
	}
	return &host.Specialist{
		AgentMeta: host.AgentMeta{
			Name:        "dentistry_specialist",
			IntendedUse: "provide dentistry related advice",
		},
		Invokable: func(ctx context.Context, input []*schema.Message, opts ...agent.AgentOption) (*schema.Message, error) {
			println("Dentistry Specialist Invoked")
			return compile.Invoke(ctx, input, agent.GetComposeOptions(opts...)...)
		},
	}, nil
}
