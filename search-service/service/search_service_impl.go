package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/rubiagatra/search-service/elastic/model"
	pb "github.com/rubiagatra/search-service/pb"
)

// SearchAnimes for demo purpose we don't use page and size
func (s *Service) SearchAnimes(ctx context.Context, req *pb.SearchRequest) (animes *pb.Animes, err error) {
	searchResult, err := s.elasticClient.Search().
		Index("anime").
		Type("anime").
		Sort("title", true).
		From(0).Size(10).
		Pretty(true).
		Do(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
	fmt.Printf("Found a total of %d animes\n", searchResult.TotalHits())

	var animeSlices []*pb.Anime
	if searchResult.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d animes \n", searchResult.Hits.TotalHits)
		for _, hit := range searchResult.Hits.Hits {
			var a model.Anime
			err := json.Unmarshal(*hit.Source, &a)
			if err != nil {
				log.Print(err)
				return nil, err
			}
			fmt.Printf("Anime Title %s: Author %s\n", a.Title, a.Author)
			animeSlices = append(animeSlices, &pb.Anime{
				Id:     a.ID,
				Title:  a.Title,
				Author: a.Author,
			})
		}
	} else {
		fmt.Print("Found no animes\n")
		return nil, nil
	}

	animes = &pb.Animes{
		Animes: animeSlices,
	}

	return
}
