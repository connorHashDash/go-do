package main

import (
//  "fmt"
)

func main() {
  exampleTodos := Todos{}
  exampleTodos.add("Doodle")
  exampleTodos.add("Invent")
  exampleTodos.add("Golf")
  exampleTodos.formattedPrint()
  exampleTodos.delete(0)
  exampleTodos.formattedPrint()
  exampleTodos.toggle(1)
  exampleTodos.formattedPrint()
  exampleTodos.toggle(1)
  exampleTodos.formattedPrint()
  exampleTodos.edit(1, "Paint")
  exampleTodos.formattedPrint()
}
