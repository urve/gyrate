package components

import (
  "os"
  "os/exec"
  "fmt"
  "strings"
)

func Logistic() {
  if len(os.Args) < 3 {
    fmt.Println("Not enough arguments.")
    os.Exit(1)
  }
  script := fmt.Sprintf(`./components/lr/lr.py --file %s
`, os.Args[2])
  cmd := exec.Command("bash", "-c", script)
  cmd.Stdout = os.Stdout
  cmd.Run()
}

func LogisticAll() {
  if len(os.Args) < 3 {
    fmt.Println("Not enough arguments.")
    os.Exit(1)
  }
  directory := os.Args[2]
  if strings.HasSuffix(directory, "/") {
    directory = directory[:len(directory)-1]
  }
  script := fmt.Sprintf(`
shopt -s nullglob
parallel ./components/lr/lr.py --file {} ::: %s/*.txt
`, directory)
  cmd := exec.Command("bash", "-c", script)
  cmd.Stdout = os.Stdout
  cmd.Run()
}

func LogisticRebuild() {
  if len(os.Args) < 4 {
    fmt.Println("Not enough arguments.")
    os.Exit(1)
  }
  script := fmt.Sprintf(`./components/lr/lr.py --rebuild %s %s`, os.Args[2], os.Args[3])
  cmd := exec.Command("bash", "-c", script)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Run()
}
