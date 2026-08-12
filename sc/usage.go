package main

import (
  "fmt"
)

func Usage() {
  fmt.Println("sc (split-and-classify) takes a file and, sentence by sentence,")
  fmt.Println("allows the user to classify them into one of two categories.")
  fmt.Println()
  fmt.Println("Usage:")
  fmt.Println("  sc CANNIBALIZED_FILE FIRST_CLASS SECOND_CLASS")
  fmt.Println()
  fmt.Println("Parameters:")
  fmt.Println("  CANNIBALIZED_FILE: The file that will be torn apart.")
  fmt.Println("  FIRST_CLASS      : File path to first class to be made.")
  fmt.Println("  SECOND_CLASS     : File path to the second class to be made.")
}

func LineInstructions(sentenceArray []string) {
  dash := "──────────────────────────"
  fmt.Printf("%s\nPlease enter 'y' or 'n' for whether to classify each line\nas either YES_FILE or NO_FILE. Quit and use option `-h` to read more.\nTotal lines: %d.\n%s\n", dash, len(sentenceArray), dash)
}
