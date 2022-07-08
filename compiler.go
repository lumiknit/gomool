package main

import (
	"io"
	"log"
	"os"
	"sync"

	"github.com/lumiknit/gomool/lexer"
)

func compile(wg *sync.WaitGroup, filename string) {
	defer (*wg).Done()
	log.Printf("[%s] Compile...\n", filename)

	var err error
	var reader io.Reader

	// Read File
	reader, err = os.Open(filename)
	if err != nil {
		log.Fatalf("[%s] Cannot open file!\n", filename)
		return
	}

	lexer := lexer.New(filename, reader)
	println(lexer)

	log.Printf("[%s] Done!\n", filename)
}
