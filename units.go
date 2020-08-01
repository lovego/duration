// Package duration parse, format and calculate duration in full and customizable units.
package duration

import (
	"errors"
	"fmt"
	"log"
)

const (
	Year        uint64 = 365 * Day
	Month              = 30 * Day
	Week               = 7 * Day
	Day                = 24 * Hour
	Hour               = 60 * Minute
	Minute             = 60 * Second
	Second             = 1000 * Millisecond
	Millisecond        = 1000 * Microsecond
	Microsecond        = 1000 * Nanosecond
	Nanosecond         = 1
)

var units = make(map[string]uint64, 64)

type Unit struct {
	// how many nanoseconds the unit has.
	Value uint64
	// A unit can have many names, but different unit cann't have same names.
	// If the first name is not empty, it will used by Format func in output.
	Names []string
}

func (unit Unit) validate() error {
	if unit.Value == 0 {
		return errors.New("duration: Unit.Value must not be 0.")
	}
	if len(unit.Names) == 0 {
		return errors.New("duration: Unit.Names must not be empty.")
	}
	for _, name := range unit.Names {
		if name != "" {
			if err := unit.validateName(name); err != nil {
				return err
			}
		}
	}
	return nil
}

func (unit Unit) register() {
	for _, name := range unit.Names {
		if name != "" {
			if err := unit.validateName(name); err != nil {
				// This should not happen, beccause of previous validation.
				// If it happend, it must be a bug, so panic here to find it out.
				log.Panic(err)
			}
			units[name] = unit.Value
		}
	}
}

func (unit Unit) validateName(name string) error {
	if value := units[name]; value != 0 && value != unit.Value {
		return fmt.Errorf(
			`duration: unit name "%s" aready exists but have a different value.`, name,
		)
	}
	return nil
}
