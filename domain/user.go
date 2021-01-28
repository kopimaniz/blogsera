package domain

import (
	"time"
)

type UserStatus int

const(
  UserStatusActive UserStatus = 1
  UserStatusDeleted UserStatus = 0
)

type User struct {
	UserID    int        `json:"user_id,omitempty"`
	Username  string     `json:"username,omitempty"`
	Password  string     `json:"password,omitempty"`
	Email     string     `json:"email,omitempty"`
	FirstName string     `json:"first_name,omitempty"`
	LastName  string     `json:"last_name,omitempty"`
	Status    UserStatus `json:"status,omitempty"`
	TCreated  *time.Time  `json:"t_created,omitempty"`
	TUpdated  *time.Time  `json:"t_updated,omitempty"`
}

type UserRepository interface{
  Get(ID int)(*User, error)
  GetByUsername(username string)(*User, error)
  GetAll(activeOnly bool)([]*User, error)
  Save(u *User) (*User, error)
  Update(ID int, u *User) (*User, error)
  Delete(ID int) error
}

type UserService interface{
  Save(u *User) (*User, error)
  Get(ID int)(*User, error)
  GetAll(activeOnly bool)([]*User, error)
  Update(ID int, u *User) (*User, error)
}
