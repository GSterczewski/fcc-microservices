package services

import (
	"fmt"
	"io"
)

//Timestamp -
type Timestamp struct{}

//Run - main function for the service
func (ts Timestamp) Run(w io.Writer) {
	fmt.Fprint(w, "Timestamp service")
}
