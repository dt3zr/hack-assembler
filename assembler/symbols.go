package assembler

import (
	"fmt"
)

type symbolID int

const (
	instruction symbolID = iota
	statement
	aStatement
	mStatement
	dStatement
	unaryOperation
	mDest
	dDest
	expression
	amBiOperation
	dBiOperation
	amOperand
	amValOperand
	valOperand
	dValOperand
	amdOperand
	amdValOperand
	jump

	epsilon
	invalidSymbol
)

const (
	terminalStart symbolID = 256 + iota
	areg
	mreg
	dreg
	equal
	plus
	minus
	logicalAnd
	logicalOr
	logicalNot
	zero
	one
	atDigits
	semicolon
	newline
	jumpVerb
	label
	atIdentifier

	eoi
	invalidToken
)

const (
	actionStart symbolID = 512 + iota
	seta                 // a=1
	setc101010           // 0
	setc111111           // 1
	setc111010           // -1
	setc001100           // D
	setc110000           // A or M
	setc001101           // !D
	setc110001           // !A or !M
	setc001111           // -D
	setc110011           // -A or -M
	setc011111           // D+1
	setc110111           // A+1 or M+1
	setc001110           // D-1
	setc110010           // A-1 or M-1
	setc000010           // D+A or D+M
	setc010011           // D-A or D-M
	setc000111           // A-D or M-D
	setc000000           // D&A or D&M
	setc010101           // D|A or D|M
	setd1                // d1=1
	setd2                // d2=1
	setd3                // d3=1
	setj
	setDigits
	setValuePlus
	setValueMinus
	setValueLogicalNot
	actionOnValueWithOne
	actionOnValueWithAMOperand
	actionOnValueWithA
	actionOnValueWithM
	actionOnValueWithD
)

func (sid symbolID) isTerminal() bool {
	return sid > terminalStart
}

func (sid symbolID) isAction() bool {
	return sid > actionStart
}

func (sid symbolID) String() string {
	switch sid {
	case instruction:
		return fmt.Sprintf("instruction")
	case statement:
		return fmt.Sprintf("statement")
	case aStatement:
		return fmt.Sprintf("aStatement")
	case mStatement:
		return fmt.Sprintf("mStatement")
	case dStatement:
		return fmt.Sprintf("dStatement")
	case unaryOperation:
		return fmt.Sprintf("unaryOperation")
	case atDigits:
		return fmt.Sprintf("atDigits")
	case mDest:
		return fmt.Sprintf("mDest")
	case dDest:
		return fmt.Sprintf("dDest")
	case expression:
		return fmt.Sprintf("expression")
	case amBiOperation:
		return fmt.Sprintf("amBiOperation")
	case dBiOperation:
		return fmt.Sprintf("dBiOperation")
	case amOperand:
		return fmt.Sprintf("amOperand")
	case amValOperand:
		return fmt.Sprintf("amValOperand")
	case valOperand:
		return fmt.Sprintf("valOperand")
	case dValOperand:
		return fmt.Sprintf("dValOperand")
	case amdOperand:
		return fmt.Sprintf("amdOperand")
	case amdValOperand:
		return fmt.Sprintf("amdValOperand")
	case jump:
		return fmt.Sprintf("jump")
	case epsilon:
		return fmt.Sprintf("epsilon")
	case invalidSymbol:
		return fmt.Sprintf("invalidSymbol")
	case areg:
		return fmt.Sprintf("areg")
	case mreg:
		return fmt.Sprintf("mreg")
	case dreg:
		return fmt.Sprintf("dreg")
	case equal:
		return fmt.Sprintf("equal")
	case plus:
		return fmt.Sprintf("plus")
	case minus:
		return fmt.Sprintf("minus")
	case logicalAnd:
		return fmt.Sprintf("logicalAnd")
	case logicalOr:
		return fmt.Sprintf("logicalOr")
	case logicalNot:
		return fmt.Sprintf("logicalNot")
	case zero:
		return fmt.Sprintf("zero")
	case one:
		return fmt.Sprintf("one")
	case semicolon:
		return fmt.Sprintf("semicolon")
	case newline:
		return fmt.Sprintf("newline")
	case jumpVerb:
		return fmt.Sprintf("jumpVerb")
	case invalidToken:
		return fmt.Sprintf("invalidToken")
	case eoi:
		return fmt.Sprintf("eoi")
	case seta:
		return fmt.Sprintf("seta")
	case setc101010:
		return fmt.Sprintf("setc101010")
	case setc111111:
		return fmt.Sprintf("setc111111")
	case setc111010:
		return fmt.Sprintf("setc111010")
	case setc001100:
		return fmt.Sprintf("setc001100")
	case setc110000:
		return fmt.Sprintf("setc110000")
	case setc001101:
		return fmt.Sprintf("setc001101")
	case setc110001:
		return fmt.Sprintf("setc110001")
	case setc001111:
		return fmt.Sprintf("setc001111")
	case setc110011:
		return fmt.Sprintf("setc110011")
	case setc011111:
		return fmt.Sprintf("setc011111")
	case setc110111:
		return fmt.Sprintf("setc110111")
	case setc001110:
		return fmt.Sprintf("setc001110")
	case setc110010:
		return fmt.Sprintf("setc110010")
	case setc000010:
		return fmt.Sprintf("setc000010")
	case setc010011:
		return fmt.Sprintf("setc010011")
	case setc000111:
		return fmt.Sprintf("setc000111")
	case setc000000:
		return fmt.Sprintf("setc000000")
	case setc010101:
		return fmt.Sprintf("setc010101")
	case setd1:
		return fmt.Sprintf("setd1")
	case setd2:
		return fmt.Sprintf("setd2")
	case setd3:
		return fmt.Sprintf("setd3")
	case setj:
		return fmt.Sprintf("setj")
	case setDigits:
		return fmt.Sprintf("setDigits")
	case setValuePlus:
		return fmt.Sprintf("setValuePlus")
	case setValueMinus:
		return fmt.Sprintf("setValueMinus")
	case setValueLogicalNot:
		return fmt.Sprintf("setValueLogicalNot")
	case actionOnValueWithOne:
		return fmt.Sprintf("actionOnValueWithOne")
	case actionOnValueWithAMOperand:
		return fmt.Sprintf("actionOnValueWithAMOperand")
	case actionOnValueWithA:
		return fmt.Sprintf("actionOnValueWithA")
	case actionOnValueWithM:
		return fmt.Sprintf("actionOnValueWithM")
	case actionOnValueWithD:
		return fmt.Sprintf("actionOnValueWithD")
	default:
		return ""
	}
}
