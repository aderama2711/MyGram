package service

import "MyGram/repository"

type Service struct {
	repo repository.RepoInterface
}

type ServiceInterface interface {
	MyGramService
}

func NewService(repo repository.RepoInterface) ServiceInterface {
	return &Service{repo: repo}
}
