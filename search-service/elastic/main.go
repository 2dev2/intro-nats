package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/olivere/elastic"
	"github.com/rubiagatra/search-service/elastic/model"
)

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"anime":{
			"properties":{
				"id":{
					"type":"keyword"
				},
				"title":{
					"type":"keyword"
				},
				"author":{
					"type":"keyword"
				}
			}
		}
	}
}`

func main() {
	ctx := context.Background()
	client, err := elastic.NewSimpleClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		panic(err)
	}

	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	exists, err := client.IndexExists("anime").Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		createIndex, err := client.CreateIndex("anime").BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
			log.Println("not acknowledged")
		}
	}

	anime1 := model.Anime{ID: "1", Title: "Boku no Hero Academia", Author: "Kohei Horikoshi"}
	put1, err := client.Index().
		Index("anime").
		Type("anime").
		Id("1").
		BodyJson(anime1).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed anime %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	time.Sleep(1000 * time.Millisecond)

	searchResult, err := client.Search().
		Index("anime").
		Type("anime").
		Sort("title", true).
		From(0).Size(10).
		Pretty(true).
		Do(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	var ttanime model.Anime
	for _, item := range searchResult.Each(reflect.TypeOf(ttanime)) {
		if a, ok := item.(model.Anime); ok {
			fmt.Printf("Anime Title %s: Author %s\n", a.Title, a.Author)
		}
	}
	fmt.Printf("Found a total of %d animes\n", searchResult.TotalHits())
	if searchResult.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d animes \n", searchResult.Hits.TotalHits)
		for _, hit := range searchResult.Hits.Hits {
			var a model.Anime
			err := json.Unmarshal(*hit.Source, &a)
			if err != nil {
			}
			fmt.Printf("Anime Title %s: Author %s\n", a.Title, a.Author)
		}
	} else {
		fmt.Print("Found no animes\n")
	}

}
