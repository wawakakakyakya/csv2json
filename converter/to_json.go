package converter

import (
	"encoding/json"
	"fmt"

	"github.com/wawakakakyakya/csv2json/parser"
)

func ToJson(lines parser.Blocks) {
	keys := lines[0]
	key_length := len(keys)

	dict := map[string]string{}
	for line_number, line := range lines[1:] {
		for lineIndex, v := range line {
			if len(line) != key_length {
				fmt.Printf("Number of elements in line(%d) and headers elements(%d) differ", line_number+1, key_length)
				fmt.Println(line)
				break
			}
			dict[keys[lineIndex]] = v
		}
	}
	s, err := json.Marshal(dict)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(s))
}
