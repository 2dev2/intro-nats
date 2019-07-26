package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
	"github.com/olivere/elastic"
	"github.com/rubiagatra/search-service/elastic/model"
)

func main() {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	ctx := context.Background()
	client, err := elastic.NewSimpleClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	nc.Subscribe("anime-service-subject", func(m *nats.Msg) {
		log.Printf("[Received anime-service-subject] %s", string(m.Data))
		var anime model.Anime
		err = json.Unmarshal(m.Data, &anime)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		put, err := client.Index().
			Index("anime").
			Type("anime").
			Id(anime.ID).
			BodyJson(anime).
			Do(ctx)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		fmt.Printf("Indexed anime %s to index %s, type %s\n", put.Id, put.Index, put.Type)

	})

	runtime.Goexit()
}
