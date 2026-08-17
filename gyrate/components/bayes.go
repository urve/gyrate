package components

import (
  "fmt"
  "os"
  "os/exec"
  "strings"
  "github.com/jbrukh/bayesian"
)

func BayesBase() (classifier *bayesian.Classifier,
                  toBeClassified []string,
                  toBeClassifiedName string,
                  yesName string,
                  noName string) {
  if len(os.Args) < 5 {
    // gyrate --bayes FILE YES NO
    fmt.Println("Not enough arguments.")
    os.Exit(1)
  }
  const (
    Yes  bayesian.Class = "Yes"
    No   bayesian.Class = "No"
  )
  classifier = bayesian.NewClassifier(Yes, No)

  // Yes: First file
  // No:  Second file
  yes, _ := ReadFile(os.Args[3])
  yes = CutSliceIntoWords(yes)
  yesName = GetFileName(os.Args[3])
  no, _ := ReadFile(os.Args[4])
  no = CutSliceIntoWords(no)
  noName = GetFileName(os.Args[4])

  // Learn on Yes and No
  classifier.Learn(yes, Yes)
  classifier.Learn(no,  No)

  // Split into sentences
  toBeClassified, toBeClassifiedName = ReadFile(os.Args[2])
  toBeClassified = LinesToSentences(toBeClassified)

  return classifier, toBeClassified, toBeClassifiedName, yesName, noName
}

func Bayes() {
  classifier, toBeClassified, toBeClassifiedName, yesName, noName := BayesBase()
  classes := []string{yesName, noName}
  runningYes := 0
  runningNo  := 0
  for _, line := range toBeClassified {
    lineWords := Cut(line)
    _, likely, _ := classifier.LogScores(lineWords)
    // Running count of sentences classified Yes or No
    if classes[likely] == yesName {
      runningYes += 1
    } else if classes[likely] == noName {
      runningNo += 1
    }
  }

  yesFloat := float64(runningYes)
  noFloat := float64(runningNo)
  val := 100.0 * yesFloat / (yesFloat + noFloat)
  // id,value
  fmt.Printf("%s,%.2f\n", toBeClassifiedName, val)
}

func BayesSentence() {
  classifier, toBeClassified, _, _, _ := BayesBase()
  //classes := []string{yesName, noName}

  runningTotal := 0.0
  for index, line := range toBeClassified {
    lineWords := Cut(line)
    scores, _, _ := classifier.LogScores(lineWords)
    // Score of `yes` - `no`
    score := scores[0] - scores[1]
    runningTotal += score
    fmt.Printf("%d,%.8f,%.8f\n", index, runningTotal, score)
  }
}

func BayesAll() {
  if len(os.Args) < 5 {
    fmt.Println("Not enough arguments.")
    os.Exit(1)
  }
  directory := os.Args[2]
  if strings.HasSuffix(directory, "/") {
    directory = directory[:len(directory)-1]
  }
  script := fmt.Sprintf(`
shopt -s nullglob
parallel './gyrate --bayes {} %s %s' ::: %s/*.txt
`, os.Args[3], os.Args[4], directory)
  cmd := exec.Command("bash", "-c", script)
  cmd.Stdout = os.Stdout
  cmd.Run()
}
