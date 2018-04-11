package main

import (
	"os"
	"sync"

	"github.com/jukeizu/birthday/subscribers/jobs"
	base "github.com/jukeizu/client-base"
	"github.com/jukeizu/handler"
	mdb "github.com/shawntoffel/GoMongoDb"
	"github.com/shawntoffel/services-core/command"
	configreader "github.com/shawntoffel/services-core/config"
	"github.com/shawntoffel/services-core/logging"
)

var commandArgs command.CommandArgs

type Config struct {
	SchedulingClient base.ClientConfig
	JobConfig        job.Config
	DbConfig         mdb.DbConfig
	ServiceConfig    configreader.ServiceConfig
	CommandConfig    CommandConfig
}

type CommandConfig struct {
	BirthdayClient   base.ClientConfig
	SetHandlerConfig handler.HandlerConfig
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

	wg := sync.WaitGroup{}

	StartServices(&wg, logger, config)
	StartJobs(&wg, logger, config)
	StartCommands(&wg, logger, config.CommandConfig)

	wg.Wait()
}
