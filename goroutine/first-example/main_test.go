package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_prinSomething(t *testing.T) {
	//This saves the original os.Stdout, so it can be restored later. os.Stdout refers to the standard output stream where a Go program typically writes its output.
	stdOut := os.Stdout

	//This creates a pipe (r for reading and w for writing). The pipe allows the program to capture anything written to it and read that output.
	r, w, _ := os.Pipe()

	//The standard output (os.Stdout) is now replaced by w, which means any output that would normally go to the console will be written to this pipe instead.
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go printSomething("epsilon", &wg)

	wg.Wait()

	_ = w.Close()

	//The program reads everything that was written to the pipe and converts it to a string (output). This is the output of the printSomething function.
	result, _ := io.ReadAll(r)
	output := string(result)

	//The original standard output is restored so that further output goes back to the console.
	os.Stdout = stdOut

	if !strings.Contains(output, "epsilon") {
		t.Error("Epsilon it's not there")
	}
}
