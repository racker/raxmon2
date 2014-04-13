package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/codegangsta/cli"
)

func display(obj interface{}) {
	str, _ := json.MarshalIndent(obj, "", "  ")
	fmt.Println(string(str))
}

func List(c *cli.Context) {

	entities, err := GetClient().ListEntities()
	if err != nil {
		log.Fatal(err)
	}

	display(entities)
}

func Get(c *cli.Context) {
	enId := c.String("entity-id")
	if len(enId) == 0 {
		fmt.Println("Entity ID Missing")
		return
	}
	entity, err := GetClient().GetEntity(enId)
	if err != nil {
		log.Fatal(err)
	}
	display(entity)
}

func Delete(c *cli.Context) {
	enId := c.String("entity-id")
	if len(enId) == 0 {
		fmt.Println("Entity ID Missing")
		return
	}
	_, err := GetClient().DeleteEntity(enId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Entity Deleted", enId)
}

func HostInfo(c *cli.Context) {
	enId := c.String("entity-id")
	if len(enId) == 0 {
		fmt.Println("Entity ID Missing")
		return
	}
	hostInfoType := c.String("type")
	if len(hostInfoType) == 0 {
		fmt.Println("Type Missing")
		return
	}

	hostinfo, err := GetClient().HostInfoEntity(enId, hostInfoType)
	if err != nil {
		log.Fatal(err)
	}

	display(hostinfo)
}

var EntitiesExports []cli.Command = []cli.Command{
	{
		Name:   "entities.list",
		Usage:  "Entity List",
		Action: List,
	},
	{
		Name:   "entities.get",
		Usage:  "Entity Get",
		Action: Get,
		Flags: []cli.Flag{
			cli.StringFlag{"entity-id", "", "The Entity ID"},
		},
	},
	{
		Name:   "entities.delete",
		Usage:  "Entity Delete",
		Action: Delete,
		Flags: []cli.Flag{
			cli.StringFlag{"entity-id", "", "The Entity ID"},
		},
	},
	{
		Name:   "entities.host_info",
		Usage:  "Entity Host Info",
		Action: HostInfo,
		Flags: []cli.Flag{
			cli.StringFlag{"entity-id", "", "The Entity ID"},
			cli.StringFlag{"type", "", "Host Info Type"},
		},
	},
}
