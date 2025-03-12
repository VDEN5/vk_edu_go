package main

import (
	"fmt"
	"os"
	stream "uniq/stream_oper"
	"uniq/uniq"
)

func main() {
	var options uniq.Options
	options = uniq.Init()

	input, err := stream.GetStream(os.Stdin, 0, os.Open)
	if err != nil {
		fmt.Printf("failed to get input stream: %w", err)
		return
	}

	strings, err := stream.ReadLines(input)
	if err != nil {
		fmt.Printf("failed to readLines: %w", err)
		return
	}

	output, err := stream.GetStream(os.Stdout, 1, os.Create)
	if err != nil {
		fmt.Printf("failed to get output stream: %w", err)
		return
	}

	result,_ := uniq.Uniq(strings, options)

	err = stream.WriteLines(output, result)
	if err != nil {
		fmt.Printf("failed to write lines: %w", err)
		return
	}
}
