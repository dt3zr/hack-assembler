# HACK Assembler

## What's HACK?

HACK is a simple and interesting computer system built from scratch for the purpose of teaching CS students the [concepts of computer system from first principle in a single book](https://www.nand2tetris.org/).

## What's HACK Assembler

Like any other modern computer system, HACK computer system runs programs written by software engineers. In HACK, programs can be written in HACK assembly language and assembled by the HACK assembler into HACK machine instruction that can be executed by the HACK CPU.

## How does it work?

The HACK assembler is written in [Go](https://golang.org/). To build the assembler, run the usual build command.

```shell
go build
```

Then run the assembler with the provided HACK assembly source code.

```shell
./hack-assembler -infile draw.asm
```

A HACK machine binary file `draw.hack` is created and the file can be loaded into the [CPU Emulator](https://www.nand2tetris.org/software) and executed by the HACK CPU. That's it!

