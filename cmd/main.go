package cmd

import (
	"blogsera/config"
	"blogsera/user/userhandler"
	"blogsera/user/userrepo"
	"blogsera/user/userservice"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT = 8080

func Start(){
  conf := config.LoadMysqlConfig("config/local_mysql.json")
  db := config.NewMysqlDB(conf)
  defer db.Close()

  h := mux.NewRouter()

  userRoute(db,h)

  server := http.Server{
    Addr: fmt.Sprintf(":%v",PORT),
    Handler: h,
  }

  log.Printf("run server on port %v", PORT)
  log.Fatal(server.ListenAndServe())
}

func userRoute(db *sql.DB, h *mux.Router){
  repo := userrepo.NewMysql(db)
  service := userservice.New(repo)
  handler := userhandler.NewHTTP(service)

  h.HandleFunc("/users", handler.GetAll).Methods("GET")
  h.HandleFunc("/users", handler.Save).Methods("POST")
  h.HandleFunc("/users/{id:[0-9]+}", handler.Get).Methods("GET")
  h.HandleFunc("/users/{id:[0-9]+}", handler.Update).Methods("PUT")
  h.HandleFunc("/users/{id:[0-9]+}", handler.Delete).Methods("DELETE")
}
