package formaterrors

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "email") {
		return errors.New("email Already Taken")
	}
	return errors.New("incorrect Details")
}
