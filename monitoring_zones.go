package main

import (
	"log"

	"github.com/codegangsta/cli"
)

func monitoringZoneList(c *cli.Context) {
	zones, err := GetClient().ListMonitoringZones()
	if err != nil {
		log.Fatal(err)
	}

	Display(zones)
}

var MonitoringZonesExports []cli.Command = []cli.Command{
	{
		Name:   "monitoring_zones.list",
		Usage:  "Monitoring Zones List",
		Action: monitoringZoneList,
	},
}
