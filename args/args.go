package args

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"golang.org/x/term"
)

var (
	lf string = "\n"
)

func init() {
	flag.Parse()
}

func readStdin() (*bytes.Buffer, error) {
	var err error
	buf := bytes.NewBufferString("")

	if term.IsTerminal(0) { // end if not used pipe
		return buf, err
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		err := scanner.Err()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return buf, err
		}

		txt := scanner.Text()
		if txt == "" {
			continue
		}
		buf.WriteString(txt + lf)

	}

	return buf, err
}

func GetArgs() (string, *bytes.Buffer, error) {
	var err error
	var filePath string
	args := flag.Args()
	if len(args) > 0 {
		filePath = args[0]
	}
	stdin, err := readStdin()
	return filePath, stdin, err
}
