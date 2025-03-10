package main

import (
	"flag"
	"fmt"
	"os"
	stream "uniq/stream_oper"
	"uniq/uniq"
)

func main() {
	var options uniq.Options
	uniq.Init(&options)
	flag.Parse()

	input, err := stream.GetStream(os.Stdin, 0, os.Open)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to get input stream: %w", err).Error())
		return
	}

	strings, err := stream.ReadLines(input)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to readLines: %w", err).Error())
		return
	}

	result := uniq.Uniq(strings, options)

	output, err := stream.GetStream(os.Stdout, 1, os.Create)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to get output stream: %w", err).Error())
		return
	}

	err = stream.WriteLines(output, result)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to write lines: %w", err).Error())
		return
	}
}
