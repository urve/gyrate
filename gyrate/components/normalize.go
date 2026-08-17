package components

import (
  "fmt"
  "os"
  "encoding/csv"
  "strconv"
)

func Normalize() {
  if len(os.Args) < 2 {
    fmt.Println("Not enough arguments.")
    os.Exit(1)
  }

  // Default ranges to map scores to
  shapedMinimum := 0.0
  shapedMaximum := 100.0

  // everything after the option --normalize
  args := os.Args[2:]

  reader := os.Stdin

  // Either read a file path or STDIN

  if len(args) > 0 {
    _, parseErr := strconv.ParseFloat(args[0], 64)

    var isItNumerical bool
    if parseErr == nil {
      isItNumerical = true
    } else {
      isItNumerical = false
    }

    if !isItNumerical {
      // then args[0] must be a file path!
      csvFile, err := os.Open(args[0])
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      defer csvFile.Close()
      reader = csvFile
      args = args[1:]
    }
  }


  // Optional arguments to reshape curve values
  if len(args) >= 2 {
    parsedMinimum, err := strconv.ParseFloat(args[0], 64)
    if err != nil {
      fmt.Println("Invalid minimum: ", err)
      os.Exit(1)
    }

    parsedMaximum, err := strconv.ParseFloat(args[1], 64)
    if err != nil {
      fmt.Println("Invalid maximum: ", err)
    }

    shapedMinimum = parsedMinimum
    shapedMaximum = parsedMaximum
  }

  csvReader := csv.NewReader(reader)
  records, err := csvReader.ReadAll()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  // Placeholder variables
  var smallest float64
  var largest float64

  for index, record := range(records) {
    value, _ := strconv.ParseFloat(record[1], 64)
    if index == 0 {
      // Smallest and largest initialized
      smallest = value
      largest = value
    }
    // At end of loop, smallest and largest attain
    // proper values
    if value < smallest {
      smallest = value
    }
    if value > largest {
      largest = value
    }
  }

  // Scaling: shapedMinimum [OFFSET] +  ((x - min) / (max - min) * (shapedMaximum - shapedMinimum))

  var newRecords [][]string

  for _, record := range(records) {
    value, _ := strconv.ParseFloat(record[1], 64)
    name := record[0]
    scaledValue := shapedMinimum + ((value - smallest) / (largest - smallest) * (shapedMaximum - shapedMinimum))
    finalValue := fmt.Sprintf("%.2f", scaledValue)
    var newRecord []string
    newRecord = append(newRecord, name)
    newRecord = append(newRecord, finalValue)
    newRecords = append(newRecords, newRecord)
  }

  writer := csv.NewWriter(os.Stdout)
  writer.WriteAll(newRecords)
}
