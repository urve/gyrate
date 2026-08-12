package components

import (
  "os"
  "os/exec"
  "fmt"
)

func Embed() {
  if len(os.Args) < 4 {
    fmt.Println("Not enough arguments.")
    os.Exit(1)
  }
  script := fmt.Sprintf(`python3 -u ./components/embed/embed.py -l %s %s
`, os.Args[2], os.Args[3])
  cmd := exec.Command("bash", "-c", script)
  cmd.Stdout = os.Stdout
  cmd.Run()
}

func EmbedRebuild() {
  if len(os.Args) < 4 {
    fmt.Println("Not enough arguments.")
    os.Exit(1)
  }
  script := fmt.Sprintf(`python3 -u ./components/embed/embed.py -r %s %s
`, os.Args[2], os.Args[3])
  // Claude suggested the `-u` option to allow the script to be redirected
  // into a file, or to allow it to be piped into another process
  // (I wanted to be able to use `tee` to both direct it into a file and
  //  to watch the progress of the script.)
  cmd := exec.Command("bash", "-c", script)
  cmd.Stdout = os.Stdout
  cmd.Run()
}
