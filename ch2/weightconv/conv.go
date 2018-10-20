package weightconv

const poundsInKg = 2.20462

// ToPound converts from Kg to Pound
func (k Kg) ToPound() Pound { return Pound(k * poundsInKg) }

// ToKg converts from Pound to Kg
func (p Pound) ToKg() Kg { return Kg(p / poundsInKg) }
