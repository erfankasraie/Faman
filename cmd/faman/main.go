package main

import (
	"os"

	"github.com/faman-project/faman/internal/app"
)

func main() {
	if err := app.Execute(); err != nil {
		os.Exit(1)
	}
}
