package main

import (
//  "fmt"
)

func main() {
  exampleTodos := Todos{}
  exampleTodos.add("poo")
  exampleTodos.add("wee")
  exampleTodos.add("shart")
  exampleTodos.delete(0)
  exampleTodos.toggle(1)
  exampleTodos.formattedPrint()
}
