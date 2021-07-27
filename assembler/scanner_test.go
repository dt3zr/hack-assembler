package assembler

import (
	"bytes"
	"testing"
)

func TestInputSymbolToTransitionIndex(t *testing.T) {
	cases := []struct {
		in   byte
		want byte
	}{
		{'J', 0}, {'G', 1}, {'L', 2}, {'E', 3},
		{'T', 4}, {'Q', 5}, {'N', 6}, {'M', 7},
		{'P', 8}, {'@', 9}, {'0', 10}, {'1', 11},
		{'2', 12}, {'3', 12}, {'4', 12}, {'5', 12},
		{'6', 12}, {'7', 12}, {'8', 12}, {'9', 12},
		{'A', 13}, {'D', 14}, {'=', 15}, {'-', 16},
		{'+', 17}, {'&', 18}, {'|', 19}, {'!', 20},
		{';', 21}, {'\n', 22},
	}

	for _, c := range cases {
		// t.Logf("Getting index for '%v' and index is %v\n", c.in, transitionIndex[c.in])
		if transitionIndex[c.in] != c.want {
			t.Fatalf("Expected %v but got %v\n", c.want, transitionIndex[c.in])
		}
	}
}

func TestSingleToken(t *testing.T) {
	cases := []struct {
		in   string
		want token
	}{
		{"A", token{areg, "A"}},
		{"D", token{dreg, "D"}},
		{"M", token{mreg, "M"}},
		{"=", token{equal, "="}},
	}

	for _, c := range cases {
		buf := bytes.NewBufferString(c.in)
		s := newScanner(buf)
		tk, _ := s.scan()
		if tk.id != c.want.id || tk.lexeme != c.want.lexeme {
			t.Fatalf("Expected %v but got %v\n", c.want, tk)
		}
	}
}

func TestScanNoByte(t *testing.T) {
	buf := bytes.NewBufferString("A")
	s := newScanner(buf)
	tk, _ := s.scan()
	tk, _ = s.scan()
	if tk.id != eoi {
		t.Fatalf("Expected %v but got %v\n", eoi, tk.id)
	}
}

func TestMultipleTokens(t *testing.T) {
	cases := []struct {
		in   string
		want []token
	}{
		{"AM=", []token{{areg, "A"}, {mreg, "M"}, {equal, "="}}},
		{"AD=", []token{{areg, "A"}, {dreg, "D"}, {equal, "="}}},
		{"ADM=", []token{{areg, "A"}, {dreg, "D"}, {mreg, "M"}, {equal, "="}}},
	}

	for _, c := range cases {
		buf := bytes.NewBufferString(c.in)
		s := newScanner(buf)

		for _, w := range c.want {
			tk, _ := s.scan()
			if tk.id != w.id || tk.lexeme != w.lexeme {
				t.Fatalf("Expected %v but got %v\n", w, tk)
			}

		}
	}
}

func TestInvalidTokens(t *testing.T) {
	cases := []struct {
		in   string
		want []symbolID
	}{
		{"AN=0\n", []symbolID{areg, invalidToken}},
		{"AJ=0\n", []symbolID{areg, invalidToken}},
		{"AQI=0\n", []symbolID{areg, invalidToken}},
	}

	for _, c := range cases {
		buf := bytes.NewBufferString(c.in)
		s := newScanner(buf)

		for _, w := range c.want {
			tk, _ := s.scan()
			if tk.id != w {
				t.Fatalf("Expected %v but got %v\n", w, tk.id)
			}

		}
	}
}
