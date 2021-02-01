package userservice

import (
	"blogsera/common/cerror"
	"blogsera/domain"
)

type service struct{
  r domain.UserRepository
}

func New(r domain.UserRepository) domain.UserService{
  return &service{
    r: r,
  }
}

func(s *service) Get(ID int)(*domain.User, error){
  user, err := s.r.Get(ID)
  if err!= nil{
    return nil, err
  }

  user.Password = ""
  return user, nil
}

func(s *service) Update(ID int, u *domain.User)(*domain.User, error){
  user, err := s.r.Get(ID)
  if err!= nil{
    return nil, err
  }

  if u.Username != "" {
    user.Username = u.Username
  }

  if u.Password != "" {
    user.Password = u.Password
  }

  if u.Email != "" {
    user.Email = u.Email
  }

  if u.FirstName != "" {
    user.FirstName = u.FirstName
  }

  if u.LastName != "" {
    user.LastName = u.LastName
  }

  return s.r.Update(ID, user)
}

func(s *service) GetAll(activeOnly bool)([]*domain.User, error){
  return s.r.GetAll(activeOnly)
}

func(s *service) Save(u *domain.User) (*domain.User, error){
  user, err := s.r.GetByUsername(u.Username)

  // error bukan karena not found
  if err != nil && err != cerror.ErrUserNotFound{
    return nil, err
  }

  if err == cerror.ErrUserNotFound{
    return s.r.Save(u)
  }

  // validate
  if user.Username == u.Username {
    return nil, cerror.ErrUserExist
  }

  return s.r.Save(u)
}
