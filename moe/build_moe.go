package moe

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/flow/agent/multiagent/host"
	"github.com/cloudwego/eino/schema"
)

func BuildMOE() {
	ctx := context.Background()
	// 准备专家和路由模型
	h, err := newHost(ctx)
	if err != nil {
		panic(err)
	}
	cardiologySpecialist, err := newCardiologySpecialist(ctx)
	if err != nil {
		panic(err)
	}
	gastroenterologySpecialist, err := newGastroenterologySpecialist(ctx)
	if err != nil {
		panic(err)
	}
	psychiatrySpecialist, err := newPsychiatrySpecialist(ctx)
	if err != nil {
		panic(err)
	}
	orthopedicsSpecialist, err := newOrthopedicsSpecialist(ctx)
	if err != nil {
		panic(err)
	}
	dentistrySpecialist, err := newDentistrySpecialist(ctx)
	if err != nil {
		panic(err)
	}
	// 构建MOE
	hostMA, err := host.NewMultiAgent(ctx, &host.MultiAgentConfig{
		Host: *h,
		Specialists: []*host.Specialist{
			cardiologySpecialist,
			gastroenterologySpecialist,
			psychiatrySpecialist,
			orthopedicsSpecialist,
			dentistrySpecialist,
		},
	})
	if err != nil {
		panic(err)
	}
	msg := &schema.Message{
		Role: schema.User,
		Content: "病人病情描述:\n+" +
			`我最近两周总觉得胸口不舒服，像压着一样，有时候会心慌、心跳很快，偶尔还会有点喘不上气。
这种感觉经常在吃完饭、躺下或者加班熬夜后更明显，还会反酸、嗳气、喉咙有点烧。
最近压力也很大，晚上睡不好，发作时会很紧张，担心是不是心脏出问题。
没有明确发烧，也没有咳嗽，但偶尔左胸和上腹部有点隐隐疼。`,
	}
	out, err := hostMA.Generate(ctx, []*schema.Message{msg}, host.WithAgentCallbacks(&logCallback{}))
	if err != nil {
		panic(err)
	}

	println()
	println("最终结果：")
	println(out.Content)
}

type logCallback struct{}

func (l *logCallback) OnHandOff(ctx context.Context, info *host.HandOffInfo) context.Context {
	println(fmt.Sprintf("中间过程:%s", info))
	return ctx
}
