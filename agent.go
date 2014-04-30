package main

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
)

func agentTokenList(c *cli.Context) {
	tokens, err := GetClient().AgentTokenList()
	if err != nil {
		log.Fatal(err)
	}
	Display(tokens)
}

func agentHostInfo(c *cli.Context) {
	agId := c.String("agent-id")
	if len(agId) == 0 {
		fmt.Println("Agent ID Missing")
		return
	}
	agentType := c.String("type")
	if len(agentType) == 0 {
		fmt.Println("Type Missing")
		return
	}

	info, err := GetClient().AgentHostInfo(agId, agentType)
	if err != nil {
		log.Fatal(err)
	}
	Display(info)
}

func agentTokenDelete(c *cli.Context) {
	id := c.String("id")
	if len(id) == 0 {
		fmt.Println("ID Missing")
		return
	}
	err := GetClient().DeleteAgentToken(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Token Deleted")
}

func agentUpgrade(c *cli.Context) {
	id := c.String("agent-id")
	if len(id) == 0 {
		fmt.Println("agent-id Missing")
		return
	}
	err := GetClient().UpgradeAgent(id)
	if err != nil {
		log.Fatal(err)
	}
}

func agentConnectionsList(c *cli.Context) {
	agId := c.String("agent-id")
	if len(agId) == 0 {
		fmt.Println("Agent ID Missing")
		return
	}
	conns, err := GetClient().AgentConnectionsList(agId)
	if err != nil {
		log.Fatal(err)
	}
	Display(conns)
}

var AgentsExports []cli.Command = []cli.Command{
	{
		Name:   "agent_tokens.list",
		Usage:  "Agent Token List",
		Action: agentTokenList,
	},
	{
		Name:   "agent_tokens.delete",
		Usage:  "Agent Token Delete",
		Action: agentTokenDelete,
		Flags: []cli.Flag{
			cli.StringFlag{"id", "", "ID"},
		},
	},
	{
		Name:   "agent_host.info",
		Usage:  "Agent Host Info",
		Action: agentHostInfo,
		Flags: []cli.Flag{
			cli.StringFlag{"agent-id", "", "The Agent ID"},
			cli.StringFlag{"type", "", "Host Info Type"},
		},
	},
	{
		Name:   "agent_connections.list",
		Usage:  "Agent Connections List",
		Action: agentConnectionsList,
		Flags: []cli.Flag{
			cli.StringFlag{"agent-id", "", "The Agent ID"},
		},
	},
	{
		Name:   "agents.upgrade",
		Usage:  "Agent Upgrade",
		Action: agentUpgrade,
		Flags: []cli.Flag{
			cli.StringFlag{"agent-id", "", "The Agent ID"},
		},
	},
}
