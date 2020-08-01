package duration

import (
	"errors"
	"fmt"
)

var unitSets = make([][]Unit, 0, 4)

var UnitsEN = []Unit{
	{},
}
var UnitsZH = []Unit{}

// RegisterUnitSet register a set of units which will be used by Parse() and Duration.String().
func RegisterUnitSet(unitSet []Unit) error {
	if err := validateUnitSet(unitSet); err != nil {
		return err
	}
	for _, unit := range unitSet {
		// unit.register do not return error, otherwise an unit set may be partially registered.
		unit.register()
	}
	unitSets = append(unitSets, unitSet)
	return nil
}

func validateUnitSet(unitSet []Unit) error {
	if len(unitSet) == 0 {
		return errors.New("duration: unit set must not be empty.")
	}

	values := map[uint64]bool{}
	names := map[string]bool{}
	for _, unit := range unitSet {
		if err := unit.validate(); err != nil {
			return err
		}
		if values[unit.Value] {
			return fmt.Errorf(`duration: duplicte value "%d" in unit set.`, unit.Value)
		} else {
			values[unit.Value] = true
		}
		for _, name := range unit.Names {
			if names[name] {
				return fmt.Errorf(`duration: duplicte name "%s" in unit set.`, name)
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
