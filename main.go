package main

import (
	"log"
	"myz-torrent-api/server"
)

func main() {
	s := &server.Server{}

	log.Panic(s.Run())
}
