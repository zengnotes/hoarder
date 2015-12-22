package main

import (
	"github.com/jcelliott/lumber"
	// "github.com/nanopack/hoarder/commands"
	"github.com/pagodabox/nanobox-config"
	"os"
	"strings"
)

var log lumber.Logger
var dataDir string

func main() {
	configFile := ""
	if len(os.Args) > 1 && !strings.HasPrefix(os.Args[1], "-") {
		configFile = os.Args[1]
	}

	conf := map[string]string{
		"listenAddr": ":1234",
		"logLevel":   "info",
		"token":      "token",
		"dataDir":    "/tmp/hoarder/",
	}

	config.Load(conf, configFile)
	conf = config.Config
	// do the stuff
	level := lumber.LvlInt(conf["logLevel"])
	log = lumber.NewConsoleLogger(level)
	log.Prefix("[hoarder]")
	dataDir = conf["dataDir"]

	err := Start(conf["listenAddr"], conf["token"])
	if err != nil {
		log.Error("could not start: %s", err.Error())
		os.Exit(1)
	}
}
