package main

import (
  "os"
  "gyrate/components"
)

func main() {
  if len(os.Args) == 1 {
    components.Usage()
  } else if os.Args[1] == "--bayes" {
    components.Bayes()
  } else if os.Args[1] == "--bayes-all" {
    components.BayesAll()
  } else if os.Args[1] == "--bayes-sentence" {
    components.BayesSentence()
  } else if os.Args[1] == "--logistic" {
    components.Logistic()
  } else if os.Args[1] == "--logistic-all" {
    components.LogisticAll()
  } else if os.Args[1] == "--logistic-rebuild" {
    components.LogisticRebuild()
  } else if os.Args[1] == "--year-bridge" {
    components.YearBridge()
  } else if os.Args[1] == "--examples" {
    components.Examples()
  } else if os.Args[1] == "--duplicates" {
    components.Duplicates()
  } else if os.Args[1] == "--embed" {
    components.Embed()
  } else if os.Args[1] == "--embed-rebuild" {
    components.EmbedRebuild()
  } else if os.Args[1] == "--normalize" {
    components.Normalize()
  } else if os.Args[1] == "--gazetteer" {
    components.Gazetteer()
  } else if os.Args[1] == "--dependency" {
    components.Dependency()
  } else if os.Args[1] == "--license" {
    components.License()
  } else {
    components.Usage()
  }
}
