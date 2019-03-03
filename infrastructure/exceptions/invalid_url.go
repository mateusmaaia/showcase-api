package exceptions

import "fmt"

type InvalidUrlError struct {
}

func (i *InvalidUrlError) Error() string {
	return fmt.Sprintf("Invalid Url")
}
