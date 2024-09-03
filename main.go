package main

import (
	"golang-training/config"
	"golang-training/config/clients"
	"golang-training/server"
)

func main() {
	config.Init()
	clients.New()

	sHTTP := server.NewHTTPServer()
	sHTTP.Start()
}
