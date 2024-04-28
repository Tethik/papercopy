package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"flag"

	"github.com/Tethik/go-template/internal/niceware"
)

var versionFlag = flag.Bool("v", false, "Print version information and quit")
var versionString = fmt.Sprintf("niceware %s, commit %s, build %s", version, commit, build)
var decodeFlag = flag.Bool("d", false, "Decode words to bytes")

var (
	version string
	commit  string
	build   string
)

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Println(versionString)
		return
	}

	input, err := io.ReadAll(os.Stdin)

	if *decodeFlag {
		bytes, err := niceware.WordsToBytes(strings.Fields(string(input)))
		if err != nil {
			fmt.Printf("could not convert words to bytes: %v", err)
			return
		}
		os.Stdout.Write(bytes)
		return
	}

	if err != nil {
		fmt.Printf("could not read input: %v", err)
		return
	}

	words, err := niceware.BytesToWords(input)
	if err != nil {
		fmt.Printf("could not convert bytes to words: %v", err)
		return
	}

	fmt.Println(strings.Join(words, " "))
}
