package client

import (
	"bufio"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"net"
	"strings"
)

func Send(
	socketPath string,
	command string,
	arguments []string,
) (string, error) {
	connection, e := net.Dial("unix", socketPath)

	if e != nil {
		return "", fmt.Errorf("connect: %w", e)
	}

	defer errors.PanicClose(connection)
	line := command

	if len(arguments) > 0 {
		line = fmt.Sprintf("%s %s", command, strings.Join(arguments, " "))
	}

	_, e = fmt.Fprintln(connection, line)
	errors.PanicOnError(e)
	scanner := bufio.NewScanner(connection)
	var lines []string

	for scanner.Scan() {
		text := scanner.Text()

		if text == "." {
			break
		}

		lines = append(lines, text)
	}

	if len(lines) == 0 {
		return "", fmt.Errorf("no response from server")
	}

	return strings.Join(lines, "\n"), nil
}
