package main

import (
	"fmt"
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

func monitoringZoneTraceroute(c *cli.Context) {
	id := c.String("id")
	if len(id) == 0 {
		fmt.Println("Monitoring Zone ID Missing")
		return
	}
	target := c.String("target")
	if len(target) == 0 {
		fmt.Println("Target Missing")
		return
	}
	resolver := c.String("target_resolver")

	route, err := GetClient().TracerouteMonitoringZone(id, target, resolver)
	if err != nil {
		log.Fatal(err)
	}

	Display(route)
}

var MonitoringZonesExports []cli.Command = []cli.Command{
	{
		Name:   "monitoring_zones.list",
		Usage:  "Monitoring Zones List",
		Action: monitoringZoneList,
	},
	{
		Name:   "monitoring_zones.traceroute",
		Usage:  "Monitoring Zones Traceroute",
		Action: monitoringZoneTraceroute,
		Flags: []cli.Flag{
			cli.StringFlag{"id", "", "The Monitoring Zone ID"},
			cli.StringFlag{"target", "", "target"},
			cli.StringFlag{"target_resolver", "IPv4", "Target Resolver (Default: IPv4)"},
		},
	},
}
