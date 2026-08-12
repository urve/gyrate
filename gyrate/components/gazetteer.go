package components

import (
  "fmt"
  "os"
  "encoding/csv"
  "slices"
)

/*

I wonder whether this should be done manually,
as I have been doing throughout this code, or
whether I should be using named-entity recognition
methods in Python to determine where place names
occur...

It is likely to me that NER is better for
this sort of thing.

spaCy or (in the absence of spaCy doing a good
job) NLTK seem like good options for this.

*/

func Gazetteer() {
  // gyrate --gazetteer CSV DIRECTORY
  if len(os.Args) < 4 {
    fmt.Println("Not enough arguments.")
  }
  records := GetCSV()
  locations := GetLocations(records)
  uniqueLocations := Unique(locations)
  for _, location := range(uniqueLocations) {
    fmt.Println(location)
  }
}

func GetCSV() [][]string {
  csvFile, err := os.Open(os.Args[2])
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  defer csvFile.Close()
  reader := csvFile

  csvReader := csv.NewReader(reader)
  allRecords, err := csvReader.ReadAll()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  var records [][]string

  for i := range(allRecords) {
    if i != 0 {
      records = append(records, allRecords[i])
    }
  }

  return records
}

func GetLocations(records [][]string) []string {
  var locations []string
  for _, row := range(records) {
    locations = append(locations, row[0])
  }
  return locations 
}

func Unique(locations []string) []string {
  var uniqueLocations []string
  for i := range(locations) {
    var numberSeen int
    for j := range(locations) {
      if i != j && locations[i] == locations[j] {
        numberSeen += 1
      }
    }
    if numberSeen == 0 {
      //fmt.Println(locations[i])
      uniqueLocations = append(uniqueLocations, locations[i])
    }
  }

  slices.Sort(uniqueLocations)

  return uniqueLocations
}
