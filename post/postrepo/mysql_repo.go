package userRepo

import (
	"blogsera/domain"
	"database/sql"
)

type mysqlRepo struct{
  db *sql.DB
}

func NewMysql(db *sql.DB) domain.PostRepository{
  return &mysqlRepo{db}
}

func(r *mysqlRepo) Save(post *domain.Post)(*domain.Post, error){
  result, err := r.db.Exec("insert into post(user_id, title, content) values(?,?,?)", post.UserID, post.Title, post.Content)

  if err!= nil{
    return nil, err
  }

  _, err := result.LastInsertId()
  if err!= nil{
    return nil, err
  }

  return nil, nil
}
