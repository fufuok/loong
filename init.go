package loong

import (
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/phuslu/log"
)

var (
	req     *resty.Client
	conf    *TConfig
	retries = 3
)

type TConfig struct {
	Debug        bool
	LogLevel     string
	URL          string
	StatusCode   int
	ContainsText string
	WebService   string
	Interval     time.Duration
	ResetCmd     map[string]string
}

func InitLoong(c *TConfig) {
	// 更新系统配置
	conf = c

	initLogger()
	initRequest()
}

func initLogger() {
	log.DefaultLogger = log.Logger{
		Level:      log.ParseLevel(conf.LogLevel),
		TimeFormat: "0102 15:04:05",
		Writer: &log.FileWriter{
			Filename:     "logs/loong.daemon.log",
			FileMode:     0600,
			MaxSize:      100 << 20,
			MaxBackups:   7,
			EnsureFolder: true,
			LocalTime:    true,
		},
	}
	if conf.Debug {
		log.DefaultLogger.Writer = &log.ConsoleWriter{
			ColorOutput:    true,
			QuoteString:    true,
			EndWithMessage: true,
		}
	}
}

func initRequest() {
	req = resty.New()
	req.SetDebug(conf.Debug)
	req.SetLogger(&logger{})
	req.SetRetryCount(retries)
	req.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 "+
		"(KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36")
}
