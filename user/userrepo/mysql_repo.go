package userrepo

import (
	cerror "blogsera/common/error"
	"blogsera/domain"
	"database/sql"
)

type mySqlRepo struct{
  db *sql.DB
}

func NewMysql(db *sql.DB) domain.UserRepository{
  return &mySqlRepo{db}
}

// Get
func(r *mySqlRepo) Get(ID int)(*domain.User, error){
  var user domain.User
  var nullTm MysqlNullTIme

  row := r.db.QueryRow("select user_id, username, password, email, first_name, last_name, status, t_created, t_updated from user where user_id=?", ID)
  err := row.Scan(&user.UserID, &user.Username, &user.Password, &user.Email, &user.FirstName, &user.LastName, &user.Status, &user.TCreated, &nullTm)

  if err != nil {
    if err == sql.ErrNoRows {
      return nil, cerror.ErrUserNotFound
    }
    return nil, err
  }

  // handle null time
  if nullTm.Valid {
    user.TUpdated = &nullTm.Time
  }


  return &user, nil
}

func(r *mySqlRepo) GetByUsername(username string)(*domain.User, error){
  var user domain.User
  var nullTm MysqlNullTIme

  row := r.db.QueryRow("select user_id, username, email, first_name, last_name, status, t_created, t_updated from user where username=?", username)
  err := row.Scan(&user.UserID, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.Status, &user.TCreated, &nullTm)

  if err != nil {
    if err == sql.ErrNoRows {
      return nil, cerror.ErrUserNotFound
    }
    return nil, err
  }

  // handle null time
  if nullTm.Valid {
    user.TUpdated = &nullTm.Time
  }

  return &user, nil
}

func(r *mySqlRepo) GetAll(activeOnly bool)([]*domain.User, error){
  var users []*domain.User

  var rows *sql.Rows
  defer rows.Close()
  var err error
  if activeOnly {
    rows, err = r.db.Query("select user_id, username, email, first_name, last_name, status, t_created, t_updated from user where status=?", true)
  } else {
    rows, err = r.db.Query("select user_id, username, email, first_name, last_name, status, t_created, t_updated from user")
  }
  if err!= nil{
    return nil, err
  }

  for rows.Next(){
    var user domain.User
    var nullTm MysqlNullTIme

    err := rows.Scan(&user.UserID, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.Status,  &user.TCreated, &nullTm)
    if err!= nil {
      return nil, err
    }

    if nullTm.Valid {
      user.TUpdated = &nullTm.Time
    }

    users = append(users, &user)
  }

  if rows.Err() != nil {
    return nil, rows.Err()
  }

  return users, nil
}

// Save save user to table user
func(r *mySqlRepo) Save(u *domain.User) (*domain.User, error){
  result, err := r.db.Exec("insert into user(username, password, email, first_name, last_name) values(?, ?, ?, ?, ?)", u.Username, u.Password, u.Email, u.FirstName, u.LastName)
  if err != nil{
    return nil, err
  }

  lastID, err := result.LastInsertId()
  if err!= nil{
    return nil, err
  }
  return r.Get(int(lastID))
}

// Update update user
func(r *mySqlRepo) Update(ID int, u *domain.User) (*domain.User, error){
  _, err := r.db.Exec("update user set username=?, password=?, email=?, first_name=?, last_name=?, status=? where user_id=?", u.Username, u.Password, u.Email, u.FirstName, u.LastName, u.Status, ID)
  if err != nil{
    return nil, err
  }

  return r.Get(ID)
}

// Delete delete user
func(r *mySqlRepo) Delete(ID int) error{
  _, err := r.db.Exec("delete from user where user_id=?", ID)
  if err != nil{
    return err
  }
  return nil
}

