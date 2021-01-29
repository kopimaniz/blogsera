package postrepo

import (
	"reflect"

	"github.com/go-sql-driver/mysql"
)

// handle null time
type MysqlNullTIme mysql.NullTime

func(nt *MysqlNullTIme) Scan(value interface{}) error{
  var t mysql.NullTime

  if err := t.Scan(value); err != nil{
    return err
  }

  if reflect.TypeOf(value) == nil {
    *nt = MysqlNullTIme{t.Time, false}
  } else {
    *nt = MysqlNullTIme{t.Time, true}
  }

  return nil
}
