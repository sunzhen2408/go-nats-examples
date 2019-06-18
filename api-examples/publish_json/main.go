package main

import (
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	// [begin publish_json]
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		nc.Close()
		log.Fatal(err)
	}
	defer ec.Close()

	// Define the object
	type stock struct {
		Symbol string
		Price  int
	}

	// Publish the message
	if err := ec.Publish("updates", &stock{Symbol: "GOOG", Price: 1200}); err != nil {
		log.Fatal(err)
	}
	// Make sure the message goes through before we close
	ec.Flush()
	// [end publish_json]
}
