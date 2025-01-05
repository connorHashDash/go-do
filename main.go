package main

import (
//  "fmt"
)

func main() {
  exampleTodos := Todos{}
  exampleTodos.add("Doodle")
  exampleTodos.add("Invent")
  exampleTodos.add("Golf")
  exampleTodos.delete(0)
  exampleTodos.toggle(1)
//  exampleTodos.toggle(1)
  exampleTodos.edit(1, "Paint") 
  exampleTodos.tableView()
}
