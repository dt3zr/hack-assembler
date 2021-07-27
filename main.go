package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/derktes/hack-assembler/assembler"
)

// Main is entry point of the HACK assembler
func main() {
	var infile, outfile string

	flag.StringVar(&infile, "infile", "", "HACK assembly filename")
	flag.Parse()

	if err := validateInfile(infile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	outfile = makeOutfileName(infile)

	if err := assembler.Assemble(infile, outfile); err != nil {
		fmt.Println(err)
	}
}

func validateInfile(infile string) error {
	lowerInfile := strings.ToLower(infile)

	if lowerInfile == "" {
		return errors.New("No input HACK assembly file available")
	} else if !strings.HasSuffix(lowerInfile, ".asm") {
		return errors.New("Input HACK assembly file name does not end with .asm")
	}

	return nil
}

func makeOutfileName(infile string) string {
	base := strings.TrimSuffix(infile, ".asm")
	return fmt.Sprintf("%s.hack", base)
}
