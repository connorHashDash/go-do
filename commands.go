package main

import (
  "flag"
)

func (t *Todos) commandParse() {
  var view = flag.Bool("view", false, "View the Todos in a table")
  var add = flag.String("add", "", "Add a new Todo")

  flag.Parse()
  if *view {
    t.tableView()
  }

  if *add != "" {
    t.add(*add)
    t.tableView()
  }



}

