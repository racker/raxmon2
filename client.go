package main

import (
	"github.com/rphillips/gorax/identity"
	"github.com/rphillips/gorax/monitoring"
)

var MAAS_URL string = "https://monitoring.api.rackspacecloud.com/v1.0"
var IDENTITY_URL string = identity.USIdentityService

func GetClient() *monitoring.MonitoringClient {
	cli := monitoring.MakeAPIKeyMonitoringClient(MAAS_URL, IDENTITY_URL, Username, ApiKey)
	cli.SetDebug(Debug)
	return cli
}
