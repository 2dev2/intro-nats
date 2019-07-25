package anime_service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/labstack/gommon/log"
)

type Resolver struct {
	animes map[string]*Anime
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func NewResolver() *Resolver {
	return &Resolver{
		animes: make(map[string]*Anime),
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
		log.Error(err)
		return nil, err
	}
	return r.animes[id], nil
}
