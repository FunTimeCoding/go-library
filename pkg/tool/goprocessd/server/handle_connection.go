package server

import (
	"bufio"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system/writer"
	"net"
)

func (s *Server) handleConnection(connection net.Conn) {
	defer errors.PanicClose(connection)
	scanner := bufio.NewScanner(connection)

	if !scanner.Scan() {
		return
	}

	response := s.handle(scanner.Text())
	writer.Print(connection, "%s\n.\n", response)
}
