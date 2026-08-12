package components

import (
  "github.com/fatih/color"
)

func Command() func(a ...interface{}) string {
  // The return type is suggested by Claude
  return color.New(color.Bold, color.FgGreen).SprintFunc()
}

func Option() func(a ...interface{}) string {
  return color.New(color.Bold, color.FgBlue).SprintFunc()
}

func Note() func(a ...interface{}) string {
  return color.New(color.Bold).SprintFunc()
}

func Header() func(a ...interface{}) string {
  return color.New(color.Bold, color.FgYellow).SprintFunc()
}
