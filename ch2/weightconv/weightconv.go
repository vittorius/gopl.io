package weightconv

import (
	"fmt"
)

// Kg is a type for kilograms
type Kg float64

// Pound is a type for pounds
type Pound float64

const (
	// KgAbbr is the abbreviation for kilogram unit
	KgAbbr = "kg"
	// PoundAbbr is the abbreviation for pound unit
	PoundAbbr = "lb"
)

// String returns the human-readable representation of a Kg object
func (k Kg) String() string { return fmt.Sprintf("%g %s", k, KgAbbr) }

// String returns the human-readable representation of a Pound object
func (p Pound) String() string { return fmt.Sprintf("%g %s", p, PoundAbbr) }
