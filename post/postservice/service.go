package service

import "blogsera/domain"

type service struct{
  r domain.PostRepository
}

func New(repo domain.PostRepository) domain.PostService{
  return &service{r: repo}
}

func (s *service) Save(post *domain.Post) (*domain.Post, error) {
	panic("not implemented") // TODO: Implement
}

