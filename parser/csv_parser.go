package parser

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Block []string
type Blocks []Block

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func parseCSV(reader io.Reader) *Blocks {
	var blocks Blocks
	r := csv.NewReader(reader) // csv.NewReaderを使ってcsvを読み込む

	// []stringなのでループする
	for {
		row, err := r.Read() // csvを1行ずつ読み込む
		if err == io.EOF {
			break // 読み込むべきデータが残っていない場合，Readはnil, io.EOFを返すため、ここでループを抜ける
		}
		blocks = append(blocks, Block(row))
	}
	return &blocks
}

func ReadFromFile(path string) (*Blocks, error) {
	var err error
	blocks := new(Blocks)
	if !Exists(path) {
		return blocks, fmt.Errorf("[ERROR] %s does not exists", path)
	}

	file, err := os.Open(path)
	if err != nil {
		return blocks, err
	}
	defer file.Close()

	blocks = parseCSV(file)
	return blocks, err
}

func ReadFromStdin(stdin *bytes.Buffer) (*Blocks, error) {
	var err error
	blocks := parseCSV(stdin)
	return blocks, err
}
