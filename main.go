package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	var src io.Reader
	var err error

	switch len(os.Args) {
	case 2:
		src = os.Stdin
	case 3:
		if src, err = os.Open(os.Args[2]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("missing args")
		os.Exit(1)
	}

	regex, err := regexp.Compile(os.Args[1])
	if err != nil {
		fmt.Printf("bad regex: %s\n", os.Args[1])
		os.Exit(1)
	}

	stdin := bufio.NewScanner(src)
	stdout := bufio.NewWriter(os.Stdout)

	for stdin.Scan() {
		if x := regex.FindString(stdin.Text()); x != "" {
			stdout.WriteString(x)
			stdout.WriteByte(10)
		}
	}

	stdout.Flush()
}
