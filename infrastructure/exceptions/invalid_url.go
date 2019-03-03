package exceptions

import "fmt"

type InvalidUrlError struct {
	Url string
}

func (i *InvalidUrlError) Error() string {
	return fmt.Sprintf("Invalid Url: %v", i.Url)
}
