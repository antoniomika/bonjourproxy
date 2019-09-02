package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/oleksandr/bonjour"
)

var (
	serviceServers []*bonjour.Server
	servicesConfig = flag.String("services", "services.toml", "The config file defining services to proxy")
)

type service struct {
	Name        string
	ServiceType string
	Domain      string
	Port        int
	Host        string
	IP          string
	TextData    []string
}

type config struct {
	ProxyService []service
}

func main() {
	flag.Parse()

	var services config

	if _, err := toml.DecodeFile(*servicesConfig, &services); err != nil {
		log.Fatal(err)
	}

	for _, serviceServer := range services.ProxyService {
		log.Printf("Starting Bonjour Proxy for %+v", serviceServer)
		s, err := bonjour.RegisterProxy(serviceServer.Name, serviceServer.ServiceType, serviceServer.Domain, serviceServer.Port, serviceServer.Host, serviceServer.IP, serviceServer.TextData, nil)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Started Bonjour Proxy for %s", serviceServer.Name)

		serviceServers = append(serviceServers, s)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	for range c {
		for _, serviceServer := range serviceServers {
			serviceServer.Shutdown()
		}
		time.Sleep(1 * time.Second)

		log.Println("Stopped all Bonjour Proxies")

		os.Exit(0)
	}
}
