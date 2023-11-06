package main

import (
	"fmt"
	"log"

	"github.com/bogdanguranda/rest-service/db"
	"github.com/bogdanguranda/rest-service/service"
	"github.com/bogdanguranda/rest-service/util/config"
	"github.com/bogdanguranda/rest-service/util/logging"
	"github.com/bogdanguranda/rest-service/web"
)

func main() {
	configReader := config.JSONConfig{}
	cfg, err := configReader.ReadConfig("config.json")
	if err != nil {
		log.Fatalf(fmt.Sprintf("Error reading config.json: %s", err))
	}

	logger := logging.NewStdLogger(logging.LogLevel(cfg.LogLevel))

	fileDB := db.NewFileDB("input.txt")
	searcher := service.NewBinarySearch(fileDB, logger)

	web.StartServer(cfg.Port, searcher, logger)
}
