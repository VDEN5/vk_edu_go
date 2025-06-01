package streamOperations

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Oper func(string) (*os.File, error)

/*помощники в обработке потока*/
func GetStream(defaultValue *os.File, numArg int, operation Oper) (*os.File, error) {
	var err error
	stream, filename := defaultValue, flag.Arg(numArg)
	if filename != "" {
		stream, err = operation(filename)
		if err != nil {
			return nil, fmt.Errorf("failed to open/create file: %w", err)
		}
		defer stream.Close()
	}
	return stream, nil
}

func ReadLines(input *os.File) ([]string, error) {
	scanner, lines := bufio.NewScanner(input), make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func WriteLines(output *os.File, strings []string) error {
	writer := bufio.NewWriter(output)
	for _, line := range strings {
		fmt.Fprintln(writer, line)
	}
	return writer.Flush()
}
