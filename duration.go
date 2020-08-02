package duration

import "strconv"

type Duration struct {
	value int64 // max duration: (1<<63 - 1) / (365 * 86400 * 1000 * 1000 * 1000) = 292 year
	units []Unit
}

func (d Duration) String() (s string) {
	if d.value == 0 {
		return "0"
	}

	if d.value < 0 {
		s += "-"
		d.value = -d.value
	}

	if len(d.units) == 0 {
		d.units = EN
	}
	for _, u := range d.units {
		v := d.value / u.Value
		if v != 0 {
			if v > 1 && u.PluralName != "" {
				s += strconv.FormatInt(v, 10) + u.PluralName
			} else {
				s += strconv.FormatInt(v, 10) + u.Name
			}
		}
		d.value %= u.Value
		if d.value == 0 {
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
	drt, err := Parse(s)
	if err != nil {
		return err
	}
	*d = drt
	return nil
}
