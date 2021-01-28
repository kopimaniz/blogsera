package config

import (
	"database/sql"
	"fmt"
	"os"

	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
}

func LoadMysqlConfig(path string) *MysqlConfig{
  file, err := os.Open(path)
  defer file.Close()
  if err!= nil {
    panic(err)
  }

  var conf MysqlConfig
  err = json.NewDecoder(file).Decode(&conf)
  if err!= nil{
    panic(err)
  }
  return &conf
}

func NewMysqlDB(conf *MysqlConfig) *sql.DB{
  sqlStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",conf.Username,conf.Password,conf.Host, conf.Port,conf.Name)
  db, err := sql.Open("mysql", sqlStr)
  if err!= nil{
    panic(err)
  }

  return db
}
