package main

import (
	"github.com/rphillips/gorax/identity"
	"github.com/rphillips/gorax/monitoring"
)

const (
	MAAS_URL string = "https://monitoring.api.rackspacecloud.com/v1.0"
)

func GetClient() *monitoring.MonitoringClient {
	cli := monitoring.MakeAPIKeyMonitoringClient(MAAS_URL, identity.USIdentityService, Username, ApiKey)
	cli.SetDebug(Debug)
	return cli
}
