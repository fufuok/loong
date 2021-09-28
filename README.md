# Loong (东方青龙, 守护 Web Service)

## 功能

- 周期检查指定网址的响应码和包含的文本
- 异常时自动执行重启服务命令

## 使用

`loong.exe -h`

`./loong -h`

```shell
NAME:
   Daemon Web Server - 守护 Windows / Linux 的网站服务

USAGE:
   - 请使用管理员身份运行
   - 用于老旧边缘服务, 临时守护
   - 支持 Windows / Linux, 可指定重启命令

VERSION:
   v0.0.2.21092818

AUTHOR:
   Fufu <fufuok.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -d                    调试模式 (default: false)
   --log value, -l value          文件日志级别: debug, info, warn, error, fatal, panic (default: "info")
   --logfile value                日志文件位置 (default: "log/loong.log")
   --errorlogfile value           错误级别的日志文件位置 (default: "log/loong.error.log")
   --url value, -u value          检查的网址
   --status-code value, -s value  网址需要返回的状态码 (default: 200)
   --text value, -t value         指定网页要包含的文本
   --web-service value, -w value  Web 服务类型 (IIS / Apache) (default: "iis")
   --interval value, -i value     检查时间间隔 (秒) (default: 180)
   --cmd value, -c value          指定重启命令 (优先使用)
   --help, -h                     show help (default: false)
   --version, -v                  print the version (default: false)

COPYRIGHT:
   https://github.com/fufuok/loong
```

## 示例

`loong.exe -d -u http://111.222 -w test -i 5`

`loong.exe -d -u https://www.baidu.com -i 30`

`./loong -u http://127.0.0.1:8080/ping -t PONG -i 5`





*ff*