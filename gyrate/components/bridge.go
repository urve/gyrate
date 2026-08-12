package components

import (
  "database/sql"
  "encoding/csv"
  _ "github.com/mattn/go-sqlite3"
  "os"
  "fmt"
)

func OpenDB() (*sql.DB, string, [][]string) {
  if len(os.Args) < 4 {
    fmt.Println("Not enough arguments.")
    os.Exit(1)
  }

  db, _ := sql.Open("sqlite3", os.Args[2])

  table := os.Args[3]

  reader := csv.NewReader(os.Stdin)
  records, _ := reader.ReadAll()

  return db, table, records
}

func YearBridge() {
  // gyrate --year-bridge DB TABLE_NAME

  db, table, records := OpenDB()

  defer db.Close()

  for _, record := range records {
    id := record[0]
    val := record[1]
    query := fmt.Sprintf("SELECT year FROM %s WHERE original_id = ?;", table)
    row := db.QueryRow(query, id)
    // gets a single record, but that's OK because
    // original_id are unique in the corpus
    var year string
    row.Scan(&year)
    fmt.Printf("%s,%s,%s\n", id, year, val)
  }
}

