package main

import (
  "fmt"
  "os"
  "bufio"
)

func ReadFile(path string) []string {
  file, err := os.Open(path)
  if err != nil {
    fmt.Println("Error opening file", err)
    os.Exit(1)
  }
  defer file.Close()

  var lines []string

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  return lines
}

func OpenToAppend(path string, bookDone string) (*os.File, error) {
  f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  dash := "───────────────────"
  fmt.Fprintln(f, dash, bookDone, dash)

  return f, err
}

func SendToFile(sentence string,
                sentenceOption string,
                fileYes *os.File,
                fileNo *os.File) {
  if sentenceOption == "true" {
    fmt.Fprintln(fileYes, sentence)
  } else if sentenceOption == "false" {
    fmt.Fprintln(fileNo, sentence)
  }
}
