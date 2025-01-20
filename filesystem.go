package main

import (
  "encoding/json"
  "fmt"
  "os"
)

func (t Todos) toJson() []byte {
  var todoJson, err = json.MarshalIndent(t, "", "  ")
  if err != nil {
    fmt.Println(err)
    return nil
  }

  return todoJson
}

func (t Todos) save() {
  json := t.toJson()

  os.WriteFile("./todo-list.json", json, 0644)
}

func (t *Todos) load() error {
  //checks to see if file exists, ,creates blank file if not and returns
  if _, err := os.Stat("./todo-list.json"); err != nil {
    os.WriteFile("./todo-list.json", []byte(""), 0644)
    return nil
  }

  //loads file
  var readData, err = os.ReadFile("./todo-list.json")
  if err != nil {
    fmt.Println("error reading file")
    fmt.Println(err)
    return err
  }

  json.Unmarshal(readData, t)

  return nil
}
