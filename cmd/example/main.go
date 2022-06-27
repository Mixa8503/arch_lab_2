package main

import (
  "flag"
  "fmt"
  "io"
  "log"
  "os"
  "strings"

  lab2 "github.com/Mixa8503/arch_lab_2"
)

var (
  inputExpression = flag.String("e", "", "Expression to compute")
  Input_File = flag.String("f", "", "File with expression to compute")
  Output_File = flag.String("o", "", "File to store computed expression")
)

func main() {
  flag.Parse()
  var err error

  if *inputExpression == "" && *Input_File == "" {
    log.Fatal("no expression provided. use -e \"{expression}\" or -f {file with expression}")
  }
  if *inputExpression != "" && *Input_File != "" {
    log.Fatal("flags -f and -e can't be used at the same time")
  }

  var reader io.Reader
  if *inputExpression != "" {
    reader = strings.NewReader(*inputExpression)
  } else {
    if file, err := os.Open(*Input_File); err == nil {
      defer file.Close()
      reader = file
    } else {
      log.Fatal("There is no such file")
    }
  }

  var writer io.Writer
  if *Output_File != "" {
    if file, err := os.Create(*Output_File); err == nil {
      defer file.Close()
      writer = file
    } else {
      log.Fatal("Something is wrong, check your file and try again")
    }
  } else {
    writer = &Writer{}
  }

  handler := &lab2.ComputeHandler{
    Input:  reader,
    Output: writer,
  }
  err = handler.Compute()
  if err != nil {
    panic(err)
  }
}

type Writer struct{}

func (w *Writer) Write(data []byte) (n int, err error) {
  fmt.Println(string(data))
  return len(data), nil
}
