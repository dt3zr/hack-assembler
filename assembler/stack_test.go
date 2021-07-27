package assembler

import "testing"

func TestCreateNewStack(t *testing.T) {
	st := newStack()

	var p interface{} = st
	s := p.(*itemStack)

	if len(s.items) != 0 {
		t.Fatalf("Expected stack is initially 0 length but got %v\n", len(s.items))
	}
}

func TestStackPush(t *testing.T) {
	d := 10
	st := newStack()

	st.push(d)

	var p interface{} = st
	s := p.(*itemStack)

	if len(s.items) != 1 {
		t.Fatalf("Expected stack size to be 1 but got %d\n", len(s.items))
	}
}

func TestStackPop(t *testing.T) {
	d := 10
	st := newStack()

	st.push(d)
	v, err := st.pop()

	var p interface{} = st
	s := p.(*itemStack)

	if v != d || err != nil || len(s.items) != 0 {
		t.Fatalf("Expected popped item to be %d, nil error and length is 0 but some failed\n", d)
	}
}

func TestStackIsEmpty(t *testing.T) {
	st := newStack()

	v, err := st.pop()
	empty := st.isEmpty()

	if v != nil || err == nil || empty != true {
		t.Fatal("Expected stack to be empty")
	}
}

func TestStackPeek(t *testing.T) {
	d := 10
	st := newStack()

	st.push(d)
	v, err := st.peek()

	var p interface{} = st
	s := p.(*itemStack)

	if v != d || err != nil || len(s.items) != 1 {
		t.Fatalf("Expected popped item to be %d, nil error and length is 0 but some failed\n", d)
	}
}
