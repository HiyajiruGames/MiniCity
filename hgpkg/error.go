package hgpkg

import "fmt"

type IllegalStateError struct {
	State string
}

func (e *IllegalStateError) Error() string {
	return fmt.Sprintf("illegal state error... state is %s", e.State)
}
