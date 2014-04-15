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

var ChecksExports []cli.Command = []cli.Command{
	{
		Name:   "checks.list",
		Usage:  "Check List",
		Action: checkList,
		Flags: []cli.Flag{
			cli.StringFlag{"entity-id", "", "The Entity ID"},
		},
	},
}
