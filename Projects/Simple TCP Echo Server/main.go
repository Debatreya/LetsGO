package main

import (
	"TCP-echo-server/config"
	"TCP-echo-server/server"
	"flag" // flag package is used to parse command line arguments
	"log" // log package is used to log messages to the console
)

func setupFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "Host to listen on") // Hosts from which to accept connections, "0.0.0.0" to accept from all hosts
	flag.IntVar(&config.Port, "port", 7280, "Port to listen on") // Port to listen on
	flag.Parse()
}

func main() {
	setupFlags()
	log.Println("Starting server on", config.Host, ":", config.Port)
	server.RunSyncTCPServer()
}