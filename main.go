package main

import (
	"bytes"
	"fmt"
	"os"

	args "github.com/wawakakakyakya/csv2json/args"
	"github.com/wawakakakyakya/csv2json/converter"
	"github.com/wawakakakyakya/csv2json/logger"
	"github.com/wawakakakyakya/csv2json/parser"
)

func read(filePath string, stdin *bytes.Buffer) (*parser.Blocks, error) {
	var lines *parser.Blocks
	var err error
	if filePath != "" {
		lines, err = parser.ReadFromFile(filePath)

	} else if stdin != nil {
		lines, err = parser.ReadFromStdin(stdin)
	}
	return lines, err
}

func main() {
	logger := logger.NewLogger("main")
	filepath, stdin, errs := args.GetArgs()
	if len(errs) > 0 {
		for _, e := range errs {
			logger.Error(e.Error())
		}
		os.Exit(1)
	}

	lines, err := read(filepath, stdin)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	jsons, custom_errs := converter.ToJson(*lines)
	if len(custom_errs) > 0 {
		for _, err := range errs {
			logger.Error(err.Error())
		}
		os.Exit(1)
	}

	fmt.Println(string(jsons))

}
