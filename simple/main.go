package main

import (
	"log"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {

		log.Fatalf("Error: %s", err)
	}
	nc.Subscribe("kum.bah", func(m *nats.Msg) {
		log.Printf("[Received kum.bah] %s", string(m.Data))
	})

	nc.Subscribe("kum.*", func(m *nats.Msg) {
		log.Printf("[Received kum.*] %s", string(m.Data))
	})

	for range time.NewTicker(1 * time.Second).C {
		nc.Publish("kum.bah", []byte("hello world"))
	}
	runtime.Goexit()
}
