package duration

import (
	"errors"
	"fmt"
)

// Parse like time.ParseDuration, no layout is needed. All registered units are recoganized.
func Parse(s string) (Duration, error) {
	if s == "" {
		return Duration{}, nil
	}
	orig := s // [-+]?([0-9]*(\.[0-9]*)?[^.0-9]+)+

	negative, s := leadingSignSymbol(s) // Consume [-+]?

	// Special case: if all left is "0" or empty, this is zero.
	if s == "0" {
		return Duration{}, nil
	}

	var value int64
	var unitNames []string
	for s != "" {
		integer, fraction, scale, size, err := leadingNumber(s, orig) // Consume [0-9]*(\.[0-9]*)?
		if err != nil {
			return Duration{}, err
		}
		s = s[size:]

		name, unit, size, err := leadingUnit(s, orig) // [^.0-9]+
		if err != nil {
			return Duration{}, err
		}
		s = s[size:]

		if integer > 0 {
			if integer > (1<<63-1)/unit {
				return Duration{}, errors.New("duration.Parse: duration " + orig + " overflows int64")
			}
			value += integer * unit
			if value < 0 {
				return Duration{}, errors.New("duration.Parse: duration " + orig + " overflows int64")
			}
		}
		if fraction > 0 {
			// float64 is needed to be nanosecond accurate for fractions of hours.
			// v >= 0 && (f*unit/scale) <= 3.6e+12 (ns/h, h is the largest unit)
			value += int64(float64(fraction) * (float64(unit) / scale))
			if value < 0 {
				return Duration{}, errors.New("duration.Parse: duration " + orig + " overflows int64")
			}
		}
		unitNames = append(unitNames, name)
	}
	if negative {
		value = -value
	}

	return Duration{value: value, units: nil}, nil
}

// leadingSignSymbol consumes the leading [-+]? from s.
func leadingSignSymbol(s string) (negative bool, rem string) {
	if s != "" && (s[0] == '-' || s[0] == '+') {
		return s[0] == '-', s[1:]
	}
	return false, s
}

// leadingNumber consumes the leading [0-9]*(\.[0-9]*)? from s.
// the number value is:  integer + fraction/scale
func leadingNumber(s, orig string) (integer, fraction int64, scale float64, size int, err error) {
	// Consume [0-9]*
	integer, iSize, err := leadingInt(s)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	// Consume (\.[0-9]*)?
	fraction, scale, fSize, err := leadingFraction(s[iSize:], orig)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	if iSize == 0 && fSize == 0 { // only unit (e.g. "year" or "minute")
		integer = 1
	}

	size = iSize + fSize
	return
}

// leadingInt consumes the leading [0-9]* from s.
func leadingInt(s string) (x int64, size int, err error) {
	i := 0
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			break
		}
		if x > (1<<63-1)/10 {
			return 0, 0, fmt.Errorf("duration.Parse: int64 overflow: %s", s[:i+1])
		}
		x = x*10 + int64(c) - '0'
		if x < 0 {
			return 0, 0, fmt.Errorf("duration.Parse: int64 overflow: %s", s[:i+1])
		}
	}
	return x, i, nil
}

// leadingFraction consumes the leading [0-9]* from s.
// It is used only for fractions, so does not return an error on overflow,
// it just stops accumulating precision.
func leadingFraction(s, orig string) (x int64, scale float64, size int, err error) {
	if s == "" || s[0] != '.' {
		return 0, 0, 0, nil
	}
	s = s[1:]

	i := 0
	scale = 1
	overflow := false
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			break
		}
		if overflow {
			continue
		}
		if x > (1<<63-1)/10 {
			// It's possible for overflow to give a positive number, so take care.
			overflow = true
			continue
		}
		y := x*10 + int64(c) - '0'
		if y < 0 {
			overflow = true
			continue
		}
		x = y
		scale *= 10
	}
	if i == 0 { // no digits after "." (e.g. ".s" or "-.s")
		return 0, 0, 0, fmt.Errorf("duration.Parse: invalid duration %s", orig)
	}

	return x, scale, i + 1, nil
}

// leadingUnit consumes the leading [^.0-9]+ from s.
func leadingUnit(s, orig string) (name string, unit int64, size int, err error) {
	i := 0
	for ; i < len(s); i++ {
		c := s[i]
		if c == '.' || '0' <= c && c <= '9' {
			break
		}
	}
	if i == 0 {
		return "", 0, 0, errors.New("duration.Parse: missing unit in duration " + orig)
	}
	name = s[:i]
	unit, ok := unitsMap[name]
	if !ok {
		return "", 0, 0, errors.New(
			"duration.Parse: unknown unit " + name + " in duration " + orig,
		)
	}
	size = i
	return
}
