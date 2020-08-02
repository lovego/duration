package duration

import (
	"bytes"
	"strconv"
)

const (
	minDuration int64 = -1 << 63
	maxDuration int64 = 1<<63 - 1
)

type Duration struct {
	Value int64 // max duration: (1<<63 - 1) / (365 * 86400 * 1000 * 1000 * 1000) = 292 year
	Units []Unit
}

func (d Duration) String() (s string) {
	if d.Value == 0 {
		return "0"
	}

	if d.Value < 0 {
		s += "-"
		d.Value = -d.Value
	}

	for _, u := range d.Units {
		if u.Name != "" {
			v := d.Value / u.Value
			if v != 0 {
				if v > 1 && u.PluralName != "" {
					s += strconv.FormatInt(v, 10) + u.PluralName
				} else {
					s += strconv.FormatInt(v, 10) + u.Name
				}
			}
			d.Value %= u.Value
			if d.Value == 0 {
				break
			}
		}
	}
	return
}

func (d *Duration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	duration, err := Parse(s)
	if err != nil {
		return err
	}
	*d = duration
	return nil
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	duration, err := Parse(string(bytes.Trim(b, `"`)))
	if err != nil {
		return err
	}
	*d = duration
	return nil
}

// Truncate returns the result of rounding d toward zero to a multiple of m.
// If m <= 0, Truncate returns d unchanged.
func (d Duration) Truncate(m int64) Duration {
	if m >= 2 && d.Value != 0 {
		d.Value -= d.Value % m
	}
	return d
}

// Round returns the result of rounding d to the nearest multiple of m.
// The rounding behavior for halfway values is to round away from zero.
// If the result exceeds the maximum (or minimum)
// value that can be stored in a Duration,
// Round returns the maximum (or minimum) duration.
// If m <= 0, Round returns d unchanged.
func (d Duration) Round(m int64) Duration {
	if m >= 2 && d.Value != 0 {
		r := d.Value % m
		if d.Value > 0 {
			if lessThanHalf(r, m) {
				d.Value -= r
			} else if d1 := d.Value + (m - r); d1 > d.Value {
				d.Value = d1
			} else { // overflow
				d.Value = maxDuration
			}
		} else {
			r = -r
			if lessThanHalf(r, m) {
				d.Value += r
			} else if d1 := d.Value - (m - r); d1 < d.Value {
				d.Value = d1
			} else { // overflow
				d.Value = minDuration
			}
		}
	}
	return d
}

// lessThanHalf reports whether x+x < y but avoids overflow,
// assuming x and y are both positive (Duration is signed).
func lessThanHalf(x, y int64) bool {
	return uint64(x)+uint64(x) < uint64(y)
}
