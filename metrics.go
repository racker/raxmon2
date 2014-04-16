package main

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
)

func metricsList(c *cli.Context) {
	enId := c.String("entity-id")
	if len(enId) == 0 {
		fmt.Println("Entity ID Missing")
		return
	}
	chId := c.String("check-id")
	if len(chId) == 0 {
		fmt.Println("Check ID Missing")
		return
	}

	metrics, err := GetClient().ListMetrics(enId, chId)
	if err != nil {
		log.Fatal(err)
	}

	Display(metrics)
}

var MetricsExports []cli.Command = []cli.Command{
	{
		Name:   "metrics.list",
		Usage:  "Metrics List",
		Action: metricsList,
		Flags: []cli.Flag{
			cli.StringFlag{"entity-id", "", "The Entity ID"},
			cli.StringFlag{"check-id", "", "The Check ID"},
		},
	},
}
