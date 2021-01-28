package error

import "errors"

var ErrUserNotFound = errors.New("user tidak ditemukan")
var ErrUserExist  = errors.New("user telah terdaftar")
