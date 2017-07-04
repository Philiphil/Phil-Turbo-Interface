package main

import (
    "github.com/tidwall/buntdb"
)

var db *buntdb.DB



func init() {
  e, _ := buntdb.Open("PTS.db")
  db = e
}

func db_get(key string) string{
  var re = ""
  db.View(func(tx *buntdb.Tx) error {
    re, _ = tx.Get(key) 
    return nil
})
  return re
}

func db_set(key string, data string){
  db.Update(func(tx *buntdb.Tx) error {
    tx.Set(key, data, nil)
    return nil
})
}

type Jrep struct{
  Arg string `json:"arg"`
  Data []string `json:"data"`
}