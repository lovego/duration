package duration

import "errors"

type Duration struct {
	value uint64
	units []Unit
}

func Parse(s string) (uint64, error) {
	var result uint64
	var now uint64
	for _, b := range s {
		if b >= '0' && b <= '9' {
			now = now*10 + uint64(b-'0')
		} else {
			switch b {
			case 'y', 'Y':
				result += now * 365 * 86400
				now = 0
			case 'M':
				result += now * 30 * 86400
				now = 0
			case 'd', 'D':
				result += now * 86400
				now = 0
			case 'h', 'H':
				result += now * 3600
				now = 0
			case 'm':
				result += now * 60
				now = 0
			case 's', 'S':
				result += now
				now = 0
			default:
				return 0, errors.New("invalid duration unit: " + s)
			}
		}
	}
	return result + now, nil
}

func (d *Duration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	integer, err := Parse(s)
	if err != nil {
		return err
	}
	*d = Duration{value: integer}
	return nil
}
