package components

import(
  "fmt"
  "os"
  "bufio"
  "path/filepath"
  "strings"
)

func ReadFile(path string) ([]string, string) {
  file, err := os.Open(path)
  if err != nil {
    fmt.Println("Error opening file", err)
    os.Exit(1)
  }
  defer file.Close()

  var lines []string

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    if !strings.HasPrefix(line, "─") {
      lines = append(lines, scanner.Text())
    }
  }

  return lines, GetFileName(path)
}

func GetFileName(path string) string {
  filename := filepath.Base(path)
  filename = strings.TrimSuffix(filename, filepath.Ext(filename))
  return filename
}

