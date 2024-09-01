package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"simple-serv/internal/aws"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var cred = "turso"

type conf struct {
  a  string
  b  string
}

func openMain() (*sql.DB, error) {
  db := &sql.DB{}
  sm, err := aws.GetSMClient("us-east-1")
  if err != nil { return db, err }
  s, err := sm.GetSecret(&cred)
  
  conf := &conf{}
  err = json.Unmarshal([]byte(*s), conf)
  if err != nil { return db, err }

  url := fmt.Sprintf("libsql://%s.turso.io?authToken=%s", conf.a, conf.b)
  db, err = sql.Open("libsql", url)
  if err != nil { return db, err }
  return db, err
}

func OpenUser(id *string) (*sql.DB, error) {
  return &sql.DB{}, nil
}
