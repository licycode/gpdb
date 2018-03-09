package configutils

import (
	pb "gp_upgrade/idl"
)

type ClientAndHostname struct {
	Client   pb.AgentClient
	Hostname string
}

func GetClients() ([]ClientAndHostname, error) {
	return []ClientAndHostname{}, nil
}
