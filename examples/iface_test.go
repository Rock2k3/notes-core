package examples

import "testing"

func TestSimple(t *testing.T) {
	ue := &UseExample{
		X: 5,
		Y: 6,
	}

	ca := &SimpleAdder{}

	res := ue.UseAdder(ca)
	if res != 10 {
		t.Errorf("expected = %v, actual = %v", 11, res)
	}

}

func TestComplex(t *testing.T) {

	ue := &UseExample{
		X: 5,
		Y: 6,
	}

	ca := &ComplexAdder{}

	res := ue.UseAdder(ca)
	if res != 10 {
		t.Errorf("expected = %v, actual = %v", 11, res)
	}

}
