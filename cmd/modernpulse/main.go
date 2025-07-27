// cmd/modernpulse/main.go
package main

import (
	"flag"
	"log"
	"os"

	"modernpulse/internal/modernpulse"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := modernpulse.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
