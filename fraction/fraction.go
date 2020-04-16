// fraction provides a struct to store fractions and provides mathmatical and conversion methods.

package fraction

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Fraction struct stores the numerator and denominator of a fraction as integers.
type Fraction struct {
	Denominator int
	Numerator   int
}

// NewFraction returns new fraction with n as numerator and d as denominator. Reduces fraction if possible.
// Returns an empty fraction (0/1) if d is less than 1.
func NewFraction(n, d int) *Fraction {
	f := new(Fraction)
	f.Denominator = 1
	if d > 0 {
		gcd := f.greatestCommonDenominator(n, d)
		f.Numerator = int(float64(n) / gcd)
		f.Denominator = int(float64(d) / gcd)
	}
	return f
}

// Zero returns an empty fraction equal to 0/1.
func Zero() *Fraction {
	return NewFraction(0, 1)
}

// greatestCommonDenominator calculates the gratest common denominator of n and d using Euchlid's algorithm.
func (f *Fraction) greatestCommonDenominator(n, d int) float64 {
	n = int(math.Abs(float64(n)))
	for n != d {
		if n > d {
			n -= d
		} else {
			d -= n
		}
		if n <= 0 || d <= 0 {
			n = 1
			break
		}
	}
	return float64(n)
}

// FromFloat converts a float64 to a fraction.
func FromFloat(n float64) *Fraction {
	d := 1
	// Get number of digits after decimal
	l := len(strings.Split(strconv.FormatFloat(n, 'f', 64, -1), ".")[1])
	return NewFraction(int(n*float64(l)), d*l)
}

// Copy returns a deep copy of f.
func (f *Fraction) Copy() *Fraction {
	return NewFraction(f.Numerator, f.Denominator)
}

// String returns fraction formatted as a string.
func (f *Fraction) String() string {
	return fmt.Sprintf("%d/%d", f.Numerator, f.Denominator)
}

// Float returns numerator/denominator.
func (f *Fraction) Float() float64 {
	return float64(f.Numerator) / float64(f.Denominator)
}

// Equals returns true if two fractions are equal.
func (f *Fraction) Equals(v *Fraction) bool {
	if f.Float() == v.Float() {
		return true
	}
	return false
}

// Less returns true if f is less than v.
func (f *Fraction) Less(v *Fraction) bool {
	if f.Float() < v.Float() {
		return true
	}
	return false
}

//---------------Arithmetic Operations----------------------------------------

// Add adds a fraction to f.
func (f *Fraction) Add(v *Fraction) *Fraction {
	n := f.Numerator*v.Denominator + v.Numerator*f.Denominator
	return NewFraction(n, f.Denominator*v.Denominator)
}

// AddInt adds n/d to f.
func (f *Fraction) AddInt(n, d int) *Fraction {
	return f.Add(NewFraction(n, d))
}

// AddFloat adds a floating point number to f.
func (f *Fraction) AddFloat(n float64) *Fraction {
	return FromFloat(f.Float() + n)
}

// Subtract subtracts a fraction from f. Returns the absolute value of the result if abs is true.
func (f *Fraction) Subtract(v *Fraction, abs bool) *Fraction {
	n := f.Numerator*v.Denominator - v.Numerator*f.Denominator
	if abs {
		n = int(math.Abs(float64(n)))
	}
	return NewFraction(n, f.Denominator*v.Denominator)
}

// SubtractInt subtracts n/d from f. Returns the absolute value of the result if abs is true.
func (f *Fraction) SubtractInt(n, d int, abs bool) *Fraction {
	return f.SubtractFloat(float64(n)/float64(d), abs)
}

// SubtractFloat subtracts a floating point number from f. Returns the absolute value of the result if abs is true.
func (f *Fraction) SubtractFloat(n float64, abs bool) *Fraction {
	v := f.Float() - n
	if abs {
		v = math.Abs(v)
	}
	return FromFloat(v)
}

// Multiply multiplies v times f.
func (f *Fraction) Multiply(v *Fraction) *Fraction {
	return NewFraction(f.Numerator*v.Numerator, f.Denominator*v.Denominator)
}

// MultiplyInt multiplies f by n/d.
func (f *Fraction) MultiplyInt(n, d int) *Fraction {
	return f.Multiply(NewFraction(n, d))
}

// MultiplyFloat multiplies f by a floating point number.
func (f *Fraction) MultiplyFloat(n float64) *Fraction {
	return FromFloat(f.Float() * n)
}

// Divide divides f by v.
func (f *Fraction) Divide(v *Fraction) *Fraction {
	return NewFraction(f.Numerator*v.Denominator, f.Denominator*v.Numerator)
}

// DivideInt divides f by n/d.
func (f *Fraction) DivideInt(n, d int) *Fraction {
	return f.Divide(NewFraction(n, d))
}

// DivideFloat divides f by a floating point number.
func (f *Fraction) DivideFloat(n float64) *Fraction {
	return FromFloat(f.Float() / n)
}
