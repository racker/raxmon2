package main

import (
	"log"

	"github.com/codegangsta/cli"
)

func checkTypesList(c *cli.Context) {
	types, err := GetClient().CheckTypeList()
	if err != nil {
		log.Fatal(err)
	}
	Display(types)
}

var CheckTypesExports []cli.Command = []cli.Command{
	{
		Name:   "check_types.list",
		Usage:  "Check Types List",
		Action: checkTypesList,
	},
}
