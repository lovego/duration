// Package duration parse, format and calculate duration in full and customizable units.
package duration

import (
	"errors"
	"fmt"
	"log"
)

const (
	Year        int64 = 365 * Day
	Month             = 30 * Day
	Week              = 7 * Day
	Day               = 24 * Hour
	Hour              = 60 * Minute
	Minute            = 60 * Second
	Second            = 1000 * Millisecond
	Millisecond       = 1000 * Microsecond
	Microsecond       = 1000 * Nanosecond
	Nanosecond        = 1
)

var unitsMap = make(map[string]int64, 64)

type Unit struct {
	Value      int64    // How many nanoseconds the unit has.
	Name       string   // If Name field is empty, the unit will not be used by Duration.Format.
	PluralName string   // Name in plural form, optional.
	OtherNames []string // Other names, optional.
}

func (unit Unit) validate() error {
	if unit.Value == 0 {
		return errors.New("duration: Unit.Value must not be 0.")
	}
	if unit.Name == "" && len(unit.OtherNames) == 0 {
		return errors.New("duration: Unit.Name and Unit.OtherNames must not both be empty.")
	}
	return unit.validateOrRegister(false)
}

func (unit Unit) register() {
	unit.validateOrRegister(true)
}

func (unit Unit) validateOrRegister(register bool) error {
	for _, name := range unit.allNames() {
		if value := unitsMap[name]; value != 0 && value != unit.Value {
			err := fmt.Errorf(
				`duration: unit name "%s" aready exists but have a different value.`, name,
			)
			if register {
				// This should not happen, beccause of previous validation.
				// If it happend, it must be a bug, so panic here to find it out.
				log.Panic(err)
			}
			return err
		}
		if register {
			unitsMap[name] = unit.Value
		}
	}
	return nil
}

func (unit Unit) allNames() (names []string) {
	if unit.Name != "" {
		names = append(names, unit.Name)
	}
	if unit.PluralName != "" {
		names = append(names, unit.PluralName)
	}
	if len(unit.OtherNames) > 0 {
		names = append(names, unit.OtherNames...)
	}
	return
}
