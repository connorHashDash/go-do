package main

import (
  "flag"
)

func (t *Todos) commandParse() {
  var view = flag.Bool("view", false, "View the Todos in a table")
  flag.Parse()
  if *view {
    t.tableView()
  }
}

