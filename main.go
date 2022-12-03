package main

import (
	"fmt"
	"os"

	args "github.com/wawakakakyakya/csv2json/args"
	"github.com/wawakakakyakya/csv2json/converter"
	"github.com/wawakakakyakya/csv2json/logger"
	"github.com/wawakakakyakya/csv2json/parser"
)

func main() {
	logger := logger.NewLogger("main")
	filepath, stdin, err := args.GetArgs()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("remove after %s", stdin))
	lines, err := parser.ReadFromStdin(stdin)
	if err != nil {
		logger.Error(err.Error())
	}
	lines, err = parser.ReaFromFile(filepath)
	if err != nil {
		logger.Error(err.Error())
	}
	converter.ToJson(*lines)
}
