package main

import (
	"log"

	"github.com/pranayjoshi/p_meet.git/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}
