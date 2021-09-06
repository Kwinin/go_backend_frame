package log

import (
	logging "github.com/op/go-logging"
	"os"
	"strings"
)

const (
	CRITICAL logging.Level = iota
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

var format = logging.MustStringFormatter(`%{color}[%{level:.4s}] %{time:2006-01-02 15:04:05} %{module} [%{longfunc}] : %{message}%{color:reset}`)

var backendLeveled logging.LeveledBackend

func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled = logging.AddModuleLevel(backendFormatter)
	backendLeveled.SetLevel(DEBUG, "")
	logging.SetBackend(backendLeveled)
}

func GetLogger(module string) *logging.Logger {
	return logging.MustGetLogger(module)
}

func SetLogLevel(level string) {
	switch strings.ToUpper(level) {
	case "CRITICAL":
		backendLeveled.SetLevel(CRITICAL, "")
	case "ERROR":
		backendLeveled.SetLevel(ERROR, "")
	case "WARNING":
		backendLeveled.SetLevel(WARNING, "")
	case "NOTICE":
		backendLeveled.SetLevel(NOTICE, "")
	case "INFO":
		backendLeveled.SetLevel(INFO, "")
	case "DEBUG":
		backendLeveled.SetLevel(DEBUG, "")
	default:
		backendLeveled.SetLevel(INFO, "")
	}
}

//隐藏
func Secret(s string) interface{} {
	return logging.Redact(s)
}
