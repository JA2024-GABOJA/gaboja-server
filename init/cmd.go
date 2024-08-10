package main

import (
	"junction/config"
	"junction/network"
)

type Cmd struct {
	config *config.Config
	router *network.Network

	//repository *repository.Repository
	//service    *service.Service
}
