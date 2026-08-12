package components

import (
  "fmt"
)

func Examples() {
  cmd  := Command()
  opt  := Option()
  note := Note()


  // Bridging
  fmt.Printf( "%s\n\n", note("Bridging"))
  fmt.Println("  Suppose we want to get the Bayesian scores of")
  fmt.Println("  a directory with the years, sorted. The following")
  fmt.Println("  command is one approach:")
  fmt.Println()
  fmt.Printf( "  %s %s texts yes.txt no.txt |\n", cmd("gyrate"), opt("--bayes-all"))
  fmt.Printf( "  %s %s data.db bibliography |\n", cmd("gyrate"), opt("--year-bridge"))
  fmt.Printf( "  %s sort -s 1\n", cmd("xan"))
  fmt.Println()
  // Particular bridging example
  fmt.Println("  For my setup, it looks like this to create a line")
  fmt.Println("  chart of the average Bayesian score over time.")
  fmt.Printf( "    %s %s ../../source/txt \\\n", cmd("./gyrate"), opt("--bayes-all"))
  fmt.Println("      ../sc/classifications/explicit.txt \\")
  fmt.Println("      ../sc/classifications/non-explicit.txt |")
  fmt.Printf( "    %s %s ../../db/db.db full |\n", cmd("./gyrate"), opt("--year-bridge"))
  fmt.Printf( "    %s %s -ns 1 |\n", cmd("xan"), opt("sort"))
  fmt.Printf( "    %s %s -n '1' --along-cols '2' 'mean(_)' |\n", cmd("xan"), opt("groupby"))
  fmt.Printf( "    %s %s -n -L '0' '1'\n", cmd("xan"), opt("plot"))
  // For the slower logistic-all
  fmt.Println()
  fmt.Print("  In the case of ", opt("--logistic-all"), ", which is slower,\n")
  fmt.Print("  one can use ", cmd("tee"), " to keep track of progress:\n")
  fmt.Print("    ", cmd("./gyrate"), " ", opt("--logistic-all"), " ../../source/txt |\n")
  fmt.Print("    ", cmd("tee"), " >(", cmd("./gyrate"), " ", opt("--year-bridge"), " ../../db/db.db full |\n")
  fmt.Printf( "    %s %s -ns 1 |\n", cmd("xan"), opt("sort"))
  fmt.Printf( "    %s %s -n '1' --along-cols '2' 'mean(_)' |\n", cmd("xan"), opt("groupby"))
  fmt.Printf( "    %s %s -n -L '0' '1')\n", cmd("xan"), opt("plot"))
  fmt.Println()
  // Normalizing some CSVs
  fmt.Println("  If one wants to get and normalize the output")
  fmt.Println("  of a script by year, something like this would")
  fmt.Println("  be appropriate:")
  fmt.Print("    ", cmd("./gyrate"), " ", opt("--bayes-all"), " \\\n")
  fmt.Print("      ../../source/txt \\\n")
  fmt.Print("      ../sc/classifications/explicit.txt \\\n")
  fmt.Print("      ../sc/classifications/non-explicit.txt |\n")
  fmt.Print("    ", cmd("./gyrate"), " ", opt("--year-bridge"), " ../../db/db.db full |\n")
  fmt.Print("    ", cmd("xan"), " ", opt("sort"), " -ns 1 |\n")
  fmt.Print("    ", cmd("xan"), " ", opt("groupby"), " -n '1' --along-cols '2' 'mean(_)' |\n")
  fmt.Print("    ", cmd("./gyrate"), " ", opt("--normalize"))
  fmt.Println()
  /*
  fmt.Printf(" %s\n\n", Note("Differences"))
  fmt.Println("  To get the differences between two csv files,"
  fmt.Printf("  like those output by %s and %s,\n", opt("--bayes-all"), opt("--logistic-all")
  fmt.Println("  one can use a command like:"
  */
}
