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
	var config string
	if c != nil {
		config = c.String("config")
	} else {
		config = DEFAULT_CONFIG
	}
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return path.Join(usr.HomeDir, config)
}

func parseConfig(c *cli.Context) error {
	configFilePath := getConfigFilePath(c)

	// Load INI File
	config, err := ini.LoadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
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
	app := cli.NewApp()
	app.Name = "raxmon2"
	app.Flags = []cli.Flag{
		cli.StringFlag{"config", getConfigFilePath(nil), ""},
		cli.BoolFlag{"debug", ""},
	}
	app.Usage = "raxmon2 [command] [options]"
	app.Commands = append(EntitiesExports)
	app.Version = "0.0.1"
	app.Before = parseConfig
	app.Run(os.Args)
}
