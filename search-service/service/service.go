package service

import (
	"github.com/olivere/elastic"
)

type Service struct {
	elasticClient *elastic.Client
}

func NewService() *Service {
	return new(Service)
}

func (s *Service) RegisterElasticClient(c *elastic.Client) {
	s.elasticClient = c
}
