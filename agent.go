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

var AgentsExports []cli.Command = []cli.Command{
	{
		Name:   "agent_tokens.list",
		Usage:  "Agent Token List",
		Action: agentTokenList,
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
}
