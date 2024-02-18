package main

import (
  "fmt"

  "github.com/nazhard/do"
)

func main() {
  _, err := do.Stuff()
  if err != nil {
    fmt.Println(err)
  }

  // fmt.Print(out)
}
