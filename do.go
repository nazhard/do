package do

import (
  "errors"
  "bufio"
  "fmt"
  "os"
  "os/exec"
  "strings"
)

var (
  script string
  args string
  name string
)

func Stuff() (string, error) {
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
    }

    meh, err := check(args)
    if err != nil {
      fmt.Println(err)
    }

    if strings.Contains(meh, "bash") {
      name = "bash"
    } else if strings.Contains(meh, "fish") {
      name = "fish"
    }
  }

  cmd := exec.Command(name, args)
  // fmt.Println(name, args)

  out, _ := cmd.Output()

  return string(out), nil
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

func checkFile(file string) error {
  _, err := os.Stat(file)
  if os.IsNotExist(err) {
    return errors.New(file+" doesn't exist")
  }
  
  return nil
}
