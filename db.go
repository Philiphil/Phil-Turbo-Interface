package main

import (
    "github.com/tidwall/buntdb"
)

var db *buntdb.DB



func init() {
  e, _ := buntdb.Open("PTI.db")
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

type Jrep struct{
  arg string `json:"arg"`
  data []string `json:"data"`
}

