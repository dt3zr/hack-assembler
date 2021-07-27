package assembler

import (
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
)

func init() {
	// log.SetOutput(nullWriter)
}

type productionID int
type production []symbolID
type productionList [][]symbolID

var pList productionList = productionList{
	{statement, jump, newline, instruction},      // 0 instruction
	{epsilon},                                    // 1 instruction
	{areg, aStatement},                           // 2 statement
	{mreg, mStatement},                           // 3 statement
	{dreg, dStatement},                           // 4 statement
	{unaryOperation},                             // 5 statement
	{setDigits, atDigits},                        // 6 statement
	{mDest, dDest, equal, setd1, expression},     // 7 aStatement
	{amBiOperation},                              // 8 aStatement
	{mreg, setd3},                                // 9 mDest
	{epsilon},                                    // 10 mDest
	{dreg, setd2},                                // 11 dDest
	{epsilon},                                    // 12 dDest
	{dDest, equal, setd3, expression},            // 13 mStatement
	{amBiOperation, seta},                        // 14 mStatement
	{equal, setd2, expression},                   // 15 dStatement
	{dBiOperation},                               // 16 dStatement
	{areg, amBiOperation},                        // 17 expression
	{mreg, amBiOperation, seta},                  // 18 expression
	{dreg, dBiOperation},                         // 19 expression
	{unaryOperation},                             // 20 expression
	{plus, valOperand, setc110111},               // 21 amBiOperation
	{minus, dValOperand},                         // 22 amBiOperation
	{epsilon, setc110000},                        // 23 amBiOperation
	{plus, setValuePlus, amValOperand},           // 24 dBiOperation
	{minus, setValueMinus, amValOperand},         // 25 dBiOperation
	{logicalAnd, amOperand, setc000000},          // 26 dBiOperation
	{logicalOr, amOperand, setc010101},           // 27 dBiOperation
	{epsilon, setc001100},                        // 28 dBiOperation
	{one},                                        // 29 valOperand
	{one, setc110010},                            // 30 dValOperand
	{dreg, setc000111},                           // 31 dValOperand
	{one, actionOnValueWithOne},                  // 32 amValOperand
	{amOperand, actionOnValueWithAMOperand},      // 33 amValOperand
	{areg},                                       // 34 amOperand
	{mreg, seta},                                 // 35 amOperand
	{minus, setValueMinus, amdValOperand},        // 36 unaryOperation
	{logicalNot, setValueLogicalNot, amdOperand}, // 37 unaryOperation
	{zero, setc101010},                           // 38 unaryOperation
	{one, setc111111},                            // 39 unaryOperation
	{one, setc111010},                            // 40 amdValOperand
	{amdOperand},                                 // 41 amdValOperand
	{areg, actionOnValueWithA},                   // 42 amdOperand
	{mreg, actionOnValueWithM},                   // 43 amdOperand
	{dreg, actionOnValueWithD},                   // 44 amdOperand
	{semicolon, setj, jumpVerb},                  // 45 jump
	{epsilon},                                    // 46 jump
}

type parseTable [][18]productionID

var pTable = parseTable{
	//  a     m     d     =     +     -     &     |     !     0     1     @     ;    \n   jxx   lbl   @id   eoi
	{0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0x01, 0xff, 0xff, 0xff, 0x01}, // instruction
	{0x02, 0x03, 0x04, 0xff, 0xff, 0x05, 0xff, 0xff, 0x05, 0x05, 0x05, 0x06, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, // statement
	{0xff, 0x07, 0x07, 0x07, 0x08, 0x08, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x08, 0x08, 0xff, 0xff, 0xff, 0xff}, // aStatement
	{0xff, 0xff, 0x0d, 0x0d, 0x0e, 0x0e, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x0e, 0x0e, 0xff, 0xff, 0xff, 0xff}, // mStatement
	{0xff, 0xff, 0xff, 0x0f, 0x10, 0x10, 0x10, 0x10, 0xff, 0xff, 0xff, 0xff, 0x1c, 0x1c, 0xff, 0xff, 0xff, 0xff}, // dStatement
	{0xff, 0xff, 0xff, 0xff, 0xff, 0x24, 0xff, 0xff, 0x25, 0x26, 0x27, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, // unaryOperation
	{0xff, 0x09, 0x0a, 0x0a, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, // mDest
	{0xff, 0xff, 0x0b, 0x0c, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, // dDest
	{0x11, 0x12, 0x13, 0xff, 0xff, 0x14, 0xff, 0xff, 0x14, 0x14, 0x14, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, // expression
	{0xff, 0xff, 0xff, 0xff, 0x15, 0x16, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x17, 0x17, 0xff, 0xff, 0xff, 0xff}, // amBiOperation
	{0xff, 0xff, 0xff, 0xff, 0x18, 0x19, 0x1a, 0x1b, 0xff, 0xff, 0xff, 0xff, 0x1c, 0x1c, 0xff, 0xff, 0xff, 0xff}, // dBiOperation
	{0x22, 0x23, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, // amOperand
	{0x21, 0x21, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x20, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, // amValOperand
	{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1d, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, // valOperand
	{0xff, 0xff, 0x1f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1e, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, // dValOperand
	{0x2a, 0x2b, 0x2c, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, // amdOperand
	{0x29, 0x29, 0x29, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x28, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, // amdValOperand
	{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x2d, 0x2e, 0xff, 0xff, 0xff, 0xff}, // jump
}

type actionHandler func(setter instructionSetter, vs stack, tk *token, tv *value)

var actionHandlerList = []actionHandler{
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // seta
		setter.seta()
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setc101010
		setter.setc(&[6]byte{1, 0, 1, 0, 1, 0})
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setc111111
		setter.setc(&[6]byte{1, 1, 1, 1, 1, 1})
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setc111010
		setter.setc(&[6]byte{1, 1, 1, 0, 1, 0})
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setc001100
		setter.setc(&[6]byte{0, 0, 1, 1, 0, 0})
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setc110000
		setter.setc(&[6]byte{1, 1, 0, 0, 0, 0})
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setc001101
		setter.setc(&[6]byte{0, 0, 1, 1, 0, 1})
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setc110001
		setter.setc(&[6]byte{1, 1, 0, 0, 0, 1})
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setc001111
		setter.setc(&[6]byte{0, 0, 1, 1, 1, 1})
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setc110011
		setter.setc(&[6]byte{1, 1, 0, 0, 1, 1})
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setc011111
		setter.setc(&[6]byte{0, 1, 1, 1, 1, 1})
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setc110111
		setter.setc(&[6]byte{1, 1, 0, 1, 1, 1})
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setc001110
		setter.setc(&[6]byte{0, 0, 1, 1, 1, 0})
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setc110010
		setter.setc(&[6]byte{1, 1, 0, 0, 1, 0})
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setc000010
		setter.setc(&[6]byte{0, 0, 0, 0, 1, 0})
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setc010011
		setter.setc(&[6]byte{0, 1, 0, 0, 1, 1})
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setc000111
		setter.setc(&[6]byte{0, 0, 0, 1, 1, 1})
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setc000000
		setter.setc(&[6]byte{0, 0, 0, 0, 0, 0})
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setc010101
		setter.setc(&[6]byte{0, 1, 0, 1, 0, 1})
	},

	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setd1
		setter.setd1()
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setd2
		setter.setd2()
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setd3
		setter.setd3()
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setj
		switch tk.lexeme {
		case "JGE":
			setter.setj([]byte{0, 1, 1})
		case "JGT":
			setter.setj([]byte{0, 0, 1})
		case "JLE":
			setter.setj([]byte{1, 1, 0})
		case "JLT":
			setter.setj([]byte{1, 0, 0})
		case "JEQ":
			setter.setj([]byte{0, 1, 0})
		case "JNE":
			setter.setj([]byte{1, 0, 1})
		case "JMP":
			setter.setj([]byte{1, 1, 1})
		}
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setDigits
		n, _ := strconv.Atoi(string(tk.lexeme[1:]))
		b := []byte(fmt.Sprintf("%016b", n))
		for i := 0; i < len(b); i++ {
			b[i] = b[i] - 0x30
		}
		setter.setValue(b)
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setValuePlus
		i, _ := vs.pop()
		v := i.(value)
		v.right = plus
		vs.push(v)
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setvalueMinus
		i, _ := vs.pop()
		v := i.(value)
		v.right = minus
		vs.push(v)
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // setValueLogicalNot
		i, _ := vs.pop()
		v := i.(value)
		v.right = logicalNot
		vs.push(v)
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // actionOnValueWithOne
		if tv.right == plus {
			setter.setc(&[6]byte{0, 1, 1, 1, 1, 1})
		} else {
			setter.setc(&[6]byte{0, 0, 1, 1, 1, 0})
		}
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // actionOnValueWithAMOperand
		if tv.right == plus {
			setter.setc(&[6]byte{0, 0, 0, 0, 1, 0})
		} else {
			setter.setc(&[6]byte{0, 1, 0, 0, 1, 1})
		}
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // actionOnValueWithA
		if tv.right == minus {
			setter.setc(&[6]byte{1, 1, 0, 0, 1, 1})
		} else {
			setter.setc(&[6]byte{1, 1, 0, 0, 0, 1})
		}
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // actionOnValueWithM
		if tv.right == minus {
			setter.setc(&[6]byte{1, 1, 0, 0, 1, 1})
		} else {
			setter.setc(&[6]byte{1, 1, 0, 0, 0, 1})
		}
		setter.seta()
	},
	func(setter instructionSetter, vs stack, tk *token, tv *value) { // actionOnValueWithD
		if tv.right == minus {
			setter.setc(&[6]byte{0, 0, 1, 1, 1, 1})
		} else {
			setter.setc(&[6]byte{0, 0, 1, 1, 0, 1})
		}
	},
}

func getProduction(s symbolID, t symbolID) (production, error) {
	ns, nt := int(s), int(t-terminalStart-1)
	log.Printf("Obtaining production id with %v, %v\n", ns, nt)

	id := pTable[ns][nt]

	if id == 255 {
		return nil, errors.New("Syntax error")
	}

	return pList[id], nil
}

type value struct {
	left  symbolID
	right symbolID
}

type parser struct {
	scanner
	writer io.Writer
}

func newParser(reader io.Reader, writer io.Writer) *parser {
	scanner := newScanner(reader)
	return &parser{scanner, writer}
}

func (p *parser) parse() error {
	ps := newStack()
	vs := newStack()

	ps.push(instruction)
	vs.push(value{invalidToken, invalidToken})
	machineInstruction := newMachineInstruction()

	tk, _ := p.scan()

	for !ps.isEmpty() {

		log.Printf("Lookahead token is %v", tk)

		if tk.id == invalidToken {
			return errors.New("Syntax error")
		}

		top, _ := ps.pop()
		topSymbol := top.(symbolID)
		top, _ = vs.pop()
		topValue := top.(value)

		log.Printf("Top symbol is %v", topSymbol)

		if topSymbol.isAction() {

			log.Printf("Top symbol %v is an action", topSymbol)
			actionHandlerList[int(topSymbol-actionStart-1)](machineInstruction, vs, &tk, &topValue)

		} else if topSymbol.isTerminal() {

			if topSymbol == tk.id {
				// if the top symbol is the same as the scanned input token
				// then the top symbol is a terminal symbol, a perfect match, go pop

				log.Printf("%v is equal to %v, advancing to next token", tk, topSymbol)

				if tk.id == newline {

					log.Printf("Machine instruction is %v", machineInstruction)

					r := machineInstruction.(io.Reader)
					io.Copy(p.writer, r)
					p.writer.Write([]byte{'\n'})

					machineInstruction = newMachineInstruction()

				}

				tk, _ = p.scan()

			} else {
				// if topSymbol is terminal symbol but did not match any input token
				// then it's definitely a syntax error
				return errors.New("Syntax error")
			}

		} else {

			// the token is now a non-terminal

			if topSymbol == epsilon {

				// if there is an epsilon on the stack, go pop another symbol
				log.Printf("Meet %v, do nothing", topSymbol)

			} else {

				prod, err := getProduction(topSymbol, tk.id)
				log.Printf("Production to apply is %v", prod)

				// when err is not nil, it's definitely a syntax error, return the error
				if err != nil {
					return err
				}

				lhs := topValue.right

				for i := len(prod) - 1; i >= 0; i-- {
					log.Printf("Pushing %v", prod[i])
					ps.push(prod[i])
					vs.push(value{lhs, lhs})
				}
			}
		}

	}

	return nil
}
