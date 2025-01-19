package main

import (
  "flag"
  "strings"
  "fmt"
  "strconv"
)

func (t *Todos) commandParse() {
  var view = flag.Bool("view", false, "View the Todos in a table")
  var add = flag.String("add", "", "Add a new todo")
  var remove = flag.Int("delete", -1, "Remove a todo")
  var toggle = flag.Int("toggle", -1, "Toggle a todo")
  var edit = flag.String("edit", "", "Edit the name of a todo")

  flag.Parse()

  if *view {
    t.tableView()
  }

  if *add != "" {
    t.add(*add)
    t.tableView()
  }

  if *remove != -1 {
    t.delete(*remove)
    t.tableView()
  }

  if *toggle != -1 {
    t.toggle(*toggle)
    t.tableView()
  }

  if *edit != "" {
    splitString := strings.Split(*edit, ":")
    if len(splitString) < 2 {
      fmt.Println("Correct formatting is 'index:name' \nExample: 1:NewName")
      return
    }

    if index, err := strconv.Atoi(splitString[0]); err == nil {
      t.edit(index, splitString[1])
      t.tableView()
    }
  }
}

