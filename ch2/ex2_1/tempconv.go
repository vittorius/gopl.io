// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

// Celsius represents a temperature value in Celsius scale
type Celsius float64

// Fahrenheit represents a temperature value in Fahrenheit scale
type Fahrenheit float64

// Kelvin represents a temperature value in Kelvin scale
type Kelvin float64

const (
	// AbsoluteZeroC ...
	AbsoluteZeroC Celsius = -273.15
	// FreezingC ...
	FreezingC     Celsius = 0
	// BoilingC ...
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

//!-
