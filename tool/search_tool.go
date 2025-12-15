package tool

import (
	"context"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/schema"
)

/*
创建一个tool
*/

type InputParams struct {
	Name string `json:"name"`
}

func GetResource(_ context.Context, params *InputParams) (string, error) {
	//简单模拟返回资源url
	var count int
	for i := 0; i < len(params.Name); i++ {
		u := params.Name[i]
		count += int(u)
	}
	switch count % 3 {
	case 0:
		return "https://www.cloudwego.io/zh", nil
	case 1:
		return "https://www.cloudwego.io/zh/docs/eino", nil
	case 2:
		return "https://github.com/cloudwego/eino", nil
	}

	return "", nil
}

func CreateResourceTool() tool.InvokableTool {
	getResourceTool := utils.NewTool(&schema.ToolInfo{
		Name: "get_resource",
		Desc: "get a resource url by name",
		ParamsOneOf: schema.NewParamsOneOfByParams(
			map[string]*schema.ParameterInfo{
				"name": &schema.ParameterInfo{
					Type:     schema.String,
					Desc:     "resource's name",
					Required: true,
				},
			},
		),
	}, GetResource)
	return getResourceTool
}
