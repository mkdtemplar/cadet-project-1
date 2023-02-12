package validation

import (
	"errors"
	"fmt"
	"regexp"
)

func (v *Validation) ValidateVehicleName(name string) *Validation {
	checkLetters := regexp.MustCompile(`^[a-zA-Z ]*$`)
	if !checkLetters.MatchString(name) {
		v.Err = errors.New("invalid vehicle name")
		return v
	}
	return v
}

func (v *Validation) ValidateVehicleModel(model string) *Validation {
	checkLetters := regexp.MustCompile(`^[a-zA-Z ]*$`)
	if !checkLetters.MatchString(model) {
		v.Err = errors.New("invalid vehicle name")
		return v
	}
	return v
}

func (v *Validation) ValidateVehicleMileage(mileage float32) *Validation {
	checkMileage := regexp.MustCompile(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)(?:\d[eE][+\-]?\d+)?$`)

	if mileage == 0.0 || checkMileage.MatchString(fmt.Sprintf("%f", mileage)) == false {
		v.Err = errors.New("mileage is required or invalid data format mileage must be float32")
		return v
	}
	return v
}
