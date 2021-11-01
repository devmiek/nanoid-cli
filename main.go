// Copyright (c) 2021 Handle
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"flag"
	"fmt"
	"io"
	"runtime"

	"github.com/nobody-night/nanoid-go"
)

// versionTag represents the version tag of nanoid-cli.
const versionTag string = "0.1.1"

var (
	// generateCount represents the number of Nano ID to be generated.
	generateCount int

	// generateSize represents the size of the Nano ID to be generated.
	generateSize int

	// generateAlphabet represents the alphabet to be used to generate the Nano ID.
	generateAlphabet string

	// outputVersionOnly represents whether to output version information only.
	outputVersionOnly bool
)

func init() {
	flag.IntVar(&generateCount, "n", 1, "number of Nano IDs to be generated")
	flag.IntVar(&generateSize, "s", 21, "size of Nano IDs to be generated")
	flag.StringVar(&generateAlphabet, "a", "", fmt.Sprintf(
		"custom alphabet used to generate Nano IDs (required <= %d)",
		nanoid.MaxAlphabetSize))
	flag.BoolVar(&outputVersionOnly, "v", false, "output version information only")
	flag.Parse()
}

// createReader creates a Nano ID reader with the command line options,
// then returns the created Nano ID reader and any errors encountered.
func createReader() (r io.Reader, err error) {
	switch len(generateAlphabet) {
	case 0:
		return nanoid.NewReader()
	default:
		return nanoid.NewReader(nanoid.WithAlphabet(generateAlphabet))
	}
}

func main() {
	if outputVersionOnly {
		fmt.Printf("nanoid version %s (runtime %s)\n", versionTag, runtime.Version())
		return
	}

	if generateCount < 1 {
		fmt.Printf("nanoid: invalid generate count: %d\n", generateCount)
		return
	}
	if generateSize < 1 {
		fmt.Printf("nanoid: invalid generate size: %d\n", generateSize)
		return
	}
	if len(generateAlphabet) > nanoid.MaxAlphabetSize {
		fmt.Printf("nanoid: alphabet is too long: %d\n", len(generateAlphabet))
		return
	}

	reader, err := createReader()
	if err != nil {
		fmt.Printf("nanoid: %s\n", err)
		return
	}

	buf := make([]byte, generateSize)
	for count := 0; count < generateCount; count++ {
		_, err = reader.Read(buf)
		if err != nil {
			fmt.Printf("nanoid: %s\n", err)
			return
		}
		fmt.Printf("%s\n", buf)
	}
}
