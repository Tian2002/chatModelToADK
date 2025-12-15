package tool

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudwego/eino-ext/components/tool/browseruse"
)

/*
调用一个tool
这里只是简单调用，实际使用中可以结合智能体agent来调用工具
*/

func CallTool() {
	ctx := context.Background()
	//自定义tool
	resourceUrl, err := CreateResourceTool().InvokableRun(ctx, `{"name":"测试1"}`) //获取资源url
	if err != nil {
		log.Fatal(err)
	}
	//调用浏览器工具访问该资源,官方示例工具
	but, err := browseruse.NewBrowserUseTool(ctx, &browseruse.Config{})
	if err != nil {
		log.Fatal(err)
	}
	result, err := but.Execute(&browseruse.Param{ //打开浏览器跳转到资源url
		Action: browseruse.ActionGoToURL,
		URL:    &resourceUrl,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	time.Sleep(10 * time.Second)
	but.Cleanup() //关闭浏览器
}
