package fraction

import (
	"testing"
)

func TestFraction(t *testing.T) {
	a := NewFraction(1, 3)
	b := NewFraction(3, 28)
	c := NewFraction(1, 3)
	if a.String() != "1/3" {
		t.Errorf("Actual fraction %s does not equal 1/3.", a.String())
	} else if b.String() != "3/28" {
		t.Errorf("Actual fraction %s does not equal 3/28.", b.String())
	} else if a.Equals(b) {
		t.Errorf("%s does not equal %s.", a.String(), b.String())
	} else if !a.Equals(c) {
		t.Errorf("%s equals %s.", a.String(), c.String())
	} else if a.Less(b) {
		t.Errorf("%s is greater than %s.", a.String(), c.String())
	} else {
		c = a.Add(b)
		if c.String() != "37/84" {
			t.Errorf("%s + %s does not equal 37/84.", a.String(), b.String())
		}
		c = a.Subtract(b, false)
		if c.String() != "19/84" {
			t.Errorf("%s - %s does not equal 19/84.", a.String(), b.String())
		}
		c = a.Multiply(b)
		if c.String() != "1/28" {
			t.Errorf("%s * %s does not equal 1/28.", a.String(), b.String())
		}
		c = a.Divide(b)
		if c.String() != "28/9" {
			t.Errorf("%s / %s does not equal 28/9.", a.String(), b.String())
		}
		c = a.Multiply(NewFraction(-1, 1))
		if c.String() != "-1/3" {
			t.Errorf("%s * %s does not equal -1/3.", a.String(), b.String())
		}
	}
}
