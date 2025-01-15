package main

import (
//  "fmt"
)

func main() {
  exampleTodos := Todos{}
  exampleTodos.load()
  exampleTodos.tableView()
}
