package main

import (
  "fmt"
  "strings"
  "github.com/mattn/go-tty/v2"
)

func GetOption(line int, sentence string) string {
  redo := true
  var returnOption string
  //reader := bufio.NewReader(os.Stdin)

  runeReader, _ := tty.Open()
  defer runeReader.Close()

  for redo {

    if sentence == "" {
      continue
    }
    fmt.Printf("[%d] %s\n", line, sentence)
    //option, _ := reader.ReadString('\n')
    //option = strings.TrimSpace(option)
    //option = strings.ToLower(option)

    option, _, _ := runeReader.ReadRune()

    if option == 'y' || option == 'Y' {
      redo = false
      returnOption = "true"
    } else if option == 'n' || option == 'N' {
      redo = false
      returnOption = "false"
    } else if option == 's' || option == 'S' {
      redo = false
      returnOption = "skip"
    } else {
      fmt.Println("Not recognized. Try again.")
      redo = true
    }
  }
  return returnOption
}

func SplitOnRunes(lines []string) []string {
  sentenceArray := []string{}
  for _, line := range lines {
    splitOnRunes := func(r rune) bool {
      if (r == '.' || r == '!' || r == '?') {
        return true
      } else {return false}
    }
    sentences := strings.FieldsFunc(line, splitOnRunes)
    for _, sentence := range sentences {
      sentenceArray = append(sentenceArray, sentence)
    }
  }
  return sentenceArray
}
