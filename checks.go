package main

import (
	"fmt"

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
		Die(err)
	}
	Display(checks)
}

func checkCreate(c *cli.Context) {
	check := monitoring.CheckCreateStruct{}

	enId := c.String("entity-id")
	if len(enId) == 0 {
		Die("Entity ID required")
	}

	label := c.String("label")
	if len(label) > 0 {
		check.Label = &label
	}

	checkType := c.String("type")
	if len(checkType) == 0 {
		Die("Check type required")
	} else {
		check.Type = &checkType
	}

	timeout := c.Int("timeout")
	if timeout == -1 {
		Die("Timeout required")
	}

	period := c.Int("period")
	if period == -1 {
		Die("Period required")
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

	metadata_str := c.String("metadata")
	if len(metadata_str) > 0 {
		metadata_obj := StringToDict(metadata_str)
		check.Metadata = &metadata_obj
	}

	check.Timeout = &timeout
	check.Period = &period

	url, err := GetClient().CreateCheck(enId, &check)
	if err != nil {
		Die(err)
	}
	fmt.Println(url)
}

func checkDelete(c *cli.Context) {
	enId := c.String("entity-id")
	chId := c.String("id")

	if len(enId) == 0 {
		Die("Entity ID Missing")
	}
	if len(chId) == 0 {
		Die("Check ID Missing")
	}

	err := GetClient().DeleteCheck(enId, chId)
	if err != nil {
		Die(err)
	}
	fmt.Printf("%s removed\n", chId)
}

func checkEnableDisable(c *cli.Context, enable bool) {
	var operation string

	enId := c.String("entity-id")
	chId := c.String("id")

	if len(enId) == 0 {
		Die("Entity ID Missing")
	}
	if len(chId) == 0 {
		Die("Check ID Missing")
	}

	if enable {
		operation = "disabled"
	} else {
		operation = "enabled"
	}

	check := struct {
		Disabled bool `json:"disabled"`
	}{
		enable,
	}

	err := GetClient().UpdateCheck(enId, chId, &check)
	if err != nil {
		Die(err)
	}
	fmt.Printf("%s %s\n", chId, operation)
}

func checkEnable(c *cli.Context) {
	checkEnableDisable(c, false)
}

func checkDisable(c *cli.Context) {
	checkEnableDisable(c, true)
}

func checkUpdate(c *cli.Context) {
	check := make(map[string]interface{})

	enId := c.String("entity-id")
	chId := c.String("id")

	if len(enId) == 0 {
		Die("Entity ID Missing")
	}

	if len(chId) == 0 {
		Die("Check ID Missing")
	}

	label := c.String("label")
	if len(label) > 0 {
		check["label"] = &label
	}

	timeout := c.Int("timeout")
	if timeout != -1 {
		check["timeout"] = &timeout
	}

	period := c.Int("period")
	if period != -1 {
		check["period"] = &period
	}

	details_str := c.String("details")
	if len(details_str) > 0 {
		details_obj := StringToDict(details_str)
		check["details"] = &details_obj
	}

	monitoring_zones_str := c.String("monitoring-zones")
	if len(monitoring_zones_str) > 0 {
		check["monitoring_zones_poll"] = StringToList(monitoring_zones_str)
	}

	target_alias := c.String("target-alias")
	if len(target_alias) > 0 {
		check["target_alias"] = &target_alias
	}

	target_hostname := c.String("target-hostname")
	if len(target_hostname) > 0 {
		check["target_hostname"] = &target_hostname
	}

	target_resolver := c.String("target-resolver")
	if len(target_resolver) > 0 {
		check["target_resolver"] = &target_resolver
	}

	metadata_str := c.String("metadata")
	if len(metadata_str) > 0 {
		metadata_obj := StringToDict(metadata_str)
		check["metadata"] = &metadata_obj
	}

	err := GetClient().UpdateCheck(enId, chId, &check)
	if err != nil {
		Die(err)
	}
	fmt.Printf("%s updated\n", chId)
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
			cli.StringFlag{"metadata", "", ""},
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
		Name:   "checks.delete",
		Usage:  "Check Delete",
		Action: checkDelete,
		Flags: []cli.Flag{
			cli.StringFlag{"entity-id", "", "The Entity ID"},
			cli.StringFlag{"id", "", "The Check ID"},
		},
	},
	{
		Name:   "checks.disable",
		Usage:  "Check Disable",
		Action: checkDisable,
		Flags: []cli.Flag{
			cli.StringFlag{"entity-id", "", "The Entity ID"},
			cli.StringFlag{"id", "", "The Check ID"},
		},
	},
	{
		Name:   "checks.enable",
		Usage:  "Check Enable",
		Action: checkEnable,
		Flags: []cli.Flag{
			cli.StringFlag{"entity-id", "", "The Entity ID"},
			cli.StringFlag{"id", "", "The Check ID"},
		},
	},
	{
		Name:   "checks.update",
		Usage:  "Check Update",
		Action: checkUpdate,
		Flags: []cli.Flag{
			cli.StringFlag{"entity-id", "", "The Entity ID"},
			cli.StringFlag{"id", "", "The Check ID"},
			cli.StringFlag{"label", "", ""},
			cli.StringFlag{"details", "", ""},
			cli.StringFlag{"monitoring-zones", "", ""},
			cli.StringFlag{"target-alias", "", ""},
			cli.StringFlag{"target-hostname", "", ""},
			cli.StringFlag{"target-resolver", "", ""},
			cli.StringFlag{"metadata", "", ""},
			cli.IntFlag{"timeout", -1, ""},
			cli.IntFlag{"period", -1, ""},
		},
	},
}
