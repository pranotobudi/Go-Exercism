// Package triangle will handle functions related to triangle
package triangle

import "math"

// Kind is an integer type which is correspondence to Nat, Equ, Iso, Sca constants
type Kind int

const (
	// We'll use integer for the following identifiers used by the test program.
	NaT = 0 // not a triangle
	Equ = 1 // equilateral
	Iso = 2 // isosceles
	Sca = 3 // scalene
)

// KindFromSides receive 3 float64 number and return Kind type to check whether it is a triangle or not
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if a <= 0 || b <= 0 || c <= 0 {
		k = NaT
	} else if a+b < c || b+c < a || a+c < b {
		k = NaT
	} else if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		k = NaT
	} else if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		k = NaT
	} else if math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1) {
		k = NaT
	} else if a == b && b == c && a == c {
		k = Equ
	} else if a == b || b == c || a == c {
		k = Iso
	} else if a != b && b != c && a != c {
		k = Sca
	}
	return k
}
