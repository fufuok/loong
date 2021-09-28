package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/fufuok/xdaemon"
	"github.com/urfave/cli/v2"

	"github.com/fufuok/loong"
)

var (
	version = "v0.0.2.21092818"

	// 全局配置项
	conf = &loong.TConfig{}

	resetCmd = map[string]string{
		"iis":    "iisreset",
		"apache": "net stop ff.apachex64 & net start ff.apachex64",
		"test":   "echo Fufu*中文",
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "Daemon Web Server"
	app.Usage = "守护 Windows / Linux 的网站服务"
	app.UsageText = "- 请使用管理员身份运行\n   - 用于老旧边缘服务, 临时守护\n   - 支持 Windows / Linux, 可指定重启命令"
	app.Version = version
	app.Copyright = "https://github.com/fufuok/loong"
	app.Authors = []*cli.Author{
		{
			Name:  "Fufu",
			Email: "fufuok.com",
		},
	}
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "debug",
			Usage:       "调试模式",
			Aliases:     []string{"d"},
			Destination: &conf.Debug,
		},
		&cli.StringFlag{
			Name:        "log",
			Value:       "info",
			Usage:       "文件日志级别: debug, info, warn, error, fatal, panic",
			Aliases:     []string{"l"},
			Destination: &conf.LogLevel,
		},
		&cli.StringFlag{
			Name:        "logfile",
			Value:       loong.LogFile,
			Usage:       "日志文件位置",
			Destination: &conf.LogFile,
		},
		&cli.StringFlag{
			Name:        "errorlogfile",
			Value:       loong.ErrorLogFile,
			Usage:       "错误级别的日志文件位置",
			Destination: &conf.ErrorLogFile,
		},
		&cli.StringFlag{
			Name:        "url",
			Value:       "",
			Usage:       "检查的网址",
			Aliases:     []string{"u"},
			Required:    true,
			Destination: &conf.URL,
		},
		&cli.IntFlag{
			Name:        "status-code",
			Value:       200,
			Usage:       "网址需要返回的状态码",
			Aliases:     []string{"s"},
			Destination: &conf.StatusCode,
		},
		&cli.StringFlag{
			Name:        "text",
			Value:       "",
			Usage:       "指定网页要包含的文本",
			Aliases:     []string{"t"},
			Destination: &conf.ContainsText,
		},
		&cli.StringFlag{
			Name:    "web-service",
			Value:   "iis",
			Usage:   "Web 服务类型 (IIS / Apache)",
			Aliases: []string{"w"},
		},
		&cli.UintFlag{
			Name:    "interval",
			Value:   180,
			Usage:   "检查时间间隔 (秒)",
			Aliases: []string{"i"},
		},
		&cli.StringFlag{
			Name:    "cmd",
			Value:   "",
			Usage:   "指定重启命令 (优先使用)",
			Aliases: []string{"c"},
		},
	}
	app.Action = func(c *cli.Context) error {
		// 日志目录
		_ = os.Mkdir(loong.LogDir, os.ModePerm)

		// 守护自身
		if !conf.Debug {
			xdaemon.NewDaemon(loong.DaemonLogFile).Run()
		}

		// 重启命令配置
		cmd := c.String("cmd")
		service := strings.ToLower(c.String("web-service"))
		if cmd != "" {
			resetCmd[service] = cmd
		}

		// 初始化配置
		conf.WebService = service
		conf.ResetCmd = resetCmd
		conf.Interval = time.Duration(c.Uint("interval")) * time.Second
		loong.InitMain(conf)

		// 守护 Web 服务
		loong.Daemon()

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
