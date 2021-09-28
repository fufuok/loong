package loong

import (
	"path/filepath"
	"time"

	"github.com/fufuok/utils"
	"github.com/go-resty/resty/v2"
)

var (
	// RootPath 运行绝对路径
	RootPath = utils.ExecutableDir(true)

	// LogDir 日志路径
	LogDir        = filepath.Join(RootPath, "log")
	LogFile       = filepath.Join(LogDir, APPName+".log")
	ErrorLogFile  = filepath.Join(LogDir, APPName+".error.log")
	DaemonLogFile = filepath.Join(LogDir, APPName+".daemon.log")

	req     *resty.Client
	conf    *TConfig
	retries = 3
)

type TConfig struct {
	Debug        bool
	LogLevel     string
	LogFile      string
	ErrorLogFile string
	URL          string
	StatusCode   int
	ContainsText string
	WebService   string
	Interval     time.Duration
	ResetCmd     map[string]string
}

func InitMain(c *TConfig) {
	// 更新系统配置
	conf = c

	initLogger()
	initRequest()
}

func initRequest() {
	req = resty.New()
	req.SetDebug(conf.Debug)
	req.SetLogger(&logger{})
	req.SetRetryCount(retries)
	req.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 "+
		"(KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36")
}
