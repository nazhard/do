package do

import (
  "bufio"
  "errors"
  "fmt"
  "os"
  "os/exec"
  "strings"
)

var (
  script string
  args string
  name string
  n string
)

func Stuff() (*exec.Cmd, error) {
  numArgs := len(os.Args)

  if numArgs == 1 {
    fmt.Println("do v0.0.1 \n")
    fmt.Println("automaticly run your .sh file on ./scripts/ directory.")
  }
  
  if numArgs == 2 {
    script = os.Args[1]

    args = "./scripts/"+script
    err := checkFile(args)
    if err != nil {
      args = "./scripts/"+script+".sh"
      err = checkFile(args)
      if err != nil {
        args = "./scripts/"+script+".js"
        err = checkFile(args)
        if err != nil {
          args = "./scripts/"+script+".mjs"
        }
      }
    }

    meh, err := check(args)
    if err != nil {
      fmt.Println(err)
    }

    name = contains(meh)
  }

  if numArgs == 3 && os.Args[1] == "." {
    script = os.Args[2]

    args = "./"+script
    err := checkFile(args)
    if err != nil {
      args = "./"+script+".sh"
    } else if err != nil {
      args = "./"+script+".js"
    } else if err != nil {
      args = "./"+script+".mjs"
    }

    meh, err := check(args)
    if err != nil {
      fmt.Println(err)
    }

    name = contains(meh)
  }

  cmd := exec.Command(name, args)
  // fmt.Println(name, args)

  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  return cmd, nil
}

func check(path string) (string, error) {
  filePath := path

  file, err := os.Open(filePath)
  if err != nil {
    return "", errors.New(path+" doesn't exist")
  }
  
  defer file.Close()

  scanner := bufio.NewScanner(file)
  if scanner.Scan() {
    firstLine := scanner.Text()
    return firstLine, nil
  } else {
    return "", errors.New("error reading file")
  }
}

func contains(l string) string {
  if strings.Contains(l, "bash") {
    n = "bash"
  } else if strings.Contains(l, "fish") {
    n = "fish"
  } else if strings.Contains(l, "zx") {
    n = "zx"
  }

  return n
}

func checkFile(file string) error {
  _, err := os.Stat(file)
  if os.IsNotExist(err) {
    return errors.New(file+" doesn't exist")
  }
  
  return nil
}
