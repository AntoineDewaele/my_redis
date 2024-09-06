package reader

import (
	"bufio"
	"net"
	"strings"
	"fmt"
	"strconv"
)

// Reader is a struct that holds a connection and a reader
type Reader struct {
	reader *bufio.Reader
}

// New creates a new Reader
func New(conn net.Conn) *Reader {
	return &Reader{
		reader: bufio.NewReader(conn),
	}
}

// Read reads a line from the connection
func (r *Reader) ReadCmd() (string, []string, error) {

	// Read the *<number of elements> line
	line, err := r.readOneLine()

	if strings.HasPrefix(line, "*") {
		cmd, args := r.parseCmd(line)
		return cmd, args, err
	}

	fmt.Println("Received non RESP protocol command")
	var empty []string
	return "", empty, err
}

func (r *Reader) parseCmd(line string) (string, []string) {
	nbElements := r.parseElementsNb(line)
	var cmd string
	var args []string

	for i:= 0; i < nbElements; i++ {
		// skip the $<length> line
		_,_ = r.readOneLine()

		// read the actual element
		line,_ := r.readOneLine()

		if i == 0 {
			cmd = line
		} else {
			args = append(args, line)
		}
	}

	return cmd, args
}

func (r *Reader) readOneLine() (string, error) {
	line, err := r.reader.ReadString('\n')

	if err != nil {
		fmt.Println("Failed to read from connection:", err)
	}

	line = strings.TrimSpace(line)

	return line, err
}

func (r *Reader) parseElementsNb(line string) int {
	nbElements, err := strconv.Atoi(line[1:])
	if err != nil {
		fmt.Println("Failed to parse number of elements:", err)
		return 0
	}

	return nbElements
}