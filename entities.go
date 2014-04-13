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

func list(c *cli.Context) {

	entities, err := GetClient().ListEntities()
	if err != nil {
		log.Fatal(err)
	}

	display(entities)
}

func get(c *cli.Context) {
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

func del(c *cli.Context) {
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

func hostInfo(c *cli.Context) {
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

func agentTargets(c *cli.Context) {
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

	display(info)
}

var EntitiesExports []cli.Command = []cli.Command{
	{
		Name:   "entities.list",
		Usage:  "Entity List",
		Action: list,
	},
	{
		Name:   "entities.get",
		Usage:  "Entity Get",
		Action: get,
		Flags: []cli.Flag{
			cli.StringFlag{"entity-id", "", "The Entity ID"},
		},
	},
	{
		Name:   "entities.delete",
		Usage:  "Entity Delete",
		Action: del,
		Flags: []cli.Flag{
			cli.StringFlag{"entity-id", "", "The Entity ID"},
		},
	},
	{
		Name:   "entities.host_info",
		Usage:  "Entity Host Info",
		Action: hostInfo,
		Flags: []cli.Flag{
			cli.StringFlag{"entity-id", "", "The Entity ID"},
			cli.StringFlag{"type", "", "Host Info Type"},
		},
	},
	{
		Name:   "entities.agent_targets",
		Usage:  "Entity Agent Targets",
		Action: agentTargets,
		Flags: []cli.Flag{
			cli.StringFlag{"entity-id", "", "The Entity ID"},
			cli.StringFlag{"type", "", "Host Info Type"},
		},
	},
}
