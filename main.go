package main

import (
	"log"

	"github.com/Pixium/MallData/pkg/cli"
	"github.com/Pixium/MallData/pkg/db"
)

func main() {
	_, err := db.SetupDatabase()
	if err != nil {
		log.Fatalf("unable to setup database due to %s", err.Error())
	}

	cli.StartSshServer()
}
