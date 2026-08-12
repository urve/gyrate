package components

import (
  "strings"
)

func Cut(text string) []string {
  slice := strings.Split(text, " ")
  return slice
}

func CutSliceIntoWords(line []string) []string {
  joinedLine := strings.Join(line, " ")
  words := strings.Split(joinedLine, " ")
  return words
}

func LinesToSentences(lines []string) []string {
  sentenceArray := []string{}
  allText := strings.Join(lines, " ")
  splitOnPunctuation := func(r rune) bool {
    if (r == '.' || r == '!' || r == '?') {
      return true
    } else {return false}
  }
  sentences := strings.FieldsFunc(allText, splitOnPunctuation)
  for _, sentence := range sentences {
    trimmedSentence := strings.TrimSpace(sentence)
    if strings.HasPrefix(trimmedSentence, "─") {
      continue
    } else {
      sentenceArray = append(sentenceArray, trimmedSentence)
    }
  }
  return sentenceArray
}
