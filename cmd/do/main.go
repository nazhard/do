package main

import (
  "fmt"

  "github.com/nazhard/do"
)

func main() {
  out, err := do.Stuff()
  if err != nil {
    fmt.Println(err)
  }

  fmt.Print(out)
}
