package assembler

import (
	"io"
	"os"
	"strings"
)

// Assemble takes an input filename infile and assembles the assembly instructions
// in it and output the hack machine code into the output file outfile
func Assemble(infile, outfile string) error {

	if file, err := os.Open(infile); err == nil {

		var output strings.Builder
		p := newParser(file, &output)
		defer file.Close()

		if err := p.parse(); err == nil {

			ofile, err := os.Create(outfile)
			defer ofile.Close()

			if err == nil {
				reader := strings.NewReader(output.String())
				io.Copy(ofile, reader)
			} else {
				return err
			}

		} else {
			return err
		}

	} else {
		return err
	}

	return nil
}
