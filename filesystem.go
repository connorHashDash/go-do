package main

import (
  "encoding/json"
  "fmt"
  "os"
)

func (t Todos) toJson() []byte {
  fmt.Println(t)
  var todoJson, err = json.MarshalIndent(t, "", "  ")
  if err != nil {
    fmt.Println(err)
    return nil
  }

  fmt.Println(string(todoJson))
  return todoJson
}

func (t Todos) save() {
  json := t.toJson()

  os.WriteFile("./todo-list.json", json, 0644)
}
