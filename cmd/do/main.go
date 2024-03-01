package main

import (
  "fmt"

  "github.com/nazhard/do"
)

func main() {
  cmd, err := do.Stuff()
  if err != nil {
    fmt.Println(err)
  }

  _ = cmd.Start()

  defer cmd.Wait()
}
