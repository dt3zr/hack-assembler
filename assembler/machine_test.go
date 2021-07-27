package assembler

import (
	"io"
	"strings"
	"testing"
)

func TestNewMachineInstruction(t *testing.T) {

	want := []byte{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	m := newMachineInstruction()

	var i interface{} = m
	mi := i.(*machineInstructionType)

	if len(mi.bits) != 16 {
		t.Fatalf("Expected length to be 16 but got %d\n", len(mi.bits))
	}

	for j := 0; j < len(mi.bits); j++ {
		if (mi.bits)[j] != want[j] {
			t.Fatalf("Expected %v but got %v\n", want, mi)
		}
	}

}

func TestSeta(t *testing.T) {

	want := []byte{1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	m := newMachineInstruction()
	m.seta()

	var i interface{} = m
	mi := i.(*machineInstructionType)

	for j := 0; j < len(mi.bits); j++ {
		if (mi.bits)[j] != want[j] {
			t.Fatalf("Expected %v but got %v\n", want, mi)
		}
	}

}

func TestSetd(t *testing.T) {

	cases := []struct {
		in   int
		want []byte
	}{
		{1, []byte{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}},
		{2, []byte{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0}},
		{3, []byte{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0}},
	}

	for _, c := range cases {

		m := newMachineInstruction()

		switch c.in {
		case 1:
			m.setd1()
		case 2:
			m.setd2()
		default:
			m.setd3()
		}

		var i interface{} = m
		mi := i.(*machineInstructionType)

		for j := 0; j < len(mi.bits); j++ {
			if (mi.bits)[j] != c.want[j] {
				t.Fatalf("Expected %v but got %v\n", c.want, mi.bits)
			}
		}
	}

}

func TestOutputMachineInstruction(t *testing.T) {

	want := "1110000000000000"

	var out strings.Builder
	m := newMachineInstruction()
	src := m.(io.Reader)
	io.Copy(&out, src)
	outString := out.String()

	if outString != want {
		t.Fatalf("Expected %s but got %s\n", want, outString)
	}

}

func TestOutputMachineInstructionWithSmallBuffer(t *testing.T) {

	want := []byte{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	buf := make([]byte, 4)

	m := newMachineInstruction()
	src := m.(io.Reader)

	i := 0

	for {
		n, err := src.Read(buf)

		if err == io.EOF {
			break
		}

		for j := 0; j < n; j++ {
			bit := want[i+j] + 0x30

			// t.Logf("bit=%d, buf=%d, want=%d\n", j, byte(buf[j]), bit)

			if byte(buf[j]) != bit {
				t.Fatalf("Expected %c at bit %d but got %c\n", bit, j, byte(buf[j]))

			}
		}

		i += n
	}

}
