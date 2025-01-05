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
    completedAtStr := "nil"
    if t[i].completedAt != nil {
      completedAtStr = t[i].completedAt.Format("2006-01-02 15:04:05")
    }
fmt.Printf(`Name: %s 
Completed: %t
Created at: %s
Completed at: %s
Index: %d


`, t[i].name, t[i].completed,  t[i].createdAt.Format("2006-01-02 15:04:05"), completedAtStr, i)
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
    return nil
  }
}

func (t *Todos) edit(index int, name string) error {
  if err :=  t.validateIndex(index); err != nil {
    return err
  } else {
    (*t)[index].name = name

    return nil
  }
}

