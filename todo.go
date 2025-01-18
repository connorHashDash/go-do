package main

import (
  "time"
  "fmt"
  "errors"
  "os"
  "github.com/aquasecurity/table"
  "strconv"
)

type todo struct {
  Name        string 
  Completed   bool 
  CreatedAt   time.Time 
  CompletedAt *time.Time 
}

type Todos []todo

func (t *Todos) add(name string) {
  newTodo := todo{
    Name: name,
    Completed: false,
    CreatedAt: time.Now(),
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
    (*t)[index].Completed = !(*t)[index].Completed
    if (*t)[index].Completed == true {
      completionTime := time.Now()
      (*t)[index].CompletedAt = &completionTime
    } else {
      (*t)[index].CompletedAt = nil
    }
    return nil
  }
}

func (t *Todos) edit(index int, name string) error {
  if err :=  t.validateIndex(index); err != nil {
    return err
  } else {
    (*t)[index].Name = name

    return nil
  }
}

func (t Todos) tableView() {
  printedTable := table.New(os.Stdout)

  printedTable.SetAlignment(table.AlignRight, table.AlignRight, table.AlignCenter, table.AlignRight)
  printedTable.SetHeaders("#", "Name", "Completed", "Created-At", "Completed-At")
  printedTable.SetDividers(table.UnicodeRoundedDividers)

  for i := 0; i < len(t); i++ {
    completedAtStr := "nil"
    if t[i].CompletedAt != nil {
      completedAtStr = t[i].CompletedAt.Format("2006-01-02 15:04:05")
    }
    
    completedIco := "❌"
    if t[i].Completed != false {
      completedIco =  "✅"
    }


    printedTable.AddRow(strconv.Itoa(i), t[i].Name, completedIco, t[i].CreatedAt.Format("2006-01-02 15:04:05"), completedAtStr)
  }

  printedTable.Render()
}
