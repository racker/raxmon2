package main

import (
	"github.com/racker/gorax/identity"
	"github.com/racker/gorax/monitoring"
)

const (
	MAAS_URL string = "https://monitoring.api.rackspacecloud.com/v1.0"
)

func GetClient() *monitoring.MonitoringClient {
	cli := monitoring.MakeAPIKeyMonitoringClient(MAAS_URL, identity.USIdentityService, Username, ApiKey)
	return cli
}
