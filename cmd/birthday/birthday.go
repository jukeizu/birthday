package main

import (
	"os"

	"github.com/jukeizu/birthday/subscribers/job"
	"github.com/shawntoffel/services-core/command"
	configreader "github.com/shawntoffel/services-core/config"
	"github.com/shawntoffel/services-core/logging"
)

var commandArgs command.CommandArgs

type Config struct {
	JobConfig job.Config
}

func init() {
	commandArgs = command.ParseArgs()
}

func main() {
	logger := logging.GetLogger("birthday", os.Stdout)

	config := Config{}

	err := configreader.ReadConfig(commandArgs.ConfigFile, &config)

	if err != nil {
		panic(err)
	}

	err = job.Start(logger, config.JobConfig)

	if err != nil {
		panic(err)
	}
}
