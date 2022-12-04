package converter

import (
	"encoding/json"
	"fmt"

	"github.com/wawakakakyakya/csv2json/errors"
	"github.com/wawakakakyakya/csv2json/parser"
)

func ToMap(lines parser.Blocks) ([]map[string]string, []errors.CustomError) {

	if lines == nil {
		return []map[string]string{}, []errors.CustomError{}
	}

	keys := lines[0]
	key_length := len(keys)
	var errs []errors.CustomError

	results := []map[string]string{}
	for line_number, line := range lines[1:] {
		dict := map[string]string{}
		for lineIndex, v := range line {
			if len(line) != key_length {
				errs = append(errs, errors.CustomError{Code: errors.InvalidElementError, Err: fmt.Errorf("number of elements in line(%d) and headers elements(%d) differ", line_number+2, key_length)})
			}
			dict[keys[lineIndex]] = v
		}
		results = append(results, dict)
	}

	return results, errs
}

func ToJson(lines parser.Blocks) ([]byte, []errors.CustomError) {
	results, errs := ToMap(lines)
	if errs != nil {
		return []byte{}, errs
	}
	res, err := json.Marshal(results)
	if err != nil {
		errs = append(errs, errors.CustomError{Code: errors.JsonMarshalError, Err: err})
	}
	return res, errs
}
