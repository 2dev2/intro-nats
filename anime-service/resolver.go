package anime_service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type Resolver struct {
	animes   map[string]*Anime
	natsConn *nats.Conn
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func NewResolver(natsConn *nats.Conn) *Resolver {
	return &Resolver{
		animes:   make(map[string]*Anime),
		natsConn: natsConn,
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
	if len(r.animes) > 0 {
		for _, anime := range r.animes {
			animes = append(animes, anime)
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
