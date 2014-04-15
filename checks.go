package main

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
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

func checkTypeList(c *cli.Context) {
	types, err := GetClient().CheckTypeList()
	if err != nil {
		log.Fatal(err)
	}
	Display(types)
}

var ChecksExports []cli.Command = []cli.Command{
	{
		Name:   "checks.list",
		Usage:  "Check List",
		Action: checkList,
		Flags: []cli.Flag{
			cli.StringFlag{"entity-id", "", "The Entity ID"},
		},
	},
	{
		Name:   "check_type.list",
		Usage:  "Check Type List",
		Action: checkTypeList,
	},
}
