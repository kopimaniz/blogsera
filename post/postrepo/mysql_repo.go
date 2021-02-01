package postrepo

import (
	"blogsera/common/cerror"
	"blogsera/common/cmysql"
	"blogsera/domain"
	"database/sql"
)

type mysqlRepo struct{
  db *sql.DB
}

func NewMysql(db *sql.DB) domain.PostRepository{
  return &mysqlRepo{db}
}


func (r *mysqlRepo) Get(ID int) (*domain.Post, error) {
  var post domain.Post
  var nullTm cmysql.MysqlNullTIme

  row := r.db.QueryRow("select post_id, user_id, title, content, status, t_created, t_update from post where post_id=?", ID)

  err := row.Scan(&post.PostID, &post.UserID, &post.Title, &post.Content, &post.Status, &post.TCreated, &nullTm)
  if err != nil {
    if err == sql.ErrNoRows{
      return nil, cerror.ErrPostNotFound
    }
    return nil, err
  }

  // handle null time
  if nullTm.Valid{
    post.TUpdated = &nullTm.Time
  }
}

func (r *mysqlRepo) GetAll() ([]*domain.Post, error) {
  var posts []*domain.Post


  rows, err := r.db.Query("select post_id, user_id, title, content, status, t_created, t_update from post")
  if err!= nil{
    return nil, err
  }
  defer rows.Close()

  for rows.Next(){
    var post domain.Post
    var nullTm cmysql.MysqlNullTIme

    err := rows.Scan(&post.PostID, &post.UserID, &post.Title, &post.Content, &post.TCreated, nullTm)

    if err!= nil {
      return nil, err
    }

    if nullTm.Valid {
      post.TUpdated = &nullTm.Time
    }

    posts = append(posts, &post)
  }

  if rows.Err() != nil {
    return nil, rows.Err()
  }

  return posts, nil
}

func (r *mysqlRepo) GetByUser(UserID int) ([]*domain.Post, error) {
	var posts []*domain.Post


  rows, err := r.db.Query("select post_id, user_id, title, content, status, t_created, t_update from post where user_id=?", UserID)
  if err!= nil{
    return nil, err
  }
  defer rows.Close()

  for rows.Next(){
    var post domain.Post
    var nullTm cmysql.MysqlNullTIme

    err := rows.Scan(&post.PostID, &post.UserID, &post.Title, &post.Content, &post.TCreated, nullTm)

    if err!= nil {
      return nil, err
    }

    if nullTm.Valid {
      post.TUpdated = &nullTm.Time
    }

    posts = append(posts, &post)
  }

  if rows.Err() != nil {
    return nil, rows.Err()
  }

  return posts, nil
}


func(r *mysqlRepo) Save(post *domain.Post)(*domain.Post, error){
  result, err := r.db.Exec("insert into post(user_id, title, content) values(?,?,?)", post.UserID, post.Title, post.Content)

  if err!= nil{
    return nil, err
  }

  _, err = result.LastInsertId()
  if err!= nil{
    return nil, err
  }

  return nil, nil
}

