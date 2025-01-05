package main

import (
  "time"
  "fmt"
  "errors"
)

type todo struct {
  name        string
  completed   bool
  createdAt   time.Time
  completedAt *time.Time
}

type Todos []todo


func (t Todos) formattedPrint() {
  fmt.Println("------")
  for i := 0; i < len(t); i++ {
fmt.Printf(`Name: %s 
Completed: %t
Index: %d
Created at: %s
Completed at: %s


`, t[i].name, t[i].completed, i)
  }
}

func (t *Todos) add(name string) {
  newTodo := todo{
    name: name,
    completed: false,
    createdAt: time.Now(),
  }

    *t = append(*t, newTodo)
}

func (t Todos) validateIndex(index int) error {
  if index < 0 || index >= len(t) {
    err := errors.New("Index out of range")
    fmt.Printf("%s: %d\n", err, index)
    return err
  } 
  return nil
}

func (t *Todos) delete(index int) error {
  if err :=  t.validateIndex(index); err != nil {
    return err
  } else {
    tempList := *t
    *t = append(tempList[:index], tempList[index + 1:]...)
    return nil
  }
}

func (t *Todos) toggle(index int) error {
  if err :=  t.validateIndex(index); err != nil {
    return err
  } else {
    (*t)[index].completed = !(*t)[index].completed

    if (*t)[index].completed == true {
      completionTime := time.Now()
      (*t)[index].completedAt = &completionTime
    } else {
      (*t)[index].completedAt = nil
    }

    completionTime := time.Now()


    (*t)[index].completedAt = &completionTime

    return nil
  }

}

