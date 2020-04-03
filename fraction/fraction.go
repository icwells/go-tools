// Defines struct for fraction type

package fraction

import (
	"fmt"
	"math"
)

type Fraction struct {
	Denominator int
	Numerator   int
}

// Returns new fraction with n as numerator and d as denominator. Reduces fraction if possible.
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
	if f.Numerator == v.Numerator && f.Denominator == v.Denominator {
		return true
	}
	return false
}

// Less returns true if f is less than v.
func (f *Fraction) Less(v *Fraction) bool {
	if f.Numerator/f.Denominator < v.Numerator/v.Denominator {
		return true
	}
	return false
}

// Add adds a fraction to f.
func (f *Fraction) Add(v *Fraction) *Fraction {
	n := f.Numerator*v.Denominator + v.Numerator*f.Denominator
	return NewFraction(n, f.Denominator*v.Denominator)
}

// Subtract subtracts a fraction from f.
func (f *Fraction) Subtract(v *Fraction) *Fraction {
	n := f.Numerator*v.Denominator - v.Numerator*f.Denominator
	return NewFraction(n, f.Denominator*v.Denominator)
}

// Multiply multiplies v times f.
func (f *Fraction) Multiply(v *Fraction) *Fraction {
	return NewFraction(f.Numerator*v.Numerator, f.Denominator*v.Denominator)
}

// Divide divides f by v.
func (f *Fraction) Divide(v *Fraction) *Fraction {
	return NewFraction(f.Numerator*v.Denominator, f.Denominator*v.Numerator)
}
