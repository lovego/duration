package duration

import (
	"errors"
	"fmt"
)

var unitSetsSlice = make([]unitSet, 0, 4)

type unitSet struct {
	units    []Unit
	namesMap map[string]struct{}
}

// RegisterUnitSet register a set of units which will be used by Parse() and Duration.String().
func RegisterUnitSet(units []Unit) error {
	us := unitSet{units: units}

	if err := us.validate(); err != nil {
		return err
	}
	for _, unit := range us.units {
		// unit.register do not return error, otherwise an unit set may be partially registered.
		unit.register()
	}
	unitSetsSlice = append(unitSetsSlice, us)
	return nil
}

func (us *unitSet) validate() error {
	if len(us.units) == 0 {
		return errors.New("duration: unit set must not be empty.")
	}

	// unit name and unit value in an unit set must be unique
	values, names := make(map[int64]struct{}), make(map[string]struct{})
	for _, unit := range us.units {
		if err := unit.validate(); err != nil {
			return err
		}
		if _, ok := values[unit.Value]; ok {
			return fmt.Errorf(`duration: duplicate value "%d" in unit set.`, unit.Value)
		} else {
			values[unit.Value] = struct{}{}
		}

		for _, name := range unit.allNames() {
			if _, ok := names[name]; ok {
				return fmt.Errorf(`duration: duplicate name "%s" in unit set.`, name)
			} else {
				names[name] = struct{}{}
			}
		}
	}
	us.namesMap = names
	return nil
}

func (us unitSet) match(unitNames []string) bool {
	for _, name := range unitNames {
		if _, ok := us.namesMap[name]; !ok {
			return false
		}
	}
	return true
}

func findUnitSet(unitNames []string) []Unit {
	for _, us := range unitSetsSlice {
		if us.match(unitNames) {
			return us.units
		}
	}
	return EN
}
