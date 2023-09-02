package log

import (
	"example-api/config"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger
)

func Init() {

	levelsConfigurationMapping := map[string]logrus.Level{
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
	}

	Logger = &logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.JSONFormatter{
			DisableHTMLEscape: true,
		},
		Hooks: make(logrus.LevelHooks),
	}

	configuredLevel := config.GetLogLevel()
	logrusLevel, exists := levelsConfigurationMapping[strings.ToLower(configuredLevel)]
	if !exists {
		panic(fmt.Sprintf("log level %s not implemented", configuredLevel))
	}

	Logger.Level = logrusLevel

}
