package config

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

var (
	Config Configuration
)

type Configuration struct {
	Log struct {
		Level string
	}
	Server struct {
		Address string
		Port    int
	}
}

func Init(flags *pflag.FlagSet) (err error) {

	err = flags.Parse(os.Args[1:])
	if err != nil {
		err = fmt.Errorf("unable to parse command-line arguments: %w", err)
		return
	}

	Config = Configuration{}
	Config.Log.Level, err = flags.GetString("logLevel")
	if err != nil {
		err = fmt.Errorf("unable to parse logLevel flag: %w", err)
		return
	}
	Config.Server.Address, err = flags.GetString("addr")
	if err != nil {
		err = fmt.Errorf("unable to parse addr flag: %w", err)
		return
	}
	Config.Server.Port, err = flags.GetInt("port")
	if err != nil {
		err = fmt.Errorf("unable to parse port flag: %w", err)
		return
	}

	return

}

func GetLogLevel() (level string) { return Config.Log.Level }
