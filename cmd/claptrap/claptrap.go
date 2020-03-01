package main

import (
	"flag"
	"log"

	"github.com/incidrthreat/Claptrap/claptrap"
)

func main() {
	var configFile = flag.String("config_file", "config.yaml", "config.yaml")
	flag.Parse()

	config := claptrap.NewConfig(*configFile)
	botServer, err := claptrap.NewBotServer(config)
	if err != nil {
		log.Fatalf("Error starting botserver: %s\n", err)
	}
	botServer.Start()
}
