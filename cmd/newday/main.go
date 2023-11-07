package main

import (
	"flag"
	"log"
	"os"

	"github.com/mleone10/advent-of-code-2023/internal/newday"
)

func main() {
	dir := flag.String("dir", ".", "directory in which to generate the new day files")
	day := flag.Int("day", 0, "day number to create [0-99]")
	flag.Parse()

	if *day < 0 || *day > 99 {
		log.Printf("day number must be between 0 and 99 inclusive")
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.Printf("creating new files for day %d in %s", *day, *dir)
	err := newday.Init(*dir, *day)
	if err != nil {
		log.Fatalf("failed to create new files: %v", err)
	}
	log.Println("done!")
}
