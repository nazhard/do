package do

import (
  // "errors"
  "fmt"
  "os"
  "os/exec"
)

var script string

func Stuff() (string, error) {
  numArgs := len(os.Args)

  if numArgs == 2 {
    script = os.Args[1]
    args := "./scripts/"+script+".sh"

    cmd := exec.Command("bash", args)

    out, err := cmd.Output()
    if err != nil {
      return "", err
    }
    
    return string(out), nil
  }

  if numArgs == 1 {
    fmt.Println("do v0.0.1 \n")
    fmt.Println("automaticly run your .sh file on ./scripts/ directory.")
  }

  return "", nil
}
