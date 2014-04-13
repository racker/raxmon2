package main

import (
	"github.com/racker/gorax/monitoring"
)

const (
	MAAS_URL string = "https://monitoring.api.rackspacecloud.com/v1.0"
	AUTH_URL string = "https://identity.api.rackspacecloud.com/v2.0"
)

func GetClient() *monitoring.MonitoringClient {
	return monitoring.MakeAPIKeyMonitoringClient(MAAS_URL, AUTH_URL, Username, ApiKey)
}
