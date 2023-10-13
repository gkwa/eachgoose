package main

import (
	"log/slog"
	"os"

	"github.com/taylormonacelli/eachgoose"
	"github.com/taylormonacelli/goldbug"
)

func main() {
	goldbug.SetDefaultLoggerText(slog.LevelDebug)

	status := eachgoose.Execute()
	os.Exit(status)
}
