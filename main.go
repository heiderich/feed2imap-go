package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/Necoro/feed2imap-go/internal/config"
	"github.com/Necoro/feed2imap-go/internal/imap"
	"github.com/Necoro/feed2imap-go/internal/log"
	"github.com/Necoro/feed2imap-go/internal/parse"
)

var cfgFile = flag.String("f", "config.yml", "configuration file")
var verbose = flag.Bool("v", false, "enable verbose output")

func run() error {
	flag.Parse()
	log.SetDebug(*verbose)

	log.Print("Starting up...")

	log.Printf("Reading configuration file '%s'", *cfgFile)
	cfg, err := config.Load(*cfgFile)
	if err != nil {
		return err
	}

	parse.Parse(cfg.Feeds)

	imapUrl, err := url.Parse(cfg.GlobalConfig["target"].(string))
	if err != nil {
		return fmt.Errorf("parsing 'target': %w", err)
	}

	c, err := imap.Connect(imapUrl)
	if err != nil {
		return err
	}

	defer c.Disconnect()

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
