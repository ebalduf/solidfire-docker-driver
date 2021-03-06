package main

import (
	"github.com/solidfire/solidfire-docker-driver/sfapi"
	"github.com/solidfire/solidfire-docker-driver/sfcli"
	"os"
)

const (
	VERSION = "1.3.1"
)

var (
	client *sfapi.Client
)

func main() {
	cli := sfcli.NewCli(VERSION)
	cli.Run(os.Args)

}
