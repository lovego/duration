package duration

import (
	"bytes"
	"strconv"
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
