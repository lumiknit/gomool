package lexer

import (
	"io"
)

type Lexer struct {
	Filename string
	Reader io.Reader
	Line, Col int
}

func New(filename string, reader io.Reader) *Lexer {
	return &Lexer {
		Filename: filename,
		Reader: reader,
		Line: 1,
		Col: 1 }
}
