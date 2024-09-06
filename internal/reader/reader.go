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
func (r *Reader) ReadCmd() (string, error) {

	// Read the *<number of elements> line
	line, err := r.readOneLine()

	if strings.HasPrefix(line, "*") {
		return r.parseCmd(line), err
	}

	fmt.Println("Received non RESP protocol command")
	return "", err
}

func (r *Reader) parseCmd(line string) string {
	nbElements := r.parseElementsNb(line)
	cmd := ""

	for i:= 0; i < nbElements; i++ {
		// skip the $<length> line
		_,_ = r.readOneLine()
		// read the actual element
		line,_ := r.readOneLine()
		cmd = cmd + " " + line
	}

	return cmd
}

func (r *Reader) readOneLine() (string, error) {
	line, err := r.reader.ReadString('\n')

	if err != nil {
		fmt.Println("Failed to read from connection:", err)
	}

	return strings.TrimSpace(line), err
}

func (r *Reader) parseElementsNb(line string) int {
	nbElements, err := strconv.Atoi(line[1:])
	if err != nil {
		fmt.Println("Failed to parse number of elements:", err)
		return 0
	}

	return nbElements
}