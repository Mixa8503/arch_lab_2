package implement

import (
  "bytes"
  "fmt"
  "io"
  "strings"
)

type ComputeHandler struct {
  Input  io.Reader
  Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
  Buffer_Read := make([]byte, 128)
  _, err := ch.Input.Read(Buffer_Read)
  if err != nil {
    return err
  }
  Buffer_Read = bytes.Trim(Buffer_Read, "\x00")

  Trimmed_Expr := strings.Trim(string(Buffer_Read), " \n")
  res, err := Prefix_Eval(Trimmed_Expr)
  if err != nil {
    return err
  }

  Res_Byte := []byte(fmt.Sprint(res))
  _, err = ch.Output.Write(Res_Byte)
  if err != nil {
    return err
  }
  return nil
}
