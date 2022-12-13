package product

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Entry represents each stream. If the stream fails, an error will be present.
type Entry struct {
	Error   error
	Product Product
	Status  string
}

// Stream helps transmit each streams withing a channel.
type Stream struct {
	stream chan Entry
}

// NewJSONStream returns a new `Stream` type.
func NewJSONStream() Stream {
	return Stream{
		stream: make(chan Entry),
	}
}

// Watch watches JSON streams. Each stream entry will either have an error or a
// User object. Client code does not need to explicitly exit after catching an
// error as the `Start` method will close the channel automatically.
func (s Stream) Watch() <-chan Entry {
	return s.stream
}

// Start starts streaming JSON file line by line. If an error occurs, the channel
// will be closed.
func (s Stream) Start(path string) {
	// Stop streaming channel as soon as nothing left to read in the file.
	//defer close(s.stream)

	s.stream <- Entry{Status: "start"}

	// Open file to read.
	file, err := os.Open(path)
	if err != nil {
		s.stream <- Entry{Error: fmt.Errorf("open file: %w", err)}
		s.stream <- Entry{Status: "error"}
		return
	}
	//defer file.Close not found

	decoder := json.NewDecoder(file)

	// Read opening delimiter. `[` or `{`
	if _, err := decoder.Token(); err != nil {
		s.stream <- Entry{Error: fmt.Errorf("decode opening delimiter: %w", err)}
		s.stream <- Entry{Status: "error"}
		return
	}

	// Read file content as long as there is something.
	//	firstDate := time.Now()
	for decoder.More() {
		var p Product
		if err := decoder.Decode(&p); err != nil {
			s.stream <- Entry{Error: fmt.Errorf("decode line %d: %w", 1, err)}
			s.stream <- Entry{Status: "error"}
			return
		}

		s.stream <- Entry{Product: p}

	}

	// Read closing delimiter. `]` or `}`
	if _, err := decoder.Token(); err != nil {
		s.stream <- Entry{Error: fmt.Errorf("decode closing delimiter: %w", err)}
		s.stream <- Entry{Status: "error"}
		return
	}
	s.stream <- Entry{Status: "finish"}
	s.stream <- Entry{Status: "END"}
	close(s.stream)
	println(time.Now().String())
}
