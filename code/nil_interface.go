package main

import (
	"io"
	"bytes"
)

// Playground : https://play.golang.org/p/QbMpSh_pDHz

const debug = true

func main() {
	// var buf *bytes.Buffer
	
	var buf io.Writer
	
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	
	f(buf) // OK
	
	/*
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	
	f(buf) // NOTE: subtly incorrect!
	
	*/
	
	if debug {
		// ...use buf...
	}
}}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	// ...do something...
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
