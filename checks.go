package main

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/rphillips/gorax/monitoring"
)

func checkList(c *cli.Context) {
	enId := c.String("entity-id")
	if len(enId) == 0 {
		fmt.Println("Entity ID Missing")
		return
	}
	checks, err := GetClient().ListChecks(enId)
	if err != nil {
		log.Fatal(err)
	}
	Display(checks)
}

func checkTypesList(c *cli.Context) {
	types, err := GetClient().CheckTypeList()
	if err != nil {
		log.Fatal(err)
	}
	Display(types)
}

func checkCreate(c *cli.Context) {
	check := monitoring.CheckCreateStruct{}

	enId := c.String("entity-id")
	if len(enId) == 0 {
		log.Fatal("Entity ID required")
	}

	label := c.String("label")
	if len(label) > 0 {
		check.Label = &label
	}

	checkType := c.String("type")
	if len(checkType) == 0 {
		log.Fatal("Check type required")
	} else {
		check.Type = &checkType
	}

	timeout := c.Int("timeout")
	if timeout == -1 {
		log.Fatal("Timeout required")
	}

	period := c.Int("period")
	if period == -1 {
		log.Fatal("Period required")
	}

	details_str := c.String("details")
	if len(details_str) > 0 {
		details_obj := StringToDict(details_str)
		check.Details = &details_obj
	}

	monitoring_zones_str := c.String("monitoring-zones")
	if len(monitoring_zones_str) > 0 {
		check.MonitoringZonesPoll = StringToList(monitoring_zones_str)
	}

	target_alias := c.String("target-alias")
	if len(target_alias) > 0 {
		check.TargetAlias = &target_alias
	}

	target_hostname := c.String("target-hostname")
	if len(target_hostname) > 0 {
		check.TargetHostname = &target_hostname
	}

	target_resolver := c.String("target-resolver")
	if len(target_resolver) > 0 {
		check.TargetResolver = &target_resolver
	}

	check.Timeout = &timeout
	check.Period = &period

	url, err := GetClient().CreateCheck(enId, &check)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}

var ChecksExports []cli.Command = []cli.Command{
	{
		Name:   "checks.create",
		Usage:  "Check Create",
		Action: checkCreate,
		Flags: []cli.Flag{
			cli.StringFlag{"entity-id", "", ""},
			cli.StringFlag{"label", "", ""},
			cli.StringFlag{"type", "", ""},
			cli.StringFlag{"details", "", ""},
			cli.StringFlag{"monitoring-zones", "", ""},
			cli.StringFlag{"target-alias", "", ""},
			cli.StringFlag{"target-hostname", "", ""},
			cli.StringFlag{"target-resolver", "", ""},
			cli.IntFlag{"timeout", -1, ""},
			cli.IntFlag{"period", -1, ""},
		},
	},
	{
		Name:   "checks.list",
		Usage:  "Check List",
		Action: checkList,
		Flags: []cli.Flag{
			cli.StringFlag{"entity-id", "", "The Entity ID"},
		},
	},
	{
		Name:   "check_types.list",
		Usage:  "Check Types List",
		Action: checkTypesList,
	},
}
