package duration

import (
	"errors"
	"fmt"
)

var unitSetsSlice = make([][]Unit, 0, 4)

// RegisterUnitSet register a set of units which will be used by Parse() and Duration.String().
func RegisterUnitSet(unitSet []Unit) error {
	if err := validateUnitSet(unitSet); err != nil {
		return err
	}
	for _, unit := range unitSet {
		// unit.register do not return error, otherwise an unit set may be partially registered.
		unit.register()
	}
	unitSetsSlice = append(unitSetsSlice, unitSet)
	return nil
}

func validateUnitSet(unitSet []Unit) error {
	if len(unitSet) == 0 {
		return errors.New("duration: unit set must not be empty.")
	}

	// unit name and unit value in an unit set must be unique
	values, names := map[int64]bool{}, map[string]bool{}
	for _, unit := range unitSet {
		if err := unit.validate(); err != nil {
			return err
		}
		if values[unit.Value] {
			return fmt.Errorf(`duration: duplicate value "%d" in unit set.`, unit.Value)
		} else {
			values[unit.Value] = true
		}

		for _, name := range unit.allNames() {
			if names[name] {
				return fmt.Errorf(`duration: duplicate name "%s" in unit set.`, name)
			} else {
				names[name] = true
			}
		}
	}
	return nil
}

func FindUnitSet(units []string) []Unit {
	return nil
}
