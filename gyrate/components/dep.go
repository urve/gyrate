package components

import (
  "os"
  "os/exec"
  "fmt"
)

func Dependency() {
  // ./gyrate --dep ../../source/txt ../../db.db full [TERM OR DIRECTORY]
  if len(os.Args) < 6 {
    fmt.Println("Not enough arguments.")
    os.Exit(1)
  }
  script := fmt.Sprintf(`python3 ./components/dep/dep.py %s %s %s %s`, os.Args[2], os.Args[3], os.Args[4], os.Args[5])
  cmd := exec.Command("bash", "-c", script)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Run()
}
