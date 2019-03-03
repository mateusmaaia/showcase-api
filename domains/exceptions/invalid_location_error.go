package exceptions

import "fmt"

type InvalidLocationError struct {
}

func (i *InvalidLocationError) Error() string {
	return fmt.Sprintf("Invalid Location")
}