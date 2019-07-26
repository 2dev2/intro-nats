package anime_service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	pb "anime-service/pb"

	"github.com/nats-io/nats.go"
)

type Resolver struct {
	animes              map[string]*Anime
	natsConn            *nats.Conn
	searchServiceClient pb.SearchServiceClient
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func NewResolver(natsConn *nats.Conn, searchServiceClient pb.SearchServiceClient) *Resolver {
	animes := make(map[string]*Anime)
	animes["1"] = &Anime{ID: "1", Title: "Boku no Hero Academia", Author: "Kohei Horikoshi"}

	return &Resolver{
		animes:              animes,
		natsConn:            natsConn,
		searchServiceClient: searchServiceClient,
	}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateAnime(ctx context.Context, input InputAnime) (*Anime, error) {
	anime := &Anime{
		ID:     fmt.Sprintf("%d", time.Now().Unix()),
		Title:  input.Title,
		Author: input.Author,
	}

	r.animes[anime.ID] = anime

	animeJSON, err := json.Marshal(anime)
	if err != nil {
		log.Print("something wrong when marshaling JSON")
		return nil, err
	}

	go r.natsConn.Publish("anime-service-subject", animeJSON)

	return anime, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) FindAllAnimes(ctx context.Context, page int, size int) ([]*Anime, error) {
	var animes []*Anime
	response, err := r.searchServiceClient.SearchAnimes(ctx, &pb.SearchRequest{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if len(response.Animes) > 0 {
		for _, anime := range response.Animes {
			animes = append(animes, &Anime{ID: anime.Id, Title: anime.Title, Author: anime.Author})
		}
	}

	return animes, nil
}
func (r *queryResolver) FindAnimeByID(ctx context.Context, id string) (*Anime, error) {
	if _, ok := r.animes[id]; !ok {
		err := errors.New("anime not found")
		log.Print(err)
		return nil, err
	}
	return r.animes[id], nil
}
