package userhandler

import (
	"blogsera/common/cerror"
	"blogsera/common/chttp"
	"blogsera/domain"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type httpHandler struct{
  s domain.UserService
}

func NewHTTP(s domain.UserService) *httpHandler{
  return &httpHandler{s}
}

func(h *httpHandler) Get(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  ID, err := strconv.Atoi(vars["id"])
  if err!= nil {
    chttp.ResponseFail(http.StatusText(http.StatusNotFound), nil).AsJson(w)
    return
  }

  user, err := h.s.Get(ID)
  if err!= nil{
    chttp.ResponseFail(http.StatusText(http.StatusNotFound), nil).AsJson(w)
    return
  }
  chttp.ResponseSuccess(user).AsJson(w)
}

func(h *httpHandler) Update(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  ID, err := strconv.Atoi(vars["id"])
  if err!= nil {
    log.Println("handler : "+err.Error())
    chttp.ResponseFail(http.StatusText(http.StatusNotFound), nil).AsJson(w)
    return
  }

  var userData domain.User

  err = json.NewDecoder(r.Body).Decode(&userData)
  defer r.Body.Close()
  if err!= nil {
    log.Println("handler : "+err.Error())
    chttp.ResponseError(http.StatusText(http.StatusInternalServerError)).AsJson(w)
    return
  }

  user, err := h.s.Update(ID, &userData)
  if err!= nil{
    log.Println("handler : "+err.Error())
    if err == cerror.ErrUserNotFound {
      chttp.ResponseFail(http.StatusText(http.StatusNotFound), nil).AsJson(w)
      return
    }
    chttp.ResponseError(http.StatusText(http.StatusInternalServerError)).AsJson(w)
    return
  }

  chttp.ResponseSuccess(user).AsJson(w)
}

func(h *httpHandler) Delete(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  ID, err := strconv.Atoi(vars["id"])
  if err!= nil {
    log.Println("handler : "+err.Error())
    chttp.ResponseFail(http.StatusText(http.StatusNotFound), nil).AsJson(w)
    return
  }

  userData := domain.User{UserID: ID, Status: 0}
  _, err = h.s.Update(ID, &userData)
  if err!= nil{
    log.Println("handler : "+err.Error())
    if err == cerror.ErrUserNotFound {
      chttp.ResponseFail(http.StatusText(http.StatusNotFound), nil).AsJson(w)
      return
    }
    chttp.ResponseError(http.StatusText(http.StatusInternalServerError)).AsJson(w)
    return
  }

  chttp.ResponseSuccess(nil).AsJson(w)
}

func(h *httpHandler) GetAll(w http.ResponseWriter, r *http.Request){
  users, err := h.s.GetAll(true)
  if err!= nil{
    log.Println("handler "+err.Error())
    chttp.ResponseError(http.StatusText(http.StatusInternalServerError)).AsJson(w)
    return
  }

  chttp.ResponseSuccess(users).AsJson(w)
}

func(h *httpHandler) Save(w http.ResponseWriter, r *http.Request){
  if r.Method != http.MethodPost {
    chttp.ResponseFail(http.StatusText(http.StatusMethodNotAllowed), nil).AsJson(w)
    return
  }

  var u domain.User
  err := json.NewDecoder(r.Body).Decode(&u)
  if err!= nil{
    chttp.ResponseFail(http.StatusText(http.StatusBadRequest), nil).AsJson(w)
    return
  }

  user, err := h.s.Save(&u)
  if err!= nil{
    if err == cerror.ErrUserExist {
      chttp.ResponseFail("User telah terdaftar", nil).AsJson(w)
      return
    }

    chttp.ResponseFail(http.StatusText(http.StatusBadRequest), nil).AsJson(w)
    return
  }

  chttp.ResponseSuccess(user).AsJson(w)
}
