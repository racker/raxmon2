package main

import (
	"log"
	"os"
	"os/user"
	"path"

	"github.com/codegangsta/cli"
	"github.com/vaughan0/go-ini"
)

var Username string
var ApiKey string
var Debug bool

const (
	DEFAULT_CONFIG = ".raxrc"
)

func getConfigFilePath(c *cli.Context) string {
	if c != nil {
		return c.String("config")
	}
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return path.Join(usr.HomeDir, DEFAULT_CONFIG)
}

func parseConfig(c *cli.Context) error {
	configFilePath := getConfigFilePath(c)

	// Load INI File
	config, err := ini.LoadFile(configFilePath)
	if err != nil {
		return err
	}

	var ok bool

	Username, ok = config.Get("credentials", "username")
	if !ok {
		log.Fatal("username missing from credentials section")
	}

	ApiKey, ok = config.Get("credentials", "api_key")
	if !ok {
		log.Fatal("api_key missing from credentials section")
	}

	Debug = c.Bool("debug")

	return nil
}

func main() {
	commands := make([]cli.Command, 1)
	commands = append(commands, AgentsExports...)
	commands = append(commands, ChecksExports...)
	commands = append(commands, EntitiesExports...)
	commands = append(commands, LimitsExports...)
	commands = append(commands, MetricsExports...)
	commands = append(commands, MonitoringZonesExports...)

	app := cli.NewApp()
	app.Name = "raxmon2"
	app.Flags = []cli.Flag{
		cli.StringFlag{"config", getConfigFilePath(nil), "RaxRC Config"},
		cli.BoolFlag{"debug", "Enable Debug"},
	}
	app.Usage = "raxmon2 [command] [options]"
	app.Commands = commands
	app.Version = "0.0.1"
	app.Before = parseConfig
	app.Run(os.Args)
}
