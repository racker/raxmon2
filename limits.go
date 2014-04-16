package main

import (
	"log"

	"github.com/codegangsta/cli"
)

func limitsList(c *cli.Context) {
	limits, err := GetClient().ListLimits()
	if err != nil {
		log.Fatal(err)
	}

	Display(limits)
}

var LimitsExports []cli.Command = []cli.Command{
	{
		Name:   "limits.list",
		Usage:  "Limits List",
		Action: limitsList,
	},
}
