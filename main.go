package main

import (
//  "fmt"
)

func main() {
  exampleTodos := Todos{}
  exampleTodos.add("poo")
  exampleTodos.add("wee")
  exampleTodos.add("shart")
  exampleTodos.formattedPrint()
  exampleTodos.delete(0)
  exampleTodos.formattedPrint()
  exampleTodos.toggle(1)
  exampleTodos.formattedPrint()
  exampleTodos.toggle(1)
  exampleTodos.formattedPrint()
}
