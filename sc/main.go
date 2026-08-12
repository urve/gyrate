package main

import (
  "os"
  "strings"
)

func main () {
  if len(os.Args) <= 3 || (os.Args[1] == "-h") {
    Usage()
  } else {
    fileToClassify := os.Args[1]
    yesPath := os.Args[2]
    noPath  := os.Args[3]
    yesFile, _ := OpenToAppend(yesPath, fileToClassify)
    noFile, _  := OpenToAppend(noPath,  fileToClassify)
    defer yesFile.Close()
    defer noFile.Close()

    lines := ReadFile(fileToClassify)

    sentenceArray := SplitOnRunes(lines)
    LineInstructions(sentenceArray)
    line := 0
    for _, sentence := range sentenceArray {
      line += 1
      sentence = strings.TrimSpace(sentence)
      sentence = strings.ToLower(sentence)
      sentenceOption := GetOption(line, sentence)
      SendToFile(sentence, sentenceOption, yesFile, noFile)
   }
  }
}
