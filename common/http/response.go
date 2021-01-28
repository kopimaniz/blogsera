package http

import (
	"encoding/json"
	"net/http"
)

type Status string

const(
  Success Status = "success"
  Fail Status = "fail"
  Error Status = "error"
)

type Response struct{
  Status Status         `json:"status,omitempty"`
  Message string        `json:"message,omitempty"`
  Data interface{}      `json:"data,omitempty"`
}

func(r *Response) Json(w http.ResponseWriter){
  w.Header().Set("Content-Type","application/json")
  json.NewEncoder(w).Encode(r)
}

func ResponseSuccess(data interface{}) *Response{
  return &Response{
    Status: Success,
    Data: data,
  }
}

func ResponseFail(message string, data interface{}) *Response{
  return &Response{
    Status: Fail,
    Message: message,
    Data: data,
  }
}

func ResponseError(message string) *Response{
  return &Response{
    Status: Error,
    Message: message,
  }
}
