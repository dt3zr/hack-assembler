package assembler

import (
	"errors"
	"io"
)

type machineInstructionType struct {
	bits      []byte
	bytesRead int
}

type instructionSetter interface {
	setValue([]byte) error
	seta()
	setd1()
	setd2()
	setd3()
	setj([]byte) error
	setc(c *[6]byte)
}

func newMachineInstruction() instructionSetter {
	return &machineInstructionType{
		[]byte{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0,
	}
}

func (m *machineInstructionType) setValue(v []byte) error {
	if len(v) < 16 {
		return errors.New("Syntax error")
	}

	for i := 0; i < len(m.bits); i++ {
		m.bits[i] = v[i]
	}

	return nil
}

func (m *machineInstructionType) setc(c *[6]byte) {
	m.bits[9] = c[5]
	m.bits[8] = c[4]
	m.bits[7] = c[3]
	m.bits[6] = c[2]
	m.bits[5] = c[1]
	m.bits[4] = c[0]
}

func (m *machineInstructionType) seta() {
	m.bits[3] = 1
}

func (m *machineInstructionType) setd1() {
	m.bits[10] = 1
}

func (m *machineInstructionType) setd2() {
	m.bits[11] = 1
}

func (m *machineInstructionType) setd3() {
	m.bits[12] = 1
}

func (m *machineInstructionType) setj(j []byte) error {
	if len(j) < 3 {
		return errors.New("Syntax error")
	}

	m.bits[15] = j[2]
	m.bits[14] = j[1]
	m.bits[13] = j[0]

	return nil
}

func (m *machineInstructionType) Read(p []byte) (n int, err error) {

	plen := len(p)

	// if len of input buffer is 0 then do nothing
	if plen <= 0 {
		return 0, nil
	}

	remaining := len(m.bits) - m.bytesRead
	err = io.EOF

	// if input buffer is smaller than the remaining bytes to output use the
	// length of the input buffer otherwise use the length of the remaining
	// bytes
	if n = plen - remaining; n < 0 {
		n = plen
		err = nil
	} else {
		n = remaining
	}

	// each bit is converted to ascii 0 or 1 by adding 0x30 during copy
	for i := 0; i < n; i++ {
		p[i] = m.bits[m.bytesRead+i] + 0x30
	}

	m.bytesRead += n

	return n, err
}
