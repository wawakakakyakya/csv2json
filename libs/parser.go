package libs

import (
	"github.com/wawakakakyakya/csv2json/converter"
	"github.com/wawakakakyakya/csv2json/parser"
)

// []string Block
// [][]string Blocks

func ReadFromFile() {
	lines, err := parser.ReadFromStdin(stdin)
	if err != nil {
		return err
	}
	return converter.ToJson(*lines)
}

func ReadFromStdin() {
	var res []byte
	lines, err = parser.ReadFromFile(filepath)
	if err != nil {
		return err
	}
	converter.ToJson(*lines)
}
