package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func Execute(args []string) {
	newCmd().Execute(args)
}

type cmd struct {
	app *cli.App
}

func (c *cmd) Execute(args []string) {
	if err := c.app.Run(args); err != nil {
		fmt.Println("[Error]:", err)
		os.Exit(1)
	}
}

func newCmd() *cmd {
	version := RuntimeVersion
	app := &cli.App{}
	app.Name = "goRelease"
	app.Version = version
	app.Usage = "一个简单的agent客户端，用于采集主机信息并上报到服务端"
	app.Action = func(context *cli.Context) error {
		fmt.Println(app.Version)
		return nil
	}

	return &cmd{app: app}
}
