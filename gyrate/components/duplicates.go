package components

import (
  "os"
  "os/exec"
  "fmt"
)

func Duplicates() {
  if len(os.Args) < 3 {
    fmt.Println("Not enough arguments.")
    os.Exit(1)
  }
  script := fmt.Sprintf(`./components/dup/dup.py %s
`, os.Args[2])
  cmd := exec.Command("bash", "-c", script)
  cmd.Stdout = os.Stdout
  cmd.Run()
}
