package main

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
)

func entityList(c *cli.Context) {

	entities, err := GetClient().ListEntities()
	if err != nil {
		log.Fatal(err)
	}

	Display(entities)
}

func entityGet(c *cli.Context) {
	enId := c.String("entity-id")
	if len(enId) == 0 {
		fmt.Println("Entity ID Missing")
		return
	}
	entity, err := GetClient().GetEntity(enId)
	if err != nil {
		log.Fatal(err)
	}
	Display(entity)
}

func entityDelete(c *cli.Context) {
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

func entityHostInfo(c *cli.Context) {
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

	Display(hostinfo)
}

func entityAgentTargets(c *cli.Context) {
	enId := c.String("entity-id")
	if len(enId) == 0 {
		fmt.Println("Entity ID Missing")
		return
	}
	agentType := c.String("type")
	if len(agentType) == 0 {
		fmt.Println("Type Missing")
		return
	}

	info, err := GetClient().AgentTargets(enId, agentType)
	if err != nil {
		log.Fatal(err)
	}

	Display(info)
}

var EntitiesExports []cli.Command = []cli.Command{
	{
		Name:   "entities.list",
		Usage:  "Entity List",
		Action: entityList,
	},
	{
		Name:   "entities.get",
		Usage:  "Entity Get",
		Action: entityGet,
		Flags: []cli.Flag{
			cli.StringFlag{"entity-id", "", "The Entity ID"},
		},
	},
	{
		Name:   "entities.delete",
		Usage:  "Entity Delete",
		Action: entityDelete,
		Flags: []cli.Flag{
			cli.StringFlag{"entity-id", "", "The Entity ID"},
		},
	},
	{
		Name:   "entities.host_info",
		Usage:  "Entity Host Info",
		Action: entityHostInfo,
		Flags: []cli.Flag{
			cli.StringFlag{"entity-id", "", "The Entity ID"},
			cli.StringFlag{"type", "", "Host Info Type"},
		},
	},
	{
		Name:   "entities.agent_targets",
		Usage:  "Entity Agent Targets",
		Action: entityAgentTargets,
		Flags: []cli.Flag{
			cli.StringFlag{"entity-id", "", "The Entity ID"},
			cli.StringFlag{"type", "", "Host Info Type"},
		},
	},
}
