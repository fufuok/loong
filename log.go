package loong

import (
	"github.com/phuslu/log"
)

// go-resty 日志
type logger struct{}

func (l *logger) Errorf(format string, v ...interface{}) {
	log.Error().Msgf(format, v...)
}

func (l *logger) Warnf(format string, v ...interface{}) {
	log.Error().Msgf(format, v...)
}

func (l *logger) Debugf(format string, v ...interface{}) {
	log.Error().Msgf(format, v...)
}
