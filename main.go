package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/oleksandr/bonjour"
)

func main() {
	s, err := bonjour.RegisterProxy("1115", "_hap._tcp", "", 1200, "1115", "172.31.0.151", []string{
		"MFG=ecobee Inc.",
		"c#=4",
		"ci=9",
		"ff=1",
		"id=44:61:32:CC:0D:71",
		"md=ecobee4",
		"pv=1.1",
		"s#=1",
		"serial_number=511881680099",
		"sf=1",
	}, nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Started Bonjour Proxy for thermostat device.")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	for range c {
		s.Shutdown()
		time.Sleep(1 * time.Second)
		os.Exit(0)
	}
}
