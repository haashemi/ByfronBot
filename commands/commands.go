package commands

import "byfron/pkg/bonbast"

type Commands struct {
	bonbastClient *bonbast.Client
}

func NewCommands(bonbastClient *bonbast.Client) *Commands {
	return &Commands{bonbastClient: bonbastClient}
}
