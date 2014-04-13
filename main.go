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

const (
	DEFAULT_CONFIG = ".raxrc"
)

func parseConfig(c *cli.Context) error {
	var ok bool

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	configFilePath := path.Join(usr.HomeDir, c.String("config"))

	// Load INI File
	config, err := ini.LoadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	Username, ok = config.Get("credentials", "username")
	if !ok {
		log.Fatal("username missing from credentials section")
	}

	ApiKey, ok = config.Get("credentials", "api_key")
	if !ok {
		log.Fatal("api_key missing from credentials section")
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "raxmon2"
	app.Flags = []cli.Flag{
		cli.StringFlag{"config", DEFAULT_CONFIG, ""},
	}
	app.Usage = "raxmon2 [command] [options]"
	app.Commands = append(EntitiesExports)
	app.Version = "0.0.1"
	app.Before = parseConfig
	app.Run(os.Args)
}
