package components

import (
  "fmt"
  "os"
  "encoding/csv"
  "strconv"
  "io"
)

func Normalize() {
  if len(os.Args) < 2 {
    fmt.Println("Not enough arguments.")
    os.Exit(1)
  }

  var reader io.Reader

  // Either read a file path or STDIN
  if len(os.Args) > 2 {
    csvFile, err := os.Open(os.Args[2])
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
    defer csvFile.Close()
    reader = csvFile
  } else {
    reader = os.Stdin
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

  // Scaling: 100 * (x - min) / (max - min)

  var newRecords [][]string

  for _, record := range(records) {
    value, _ := strconv.ParseFloat(record[1], 64)
    name := record[0]
    scaledValue := 100 * ((value - smallest) / (largest - smallest))
    finalValue := fmt.Sprintf("%.2f", scaledValue)
    var newRecord []string
    newRecord = append(newRecord, name)
    newRecord = append(newRecord, finalValue)
    newRecords = append(newRecords, newRecord)
  }

  writer := csv.NewWriter(os.Stdout)
  writer.WriteAll(newRecords)
}
