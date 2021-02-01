package domain

import "time"

type Post struct {
	PostID   int        `json:"post_id,omitempty"`
	UserID   int        `json:"user_id,omitempty"`
	Title    string     `json:"title,omitempty"`
	Content  string     `json:"content,omitempty"`
	Status   bool       `json:"status,omitempty"`
	TCreated  *time.Time `json:"t_creade,omitempty"`
	TUpdated *time.Time `json:"t_updated,omitempty"`
}

type PostRepository interface{
  Save(post *Post)(*Post, error)
  Get(ID int)(*Post, error)
  GetAll()([]*Post, error)
  GetByUser(UserID int)([]*Post, error)
}

type PostService interface{
  Save(post *Post)(*Post, error)
}
