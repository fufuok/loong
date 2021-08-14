package loong

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/phuslu/log"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// 检查 Web 服务状态
func checkWebStatus() error {
	resp, err := req.R().Get(conf.URL)
	if err != nil || resp.StatusCode() != conf.StatusCode {
		return fmt.Errorf("[StatusCode] expected: %d, actual: %d, err: %v", conf.StatusCode, resp.StatusCode(), err)
	}

	if conf.ContainsText != "" && !strings.Contains(resp.String(), conf.ContainsText) {
		return fmt.Errorf("[Response] don't contain: %s", conf.ContainsText)
	}

	return nil
}

// 重启 Web 服务
// exec.Command("cmd", "/c", "netstat -an", "|", "find", "/c", "8888").CombinedOutput()
// chcp 65001 && echo "中文"
func resetWebService() (out string, err error) {
	cmd, ok := conf.ResetCmd[conf.WebService]
	if !ok {
		return "abort", fmt.Errorf("%s: can't find reset command", conf.WebService)
	}

	var b []byte

	if runtime.GOOS == "windows" {
		b, err = exec.Command("cmd", "/C", cmd).CombinedOutput()
		b, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(b)
	} else {
		b, err = exec.Command("bash", "-c", cmd).CombinedOutput()
	}

	out = string(b)

	return
}

// Daemon 守护 Web 服务
func Daemon() {
	var (
		err error
		out string
	)

	log.Info().Msgf("Loong start working...\nConfig: %+v", conf)

	for range time.Tick(conf.Interval) {
		if err = checkWebStatus(); err == nil {
			log.Info().Msg("OK")
			continue
		}

		log.Error().Err(err).Msg("Run checkWebStatus")

		for i := 0; i < retries; i++ {
			out, err = resetWebService()
			log.Warn().Str("out", out).Msg("Run resetWebService")
			if err == nil {
				log.Info().Msg("Successfully reset the web service")
				break
			}
			log.Error().Err(err).Msg("Failed to resetWebService")
		}
	}
}
