package main

import (
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

var AgentsExports []cli.Command = []cli.Command{
	{
		Name:   "agent_tokens.list",
		Usage:  "Agent Token List",
		Action: agentTokenList,
	},
}
