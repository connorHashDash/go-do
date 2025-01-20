package main

func main() {
  exampleTodos := Todos{}
  exampleTodos.load()
  exampleTodos.commandParse()
  exampleTodos.save()
}
